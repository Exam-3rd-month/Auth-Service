package test

import (
	"auth-service/genprotos/auth_pb"
	"context"
	"testing"
)

func TestCreateKitchen(t *testing.T) {
	authSt := setupAuthSt()
	ctx := context.Background()

	test := auth_pb.RegisterRequest{
		Username: "testusername",
		Email:    "testemail",
		Password: "testpassword",
		FullName: "testfullname",
		UserType: "user",
	}

	resp, err := authSt.Register(ctx, &test)
	if err != nil {
		t.Error("faliled to register user", resp)
	}

	test2 := auth_pb.CreateKitchenRequest{
		OwnerId:     resp.UserId,
		Name:        "testname",
		Description: "testdescription",
		CuisineType: "testcuisinetype",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := authSt.CreateKitchen(ctx, &test2)
	if err != nil {
		t.Error("faliled to create kitchen", resp2)
	}

	_, err = authSt.DeleteKitchen(ctx, &auth_pb.DeleteKitchenRequest{KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestUpdateKitchen(t *testing.T) {
	authSt := setupAuthSt()
	ctx := context.Background()

	test := auth_pb.RegisterRequest{
		Username: "testusername",
		Email:    "testemail",
		Password: "testpassword",
		FullName: "testfullname",
		UserType: "user",
	}

	resp, err := authSt.Register(ctx, &test)
	if err != nil {
		t.Error("faliled to register user", resp)
	}

	test2 := auth_pb.CreateKitchenRequest{
		OwnerId:     resp.UserId,
		Name:        "testname",
		Description: "testdescription",
		CuisineType: "testcuisinetype",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := authSt.CreateKitchen(ctx, &test2)
	if err != nil {
		t.Error("faliled to create kitchen", resp2)
	}

	test3 := auth_pb.UpdateKitchenRequest{
		KitchenId:   resp2.KitchenId,
		Name:        "testname2",
		Description: "testdescription2",
		Address:     "testaddress2",
		PhoneNumber: "testphonenumber2",
	}

	resp3, err := authSt.UpdateKitchen(ctx, &test3)
	if err != nil {
		t.Error("faliled to update kitchen", resp3)
	}

	_, err = authSt.DeleteKitchen(ctx, &auth_pb.DeleteKitchenRequest{KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestGetKitchen(t *testing.T) {
	authSt := setupAuthSt()
	ctx := context.Background()

	test := auth_pb.RegisterRequest{
		Username: "testusername",
		Email:    "testemail",
		Password: "testpassword",
		FullName: "testfullname",
		UserType: "user",
	}

	resp, err := authSt.Register(ctx, &test)
	if err != nil {
		t.Error("faliled to register user", resp)
	}

	test2 := auth_pb.CreateKitchenRequest{
		OwnerId:     resp.UserId,
		Name:        "testname",
		Description: "testdescription",
		CuisineType: "testcuisinetype",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := authSt.CreateKitchen(ctx, &test2)
	if err != nil {
		t.Error("faliled to create kitchen", resp2)
	}

	test3 := auth_pb.GetKitchenRequest{
		KitchenId: resp2.KitchenId,
	}

	resp3, err := authSt.GetKitchen(ctx, &test3)
	if err != nil {
		t.Error("faliled to get kitchen", resp3)
	}

	_, err = authSt.DeleteKitchen(ctx, &auth_pb.DeleteKitchenRequest{KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}
