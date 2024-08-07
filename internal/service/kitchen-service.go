package service

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
)

// 8
func (s *AuthServiceSt) CreateKitchen(ctx context.Context, in *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	s.logger.Info("create kitchen request")
	return s.service.CreateKitchen(ctx, in)
}

// 9
func (s *AuthServiceSt) UpdateKitchen(ctx context.Context, in *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	s.logger.Info("update kitchen request")
	return s.service.UpdateKitchen(ctx, in)
}

// 10
func (s *AuthServiceSt) GetKitchen(ctx context.Context, in *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	s.logger.Info("get kitchen request")
	return s.service.GetKitchen(ctx, in)
}

// 11
func (s *AuthServiceSt) ListKitchens(ctx context.Context, in *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	s.logger.Info("list kitchens request")
	return s.service.ListKitchens(ctx, in)
}

// 12
func (s *AuthServiceSt) SearchKitchens(ctx context.Context, in *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	s.logger.Info("search kitchens request")
	return s.service.SearchKitchens(ctx, in)
}

// 14
func (s *AuthServiceSt) DoesKitchenExist(ctx context.Context, in *pb.DoesKitchenExistRequest) (*pb.DoesKitchenExistResponse, error) {
	s.logger.Info("does kitchen exist request")
	return s.service.DoesKitchenExist(ctx, in)
}

// 16
func (s *AuthServiceSt) IncrementTotalOrders(ctx context.Context, in *pb.IncrementTotalOrdersRequest) (*pb.IncrementTotalOrdersResponse, error) {
	s.logger.Info("increment total orders request")
	return s.service.IncrementTotalOrders(ctx, in)
}

// 17
func (s *AuthServiceSt) IncrementOrderRating(ctx context.Context, in *pb.IncrementOrderRatingRequest) (*pb.IncrementOrderRatingResponse, error) {
	s.logger.Info("increment order rating request")
	return s.service.IncrementOrderRating(ctx, in)
}

func (s *AuthServiceSt) DeleteKitchen(ctx context.Context, in *pb.DeleteKitchenRequest) (*pb.DeleteKitchenResponse, error) {
	s.logger.Info("delete kitchen request")
	return s.service.DeleteKitchen(ctx, in)
}
