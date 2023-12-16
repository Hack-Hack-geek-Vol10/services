// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/sql-service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sql_ConvertDDL_FullMethodName = "/sql.v1.Sql/ConvertDDL"
)

// SqlClient is the client API for Sql service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqlClient interface {
	ConvertDDL(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
}

type sqlClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlClient(cc grpc.ClientConnInterface) SqlClient {
	return &sqlClient{cc}
}

func (c *sqlClient) ConvertDDL(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, Sql_ConvertDDL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlServer is the server API for Sql service.
// All implementations must embed UnimplementedSqlServer
// for forward compatibility
type SqlServer interface {
	ConvertDDL(context.Context, *ConvertRequest) (*ConvertResponse, error)
	mustEmbedUnimplementedSqlServer()
}

// UnimplementedSqlServer must be embedded to have forward compatible implementations.
type UnimplementedSqlServer struct {
}

func (UnimplementedSqlServer) ConvertDDL(context.Context, *ConvertRequest) (*ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertDDL not implemented")
}
func (UnimplementedSqlServer) mustEmbedUnimplementedSqlServer() {}

// UnsafeSqlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqlServer will
// result in compilation errors.
type UnsafeSqlServer interface {
	mustEmbedUnimplementedSqlServer()
}

func RegisterSqlServer(s grpc.ServiceRegistrar, srv SqlServer) {
	s.RegisterService(&Sql_ServiceDesc, srv)
}

func _Sql_ConvertDDL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlServer).ConvertDDL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sql_ConvertDDL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlServer).ConvertDDL(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sql_ServiceDesc is the grpc.ServiceDesc for Sql service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sql_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sql.v1.Sql",
	HandlerType: (*SqlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertDDL",
			Handler:    _Sql_ConvertDDL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/sql-service.proto",
}
