package test

import (
	"auth-service/internal/config"
	"auth-service/internal/storage"
	"context"
	"log/slog"
	"testing"

	pb "auth-service/genprotos/auth_pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupKitchenTest(t *testing.T) *storage.AuthSt {
	cfg := &config.Config{
		// Test database configuration
	}
	logger := slog.Default()
	authSt, err := storage.New(cfg, logger)
	require.NoError(t, err)
	return authSt
}

func TestCreateKitchen(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.CreateKitchenRequest{
		OwnerId:     "some-owner-id",
		Name:        "Test Kitchen",
		Description: "A test kitchen",
		CuisineType: "Italian",
		Address:     "123 Test St",
		PhoneNumber: "1234567890",
	}

	resp, err := authSt.CreateKitchen(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.KitchenId)
	assert.Equal(t, req.Name, resp.Name)
}

func TestUpdateKitchen(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.UpdateKitchenRequest{
		KitchenId:   "some-kitchen-id",
		Name:        "Updated Kitchen",
		Description: "An updated test kitchen",
		Address:     "456 New St",
		PhoneNumber: "0987654321",
	}

	resp, err := authSt.UpdateKitchen(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Name, resp.Name)
	assert.Equal(t, req.Description, resp.Description)
}

func TestGetKitchen(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.GetKitchenRequest{
		KitchenId: "some-kitchen-id",
	}

	resp, err := authSt.GetKitchen(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.KitchenId, resp.KitchenId)
}

func TestListKitchens(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.ListKitchensRequest{
		Page:  1,
		Limit: 10,
	}

	resp, err := authSt.ListKitchens(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Kitchens)
	assert.Equal(t, req.Page, resp.Page)
	assert.Equal(t, req.Limit, resp.Limit)
}

func TestSearchKitchens(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.SearchKitchensRequest{
		Name:  "Test",
		Page:  1,
		Limit: 10,
	}

	resp, err := authSt.SearchKitchens(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Kitchens)
	assert.Equal(t, req.Page, resp.Page)
	assert.Equal(t, req.Limit, resp.Limit)
}

func TestDoesKitchenExist(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.DoesKitchenExistRequest{
		KitchenId: "some-kitchen-id",
	}

	resp, err := authSt.DoesKitchenExist(ctx, req)
	assert.NoError(t, err)
	assert.True(t, resp.Exists)
}

func TestIncrementTotalOrders(t *testing.T) {
	authSt := setupKitchenTest(t)
	ctx := context.Background()

	req := &pb.IncrementTotalOrdersRequest{
		KitchenId: "some-kitchen-id",
	}

	resp, err := authSt.IncrementTotalOrders(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, "Total orders incremented successfully", resp.Ok)
}
