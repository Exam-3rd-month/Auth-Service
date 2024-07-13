package storage

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"auth-service/internal/config"

	pb "auth-service/genprotos/auth_pb"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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
	return nil, nil
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

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&response.UserId,
		&response.Username,
		&response.Email,
		&response.FullName,
		&response.UserType,
		&response.Address,
		&response.PhoneNumber,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
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

	return &pb.UpdateProfileResponse{
		UserId:      in.UserId,
		FullName:    in.FullName,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		UpdatedAt:   updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 5
func (s *AuthSt) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	return nil, nil
}

// 6
func (s *AuthSt) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return nil, nil
}

// 7
func (s *AuthSt) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}
