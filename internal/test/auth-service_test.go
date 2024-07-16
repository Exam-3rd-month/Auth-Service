package test

import (
	pb "auth-service/genprotos/auth_pb"
	"auth-service/internal/config"
	"auth-service/internal/storage"
	"context"
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAuthSt(t *testing.T) *storage.AuthSt {
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

	authSt, err := storage.New(configs, logger)
	require.NoError(t, err)
	return authSt
}

func TestRegister(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	req := &pb.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpass",
		FullName: "Test User",
		UserType: "customer",
	}

	resp, err := authSt.Register(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.UserId)
	assert.Equal(t, req.Username, resp.Username)
	assert.Equal(t, req.Email, resp.Email)
	assert.Equal(t, req.FullName, resp.FullName)
	assert.Equal(t, req.UserType, resp.UserType)
}

func TestLogin(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	// First, register a user
	registerReq := &pb.RegisterRequest{
		Username: "logintest",
		Email:    "login@example.com",
		Password: "loginpass",
		FullName: "Login Test",
		UserType: "customer",
	}
	_, err := authSt.Register(ctx, registerReq)
	require.NoError(t, err)

	// Now, try to login
	loginReq := &pb.LoginRequest{
		Email:    "login@example.com",
		Password: "loginpass",
	}

	resp, err := authSt.Login(ctx, loginReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.AccessToken)
}

func TestGetProfile(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	// First, register a user
	registerReq := &pb.RegisterRequest{
		Username: "profiletest",
		Email:    "profile@example.com",
		Password: "profilepass",
		FullName: "Profile Test",
		UserType: "customer",
	}
	regResp, err := authSt.Register(ctx, registerReq)
	require.NoError(t, err)

	// Now, get the profile
	profileReq := &pb.GetProfileRequest{
		UserId: regResp.UserId,
	}

	resp, err := authSt.GetProfile(ctx, profileReq)
	assert.NoError(t, err)
	assert.Equal(t, regResp.UserId, resp.UserId)
	assert.Equal(t, registerReq.Username, resp.Username)
	assert.Equal(t, registerReq.Email, resp.Email)
	assert.Equal(t, registerReq.FullName, resp.FullName)
	assert.Equal(t, registerReq.UserType, resp.UserType)
}

func TestUpdateProfile(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	// First, register a user
	registerReq := &pb.RegisterRequest{
		Username: "updatetest",
		Email:    "update@example.com",
		Password: "updatepass",
		FullName: "Update Test",
		UserType: "customer",
	}
	regResp, err := authSt.Register(ctx, registerReq)
	require.NoError(t, err)

	// Now, update the profile
	updateReq := &pb.UpdateProfileRequest{
		UserId:      regResp.UserId,
		FullName:    "Updated Name",
		Address:     "123 Update St",
		PhoneNumber: "1234567890",
	}

	resp, err := authSt.UpdateProfile(ctx, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, updateReq.UserId, resp.UserId)
	assert.Equal(t, updateReq.FullName, resp.FullName)
	assert.Equal(t, updateReq.Address, resp.Address)
	assert.Equal(t, updateReq.PhoneNumber, resp.PhoneNumber)
}

func TestResetPassword(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	// First, register a user
	registerReq := &pb.RegisterRequest{
		Username: "resettest",
		Email:    "reset@example.com",
		Password: "resetpass",
		FullName: "Reset Test",
		UserType: "customer",
	}
	_, err := authSt.Register(ctx, registerReq)
	require.NoError(t, err)

	// Now, reset the password
	resetReq := &pb.ResetPasswordRequest{
		Email:    "reset@example.com",
		Password: "newpassword",
	}

	resp, err := authSt.ResetPassword(ctx, resetReq)
	assert.NoError(t, err)
	assert.Equal(t, "Password successfully updated", resp.Message)
}

func TestLogout(t *testing.T) {
	authSt := setupAuthSt(t)
	ctx := context.Background()

	// First, register a user
	registerReq := &pb.RegisterRequest{
		Username: "logouttest",
		Email:    "logout@example.com",
		Password: "logoutpass",
		FullName: "Logout Test",
		UserType: "customer",
	}
	regResp, err := authSt.Register(ctx, registerReq)
	require.NoError(t, err)

	// Now, logout
	logoutReq := &pb.LogoutRequest{
		UserId: regResp.UserId,
	}

	resp, err := authSt.Logout(ctx, logoutReq)
	assert.NoError(t, err)
	assert.Equal(t, "User successfully logged out", resp.Message)
}
