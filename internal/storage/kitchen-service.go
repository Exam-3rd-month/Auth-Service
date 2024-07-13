package storage

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// 8
func (s *AuthSt) CreateKitchen(ctx context.Context, in *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	kitchen_id := uuid.New().String()
	created_at := time.Now()

	qyery, args, err := s.queryBuilder.Insert("kitchens").
		Columns(
			"kitchen_id",
			"owner_id",
			"name",
			"description",
			"cuisine_type",
			"address",
			"phone_number",
			"created_at",
		).
		Values(
			kitchen_id,
			in.OwnerId,
			in.Name,
			in.Description,
			in.CuisineType,
			in.Address,
			in.PhoneNumber,
			created_at,
		).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, qyery, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.CreateKitchenResponse{
		KitchenId:   kitchen_id,
		OwnerId:     in.OwnerId,
		Name:        in.Name,
		Description: in.Description,
		CuisineType: in.CuisineType,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		Rating:      0,
		CreatedAt:   created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 9
func (s *AuthSt) UpdateKitchen(ctx context.Context, in *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	updated_at := time.Now()

	query, args, err := s.queryBuilder.Update("kitchens").
		Set("name", in.Name).
		Set("description", in.Description).
		Set("address", in.Address).
		Set("phone_number", in.PhoneNumber).
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		Prefix("RETURNING owner_id, cousine_type, rating").
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	var owner_id string
	var cuisine_type string
	var rating float32

	err = row.Scan(&owner_id, &cuisine_type, &rating)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.UpdateKitchenResponse{
		KitchenId:   in.KitchenId,
		OwnerId:     owner_id,
		Name:        in.Name,
		Description: in.Description,
		CuisineType: cuisine_type,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		Rating:      rating,
		UpdatedAt:   updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 10
func (s *AuthSt) GetKitchen(ctx context.Context, in *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"owner_id",
		"name",
		"description",
		"cuisine_type",
		"address",
		"phone_number",
		"rating",
		"total_orders",
		"created_at",
		"updated_at").
		From("kitchens").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var response pb.GetKitchenResponse

	err = row.Scan(
		&response.OwnerId,
		&response.Name,
		&response.Description,
		&response.CuisineType,
		&response.Address,
		&response.PhoneNumber,
		&response.Rating,
		&response.TotalOrders,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &response, nil
}

// 11
func (s *AuthSt) ListKitchens(ctx context.Context, in *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	return nil, nil
}

// 12
func (s *AuthSt) SearchKitchens(ctx context.Context, in *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	return nil, nil
}
