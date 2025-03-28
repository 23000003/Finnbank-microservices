// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: account.proto

package account

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountService_GetAccounts_FullMethodName       = "/AccountService/GetAccounts"
	AccountService_GetAccountById_FullMethodName    = "/AccountService/GetAccountById"
	AccountService_GetAccountByEmail_FullMethodName = "/AccountService/GetAccountByEmail"
	AccountService_UpdateAccount_FullMethodName     = "/AccountService/UpdateAccount"
	AccountService_UpdateCardStatus_FullMethodName  = "/AccountService/UpdateCardStatus"
	AccountService_UpdatePassword_FullMethodName    = "/AccountService/UpdatePassword"
	AccountService_DeleteAccount_FullMethodName     = "/AccountService/DeleteAccount"
	AccountService_AddAccount_FullMethodName        = "/AccountService/AddAccount"
	AccountService_LoginUser_FullMethodName         = "/AccountService/LoginUser"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetAccounts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccountsResponse, error)
	GetAccountById(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetAccountByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	UpdateCardStatus(ctx context.Context, in *CardUpdateRequest, opts ...grpc.CallOption) (*CardUpdateResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountById(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_UpdateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateCardStatus(ctx context.Context, in *CardUpdateRequest, opts ...grpc.CallOption) (*CardUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CardUpdateResponse)
	err := c.cc.Invoke(ctx, AccountService_UpdateCardStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, AccountService_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_AddAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AccountService_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility.
type AccountServiceServer interface {
	GetAccounts(context.Context, *Empty) (*AccountsResponse, error)
	GetAccountById(context.Context, *AccountRequest) (*AccountResponse, error)
	GetAccountByEmail(context.Context, *EmailRequest) (*AccountResponse, error)
	UpdateAccount(context.Context, *UpdateRequest) (*AccountResponse, error)
	UpdateCardStatus(context.Context, *CardUpdateRequest) (*CardUpdateResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	DeleteAccount(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountServiceServer struct{}

func (UnimplementedAccountServiceServer) GetAccounts(context.Context, *Empty) (*AccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountById(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByEmail(context.Context, *EmailRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByEmail not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *UpdateRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateCardStatus(context.Context, *CardUpdateRequest) (*CardUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardStatus not implemented")
}
func (UnimplementedAccountServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountServiceServer) AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedAccountServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}
func (UnimplementedAccountServiceServer) testEmbeddedByValue()                        {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccounts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountById(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateCardStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateCardStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdateCardStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateCardStatus(ctx, req.(*CardUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_AddAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).AddAccount(ctx, req.(*AddAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccounts",
			Handler:    _AccountService_GetAccounts_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _AccountService_GetAccountById_Handler,
		},
		{
			MethodName: "GetAccountByEmail",
			Handler:    _AccountService_GetAccountByEmail_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "UpdateCardStatus",
			Handler:    _AccountService_UpdateCardStatus_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _AccountService_UpdatePassword_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _AccountService_AddAccount_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AccountService_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
