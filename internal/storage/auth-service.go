package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"auth-service/internal/config"
	jwttokens "auth-service/internal/jwt"

	pb "auth-service/genprotos/auth_pb"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	AuthSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
		logger       *slog.Logger
	}
)

func New(config *config.Config, logger *slog.Logger) (*AuthSt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &AuthSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:       logger,
	}, nil
}

// 1
func (s *AuthSt) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user_id := uuid.New().String()
	created_at := time.Now()
	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	query, args, err := s.queryBuilder.Insert("users").
		Columns(
			"user_id",
			"username",
			"email",
			"password",
			"full_name",
			"user_type",
			"created_at").
		Values(
			user_id,
			in.Username,
			in.Email,
			hashedPassword,
			in.FullName,
			in.UserType,
			created_at).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId:    user_id,
		Username:  in.Username,
		Email:     in.Email,
		FullName:  in.FullName,
		UserType:  in.UserType,
		CreatedAt: created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 2
func (s *AuthSt) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	query, args, err := s.queryBuilder.Select("password", "user_id").
		From("users").
		Where(sq.Eq{"email": in.Email}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var password string
	var user_id string
	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&password, &user_id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	if checkPassword(hashedPassword, password) {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	token, err := jwttokens.GenerateToken(user_id, in.Email)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}

// 3
func (s *AuthSt) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"user_id",
		"username",
		"email",
		"full_name",
		"user_type",
		"address",
		"phone_number",
		"created_at",
		"updated_at").
		From("users").
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response pb.GetProfileResponse
	var tempAddress, tempPhoneNumber sql.NullString

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&response.UserId,
		&response.Username,
		&response.Email,
		&response.FullName,
		&response.UserType,
		&tempAddress,
		&tempPhoneNumber,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	if tempAddress.Valid {
		response.Address = tempAddress.String
	}
	if tempPhoneNumber.Valid {
		response.PhoneNumber = tempPhoneNumber.String
	}

	return &response, nil
}

// 4
func (s *AuthSt) UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	updated_at := time.Now()

	query, args, err := s.queryBuilder.Update("users").
		Set("full_name", in.FullName).
		Set("address", in.Address).
		Set("phone_number", in.PhoneNumber).
		Set("updated_at", updated_at).
		Where(sq.Eq{"user_id": in.UserId}).
		Suffix("RETURNING username, email, user_type").
		ToSql()

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	var username, email, user_type string
	err = row.Scan(
		&username,
		&email,
		&user_type,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		UserId:      in.UserId,
		Username:    username,
		Email:       email,
		FullName:    in.FullName,
		UserType:    user_type,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		UpdatedAt:   updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 5
func (s *AuthSt) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	hashedPassword, err := hashPassword(in.ConfirmPassword)
	if err != nil {
		s.logger.Error("Failed to hash password", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	query, args, err := s.queryBuilder.Update("users").
		Set("password", hashedPassword).
		Where(sq.Eq{"email": in.Email}).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build SQL query", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute SQL query", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Failed to get affected rows", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	if rowsAffected == 0 {
		s.logger.Warn("No user found with the provided email")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.ResetPasswordResponse{Message: "Password successfully updated"}, nil
}

// 7
func (s *AuthSt) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	query, args, err := s.queryBuilder.Update("users").
		Set("is_expired", true).
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		s.logger.Warn("No user found with the provided user_id")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.LogoutResponse{Message: "User successfully logged out"}, nil
}

// 13
func (s *AuthSt) DoesUserExist(ctx context.Context, in *pb.DoesUserExistRequest) (*pb.DoesUserExistResponse, error) {
	query, args, err := s.queryBuilder.Select("user_id").
		From("users").
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var user_id string

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&user_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.DoesUserExistResponse{Exists: false}, nil
		}
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.DoesUserExistResponse{Exists: true}, nil
}

// 15
func (s *AuthSt) IsValidToken(ctx context.Context, in *pb.IsValidTokenRequest) (*pb.IsValidTokenResponse, error) {
	query, args, err := s.queryBuilder.Select("is_expired").
		From("users").
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var is_expired bool

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&is_expired,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.IsValidTokenResponse{Valid: false}, nil
		}
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.IsValidTokenResponse{Valid: !is_expired}, nil
}

func (s *AuthSt) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	query, args, err := s.queryBuilder.
		Delete("users").
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build delete query: %w", err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute delete query: %w", err)
	}

	return &pb.DeleteUserResponse{Ok: "User deleted successfully"}, nil
}

// 20
func (s *AuthSt) IsValidUser(ctx context.Context, in *pb.IsValidUserRequest) (*pb.IsValidUserResponse, error) {
	query, args, err := s.queryBuilder.Select("user_id").
		From("users").
		Where(sq.Eq{"email": in.Email}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var user_id string

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&user_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.IsValidUserResponse{Valid: false}, nil
		}
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.IsValidUserResponse{Valid: true}, nil
}
