package service

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
)

// 1
func (s *AuthServiceSt) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.logger.Info("register request")
	return s.service.Register(ctx, req)
}

// 3
func (s *AuthServiceSt) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	s.logger.Info("get profile request")
	return s.service.GetProfile(ctx, req)
}

// 4
func (s *AuthServiceSt) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	s.logger.Info("update profile request")
	return s.service.UpdateProfile(ctx, req)
}

// 5
func (s *AuthServiceSt) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	s.logger.Info("reset password request")
	return s.service.ResetPassword(ctx, req)
}

// 7
func (s *AuthServiceSt) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	s.logger.Info("logout request")
	return s.service.Logout(ctx, req)
}

// 13
func (s *AuthServiceSt) DoesUserExist(ctx context.Context, req *pb.DoesUserExistRequest) (*pb.DoesUserExistResponse, error) {
	s.logger.Info("does user exist request")
	return s.service.DoesUserExist(ctx, req)
}

// 15
func (s *AuthServiceSt) IsValidToken(ctx context.Context, req *pb.IsValidTokenRequest) (*pb.IsValidTokenResponse, error) {
	s.logger.Info("is valid token request")
	return s.service.IsValidToken(ctx, req)
}
