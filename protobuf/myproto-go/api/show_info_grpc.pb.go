package api

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

// ShowServiceClient is the client API for ShowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShowServiceClient interface {
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowInfo, error)
}

type showServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShowServiceClient(cc grpc.ClientConnInterface) ShowServiceClient {
	return &showServiceClient{cc}
}

func (c *showServiceClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowInfo, error) {
	out := new(ShowInfo)
	err := c.cc.Invoke(ctx, "/api.ShowService/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShowServiceServer is the server API for ShowService service.
// All implementations must embed UnimplementedShowServiceServer
// for forward compatibility
type ShowServiceServer interface {
	Show(context.Context, *ShowRequest) (*ShowInfo, error)
	mustEmbedUnimplementedShowServiceServer()
}

// UnimplementedShowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShowServiceServer struct {
}

func (UnimplementedShowServiceServer) Show(context.Context, *ShowRequest) (*ShowInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedShowServiceServer) mustEmbedUnimplementedShowServiceServer() {}

// UnsafeShowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShowServiceServer will
// result in compilation errors.
type UnsafeShowServiceServer interface {
	mustEmbedUnimplementedShowServiceServer()
}

func RegisterShowServiceServer(s grpc.ServiceRegistrar, srv ShowServiceServer) {
	s.RegisterService(&ShowService_ServiceDesc, srv)
}

func _ShowService_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShowServiceServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShowService/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShowServiceServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShowService_ServiceDesc is the grpc.ServiceDesc for ShowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShowService",
	HandlerType: (*ShowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _ShowService_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/show_info.proto",
}
