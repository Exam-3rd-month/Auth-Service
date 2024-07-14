package service

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
)

func (s *AuthServiceSt) CreateKitchen(ctx context.Context, in *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	s.logger.Info("create kitchen request")
	return s.service.CreateKitchen(ctx, in)
}

func (s *AuthServiceSt) UpdateKitchen(ctx context.Context, in *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	s.logger.Info("update kitchen request")
	return s.service.UpdateKitchen(ctx, in)
}

func (s *AuthServiceSt) GetKitchen(ctx context.Context, in *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	s.logger.Info("get kitchen request")
	return s.service.GetKitchen(ctx, in)
}
