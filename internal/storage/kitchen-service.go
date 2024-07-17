package storage

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
	"database/sql"
	"fmt"
	"log"
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
		Suffix("RETURNING owner_id, cuisine_type, rating").
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var owner_id string
	var cuisine_type string
	var rating float32

	row := s.db.QueryRowContext(ctx, query, args...)

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
	response.KitchenId = in.KitchenId

	return &response, nil
}

// 11
func (s *AuthSt) ListKitchens(ctx context.Context, in *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	// Calculate total kitchens
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("kitchens").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build count query", "error", err)
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to execute count query", "error", err)
		return nil, err
	}

	// Set default values for limit and page
	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := in.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	// Main query to fetch kitchens
	query, args, err := s.queryBuilder.Select(
		"kitchen_id",
		"name",
		"cuisine_type",
		"rating",
		"total_orders").
		From("kitchens").
		OrderBy("rating DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var kitchens []*pb.KitchenListItem

	for rows.Next() {
		kitchen := &pb.KitchenListItem{}
		err = rows.Scan(
			&kitchen.Id,
			&kitchen.Name,
			&kitchen.CuisineType,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		kitchens = append(kitchens, kitchen)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, err
	}

	return &pb.ListKitchensResponse{
		Kitchens: kitchens,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

// 12
func (s *AuthSt) SearchKitchens(ctx context.Context, in *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	// Calculate total matching kitchens
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("kitchens").
		Where(sq.Like{"name": "%" + in.Name + "%"}).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build count query", "error", err)
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to execute count query", "error", err)
		return nil, err
	}

	// Set default values for limit and page
	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := in.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	// Main query to fetch matching kitchens
	query, args, err := s.queryBuilder.Select(
		"kitchen_id",
		"name",
		"cuisine_type",
		"rating",
		"total_orders").
		From("kitchens").
		Where(sq.Like{"name": "%" + in.Name + "%"}).
		OrderBy("rating DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var kitchens []*pb.KitchenListItem

	for rows.Next() {
		kitchen := &pb.KitchenListItem{}
		err = rows.Scan(
			&kitchen.Id,
			&kitchen.Name,
			&kitchen.CuisineType,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		kitchens = append(kitchens, kitchen)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, err
	}

	return &pb.SearchKitchensResponse{
		Kitchens: kitchens,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

// 14
func (s *AuthSt) DoesKitchenExist(ctx context.Context, in *pb.DoesKitchenExistRequest) (*pb.DoesKitchenExistResponse, error) {
	query, args, err := s.queryBuilder.Select("kitchen_id").
		From("kitchens").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	var kitchen_id string

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&kitchen_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.DoesKitchenExistResponse{Exists: false}, nil
		}
		s.logger.Error(err.Error())
		return nil, err
	}
	log.Println(query)

	return &pb.DoesKitchenExistResponse{Exists: true}, nil
}

// 16
func (s *AuthSt) IncrementTotalOrders(ctx context.Context, in *pb.IncrementTotalOrdersRequest) (*pb.IncrementTotalOrdersResponse, error) {
	query, args, err := s.queryBuilder.Update("kitchens").
		Set("total_orders", sq.Expr("total_orders +?", 1)).
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}

	return &pb.IncrementTotalOrdersResponse{
		Ok: "Total orders incremented successfully",
	}, nil
}

// 17
func (s *AuthSt) IncrementOrderRating(ctx context.Context, in *pb.IncrementOrderRatingRequest) (*pb.IncrementOrderRatingResponse, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := s.queryBuilder.Select("rating").
		From("kitchens").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	var rating float32

	err = tx.QueryRowContext(ctx, query, args...).Scan(&rating)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}

	var updated_rating float32
	if rating == 0 {
		updated_rating = in.Rating
	} else {
		updated_rating = (rating + in.Rating) / 2
	}

	query, args, err = s.queryBuilder.Update("kitchens").
		Set("rating", updated_rating).
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.IncrementOrderRatingResponse{
		Ok: "Order rating incremented successfully",
	}, nil
}

func (s *AuthSt) DeleteKitchen(ctx context.Context, in *pb.DeleteKitchenRequest) (*pb.DeleteKitchenResponse, error) {
	query, args, err := s.queryBuilder.
		Delete("kitchens").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build delete query: %w", err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute delete query: %w", err)
	}

	return &pb.DeleteKitchenResponse{Ok: "Kitchen deleted successfully"}, nil
}
