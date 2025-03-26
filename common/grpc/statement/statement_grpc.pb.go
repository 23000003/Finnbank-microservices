// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: statement.proto

package statement

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
	StatementService_AddStatement_FullMethodName = "/StatementService/AddStatement"
	StatementService_GetStatement_FullMethodName = "/StatementService/GetStatement"
)

// StatementServiceClient is the client API for StatementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatementServiceClient interface {
	AddStatement(ctx context.Context, in *AddStatementRequest, opts ...grpc.CallOption) (*AddStatementResponse, error)
	GetStatement(ctx context.Context, in *GetStatementRequest, opts ...grpc.CallOption) (*GetStatementResponse, error)
}

type statementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatementServiceClient(cc grpc.ClientConnInterface) StatementServiceClient {
	return &statementServiceClient{cc}
}

func (c *statementServiceClient) AddStatement(ctx context.Context, in *AddStatementRequest, opts ...grpc.CallOption) (*AddStatementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddStatementResponse)
	err := c.cc.Invoke(ctx, StatementService_AddStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) GetStatement(ctx context.Context, in *GetStatementRequest, opts ...grpc.CallOption) (*GetStatementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatementResponse)
	err := c.cc.Invoke(ctx, StatementService_GetStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatementServiceServer is the server API for StatementService service.
// All implementations must embed UnimplementedStatementServiceServer
// for forward compatibility.
type StatementServiceServer interface {
	AddStatement(context.Context, *AddStatementRequest) (*AddStatementResponse, error)
	GetStatement(context.Context, *GetStatementRequest) (*GetStatementResponse, error)
	mustEmbedUnimplementedStatementServiceServer()
}

// UnimplementedStatementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatementServiceServer struct{}

func (UnimplementedStatementServiceServer) AddStatement(context.Context, *AddStatementRequest) (*AddStatementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStatement not implemented")
}
func (UnimplementedStatementServiceServer) GetStatement(context.Context, *GetStatementRequest) (*GetStatementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatement not implemented")
}
func (UnimplementedStatementServiceServer) mustEmbedUnimplementedStatementServiceServer() {}
func (UnimplementedStatementServiceServer) testEmbeddedByValue()                          {}

// UnsafeStatementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatementServiceServer will
// result in compilation errors.
type UnsafeStatementServiceServer interface {
	mustEmbedUnimplementedStatementServiceServer()
}

func RegisterStatementServiceServer(s grpc.ServiceRegistrar, srv StatementServiceServer) {
	// If the following call pancis, it indicates UnimplementedStatementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatementService_ServiceDesc, srv)
}

func _StatementService_AddStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).AddStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_AddStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).AddStatement(ctx, req.(*AddStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_GetStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).GetStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_GetStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).GetStatement(ctx, req.(*GetStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatementService_ServiceDesc is the grpc.ServiceDesc for StatementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StatementService",
	HandlerType: (*StatementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStatement",
			Handler:    _StatementService_AddStatement_Handler,
		},
		{
			MethodName: "GetStatement",
			Handler:    _StatementService_GetStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statement.proto",
}
