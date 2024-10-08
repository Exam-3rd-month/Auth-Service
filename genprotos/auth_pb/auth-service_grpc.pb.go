// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: auth-service/auth-service.proto

package auth_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthService_Register_FullMethodName             = "/AuthService/Register"
	AuthService_Login_FullMethodName                = "/AuthService/Login"
	AuthService_GetProfile_FullMethodName           = "/AuthService/GetProfile"
	AuthService_UpdateProfile_FullMethodName        = "/AuthService/UpdateProfile"
	AuthService_ResetPassword_FullMethodName        = "/AuthService/ResetPassword"
	AuthService_Logout_FullMethodName               = "/AuthService/Logout"
	AuthService_CreateKitchen_FullMethodName        = "/AuthService/CreateKitchen"
	AuthService_UpdateKitchen_FullMethodName        = "/AuthService/UpdateKitchen"
	AuthService_GetKitchen_FullMethodName           = "/AuthService/GetKitchen"
	AuthService_ListKitchens_FullMethodName         = "/AuthService/ListKitchens"
	AuthService_SearchKitchens_FullMethodName       = "/AuthService/SearchKitchens"
	AuthService_DoesUserExist_FullMethodName        = "/AuthService/DoesUserExist"
	AuthService_DoesKitchenExist_FullMethodName     = "/AuthService/DoesKitchenExist"
	AuthService_IsValidToken_FullMethodName         = "/AuthService/IsValidToken"
	AuthService_IncrementTotalOrders_FullMethodName = "/AuthService/IncrementTotalOrders"
	AuthService_IncrementOrderRating_FullMethodName = "/AuthService/IncrementOrderRating"
	AuthService_DeleteUser_FullMethodName           = "/AuthService/DeleteUser"
	AuthService_DeleteKitchen_FullMethodName        = "/AuthService/DeleteKitchen"
	AuthService_IsValidUser_FullMethodName          = "/AuthService/IsValidUser"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// 1 Done
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// 2
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// 3 Done
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// 4 Done
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
	// 5 Done
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	// 7
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	// 8 Done
	CreateKitchen(ctx context.Context, in *CreateKitchenRequest, opts ...grpc.CallOption) (*CreateKitchenResponse, error)
	// 9 Done
	UpdateKitchen(ctx context.Context, in *UpdateKitchenRequest, opts ...grpc.CallOption) (*UpdateKitchenResponse, error)
	// 10 Done
	GetKitchen(ctx context.Context, in *GetKitchenRequest, opts ...grpc.CallOption) (*GetKitchenResponse, error)
	// 11 Done
	ListKitchens(ctx context.Context, in *ListKitchensRequest, opts ...grpc.CallOption) (*ListKitchensResponse, error)
	// 12 Done
	SearchKitchens(ctx context.Context, in *SearchKitchensRequest, opts ...grpc.CallOption) (*SearchKitchensResponse, error)
	// Yordamchi funksiyalar
	// 13 Done
	DoesUserExist(ctx context.Context, in *DoesUserExistRequest, opts ...grpc.CallOption) (*DoesUserExistResponse, error)
	// 14 Done
	DoesKitchenExist(ctx context.Context, in *DoesKitchenExistRequest, opts ...grpc.CallOption) (*DoesKitchenExistResponse, error)
	// 15
	IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*IsValidTokenResponse, error)
	// 16
	IncrementTotalOrders(ctx context.Context, in *IncrementTotalOrdersRequest, opts ...grpc.CallOption) (*IncrementTotalOrdersResponse, error)
	// 17
	IncrementOrderRating(ctx context.Context, in *IncrementOrderRatingRequest, opts ...grpc.CallOption) (*IncrementOrderRatingResponse, error)
	// 18
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// 19
	DeleteKitchen(ctx context.Context, in *DeleteKitchenRequest, opts ...grpc.CallOption) (*DeleteKitchenResponse, error)
	// 20
	IsValidUser(ctx context.Context, in *IsValidUserRequest, opts ...grpc.CallOption) (*IsValidUserResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AuthService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, AuthService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateKitchen(ctx context.Context, in *CreateKitchenRequest, opts ...grpc.CallOption) (*CreateKitchenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateKitchenResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateKitchen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateKitchen(ctx context.Context, in *UpdateKitchenRequest, opts ...grpc.CallOption) (*UpdateKitchenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateKitchenResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateKitchen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetKitchen(ctx context.Context, in *GetKitchenRequest, opts ...grpc.CallOption) (*GetKitchenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKitchenResponse)
	err := c.cc.Invoke(ctx, AuthService_GetKitchen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListKitchens(ctx context.Context, in *ListKitchensRequest, opts ...grpc.CallOption) (*ListKitchensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListKitchensResponse)
	err := c.cc.Invoke(ctx, AuthService_ListKitchens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SearchKitchens(ctx context.Context, in *SearchKitchensRequest, opts ...grpc.CallOption) (*SearchKitchensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchKitchensResponse)
	err := c.cc.Invoke(ctx, AuthService_SearchKitchens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DoesUserExist(ctx context.Context, in *DoesUserExistRequest, opts ...grpc.CallOption) (*DoesUserExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoesUserExistResponse)
	err := c.cc.Invoke(ctx, AuthService_DoesUserExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DoesKitchenExist(ctx context.Context, in *DoesKitchenExistRequest, opts ...grpc.CallOption) (*DoesKitchenExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoesKitchenExistResponse)
	err := c.cc.Invoke(ctx, AuthService_DoesKitchenExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*IsValidTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsValidTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_IsValidToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IncrementTotalOrders(ctx context.Context, in *IncrementTotalOrdersRequest, opts ...grpc.CallOption) (*IncrementTotalOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncrementTotalOrdersResponse)
	err := c.cc.Invoke(ctx, AuthService_IncrementTotalOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IncrementOrderRating(ctx context.Context, in *IncrementOrderRatingRequest, opts ...grpc.CallOption) (*IncrementOrderRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncrementOrderRatingResponse)
	err := c.cc.Invoke(ctx, AuthService_IncrementOrderRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, AuthService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteKitchen(ctx context.Context, in *DeleteKitchenRequest, opts ...grpc.CallOption) (*DeleteKitchenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteKitchenResponse)
	err := c.cc.Invoke(ctx, AuthService_DeleteKitchen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsValidUser(ctx context.Context, in *IsValidUserRequest, opts ...grpc.CallOption) (*IsValidUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsValidUserResponse)
	err := c.cc.Invoke(ctx, AuthService_IsValidUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// 1 Done
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// 2
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// 3 Done
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// 4 Done
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	// 5 Done
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	// 7
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// 8 Done
	CreateKitchen(context.Context, *CreateKitchenRequest) (*CreateKitchenResponse, error)
	// 9 Done
	UpdateKitchen(context.Context, *UpdateKitchenRequest) (*UpdateKitchenResponse, error)
	// 10 Done
	GetKitchen(context.Context, *GetKitchenRequest) (*GetKitchenResponse, error)
	// 11 Done
	ListKitchens(context.Context, *ListKitchensRequest) (*ListKitchensResponse, error)
	// 12 Done
	SearchKitchens(context.Context, *SearchKitchensRequest) (*SearchKitchensResponse, error)
	// Yordamchi funksiyalar
	// 13 Done
	DoesUserExist(context.Context, *DoesUserExistRequest) (*DoesUserExistResponse, error)
	// 14 Done
	DoesKitchenExist(context.Context, *DoesKitchenExistRequest) (*DoesKitchenExistResponse, error)
	// 15
	IsValidToken(context.Context, *IsValidTokenRequest) (*IsValidTokenResponse, error)
	// 16
	IncrementTotalOrders(context.Context, *IncrementTotalOrdersRequest) (*IncrementTotalOrdersResponse, error)
	// 17
	IncrementOrderRating(context.Context, *IncrementOrderRatingRequest) (*IncrementOrderRatingResponse, error)
	// 18
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// 19
	DeleteKitchen(context.Context, *DeleteKitchenRequest) (*DeleteKitchenResponse, error)
	// 20
	IsValidUser(context.Context, *IsValidUserRequest) (*IsValidUserResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) CreateKitchen(context.Context, *CreateKitchenRequest) (*CreateKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKitchen not implemented")
}
func (UnimplementedAuthServiceServer) UpdateKitchen(context.Context, *UpdateKitchenRequest) (*UpdateKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKitchen not implemented")
}
func (UnimplementedAuthServiceServer) GetKitchen(context.Context, *GetKitchenRequest) (*GetKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKitchen not implemented")
}
func (UnimplementedAuthServiceServer) ListKitchens(context.Context, *ListKitchensRequest) (*ListKitchensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKitchens not implemented")
}
func (UnimplementedAuthServiceServer) SearchKitchens(context.Context, *SearchKitchensRequest) (*SearchKitchensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchKitchens not implemented")
}
func (UnimplementedAuthServiceServer) DoesUserExist(context.Context, *DoesUserExistRequest) (*DoesUserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoesUserExist not implemented")
}
func (UnimplementedAuthServiceServer) DoesKitchenExist(context.Context, *DoesKitchenExistRequest) (*DoesKitchenExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoesKitchenExist not implemented")
}
func (UnimplementedAuthServiceServer) IsValidToken(context.Context, *IsValidTokenRequest) (*IsValidTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValidToken not implemented")
}
func (UnimplementedAuthServiceServer) IncrementTotalOrders(context.Context, *IncrementTotalOrdersRequest) (*IncrementTotalOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementTotalOrders not implemented")
}
func (UnimplementedAuthServiceServer) IncrementOrderRating(context.Context, *IncrementOrderRatingRequest) (*IncrementOrderRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementOrderRating not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) DeleteKitchen(context.Context, *DeleteKitchenRequest) (*DeleteKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKitchen not implemented")
}
func (UnimplementedAuthServiceServer) IsValidUser(context.Context, *IsValidUserRequest) (*IsValidUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValidUser not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateKitchen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKitchenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateKitchen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateKitchen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateKitchen(ctx, req.(*CreateKitchenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateKitchen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKitchenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateKitchen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateKitchen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateKitchen(ctx, req.(*UpdateKitchenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetKitchen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKitchenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetKitchen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetKitchen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetKitchen(ctx, req.(*GetKitchenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListKitchens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKitchensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListKitchens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListKitchens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListKitchens(ctx, req.(*ListKitchensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SearchKitchens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchKitchensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SearchKitchens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SearchKitchens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SearchKitchens(ctx, req.(*SearchKitchensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DoesUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoesUserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DoesUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DoesUserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DoesUserExist(ctx, req.(*DoesUserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DoesKitchenExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoesKitchenExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DoesKitchenExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DoesKitchenExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DoesKitchenExist(ctx, req.(*DoesKitchenExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsValidToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsValidToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IsValidToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsValidToken(ctx, req.(*IsValidTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IncrementTotalOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementTotalOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IncrementTotalOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IncrementTotalOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IncrementTotalOrders(ctx, req.(*IncrementTotalOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IncrementOrderRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementOrderRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IncrementOrderRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IncrementOrderRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IncrementOrderRating(ctx, req.(*IncrementOrderRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteKitchen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKitchenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteKitchen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteKitchen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteKitchen(ctx, req.(*DeleteKitchenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsValidUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsValidUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IsValidUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsValidUser(ctx, req.(*IsValidUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _AuthService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _AuthService_UpdateProfile_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "CreateKitchen",
			Handler:    _AuthService_CreateKitchen_Handler,
		},
		{
			MethodName: "UpdateKitchen",
			Handler:    _AuthService_UpdateKitchen_Handler,
		},
		{
			MethodName: "GetKitchen",
			Handler:    _AuthService_GetKitchen_Handler,
		},
		{
			MethodName: "ListKitchens",
			Handler:    _AuthService_ListKitchens_Handler,
		},
		{
			MethodName: "SearchKitchens",
			Handler:    _AuthService_SearchKitchens_Handler,
		},
		{
			MethodName: "DoesUserExist",
			Handler:    _AuthService_DoesUserExist_Handler,
		},
		{
			MethodName: "DoesKitchenExist",
			Handler:    _AuthService_DoesKitchenExist_Handler,
		},
		{
			MethodName: "IsValidToken",
			Handler:    _AuthService_IsValidToken_Handler,
		},
		{
			MethodName: "IncrementTotalOrders",
			Handler:    _AuthService_IncrementTotalOrders_Handler,
		},
		{
			MethodName: "IncrementOrderRating",
			Handler:    _AuthService_IncrementOrderRating_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteKitchen",
			Handler:    _AuthService_DeleteKitchen_Handler,
		},
		{
			MethodName: "IsValidUser",
			Handler:    _AuthService_IsValidUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service/auth-service.proto",
}
