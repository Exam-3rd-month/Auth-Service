package test

import (
	"auth-service/genprotos/auth_pb"
	"auth-service/internal/config"
	"auth-service/internal/storage"
	"context"
	"log"
	"log/slog"
	"os"
	"testing"
)

func setupAuthSt() *storage.AuthSt {
	logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs, logger)
	if err != nil {
		log.Fatal(err)
	}

	return storage
}

func TestRegister(t *testing.T) {
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

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestLogin(t *testing.T) {
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

	test2 := auth_pb.LoginRequest{
		Email:    "testemail",
		Password: "testpassword",
	}

	resp2, err := authSt.Login(ctx, &test2)
	if err != nil {
		t.Error("faliled to login user", resp2)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestGetProfile(t *testing.T) {
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

	test2 := auth_pb.GetProfileRequest{
		UserId: resp.UserId,
	}

	resp2, err := authSt.GetProfile(ctx, &test2)
	if err != nil {
		t.Error("faliled to get profile", resp2)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestUpdateProfile(t *testing.T) {
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

	test2 := auth_pb.UpdateProfileRequest{
		UserId:      resp.UserId,
		FullName:    "testfullname2",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := authSt.UpdateProfile(ctx, &test2)
	if err != nil {
		t.Error("faliled to update profile", resp2)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestResetPassword(t *testing.T) {
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

	test2 := auth_pb.ResetPasswordRequest{
		Email:    "testemail",
		ConfirmPassword: "testpassword",
	}

	resp2, err := authSt.ResetPassword(ctx, &test2)
	if err != nil {
		t.Error("faliled to reset password", resp2)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}

func TestLogout(t *testing.T) {
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

	test2 := auth_pb.LoginRequest{
		Email:    "testemail",
		Password: "testpassword",
	}

	resp2, err := authSt.Login(ctx, &test2)
	if err != nil {
		t.Error("faliled to login user", resp2)
	}

	test3 := auth_pb.LogoutRequest{
		UserId: resp.UserId,
	}

	resp3, err := authSt.Logout(ctx, &test3)
	if err != nil {
		t.Error("faliled to logout user", resp3)
	}

	_, err = authSt.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}
}
