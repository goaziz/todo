// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: todo/todo.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ToDoService_CreateToDo_FullMethodName = "/todo.ToDoService/CreateToDo"
	ToDoService_UpdateToDo_FullMethodName = "/todo.ToDoService/UpdateToDo"
	ToDoService_GetToDo_FullMethodName    = "/todo.ToDoService/GetToDo"
	ToDoService_ListToDo_FullMethodName   = "/todo.ToDoService/ListToDo"
	ToDoService_DeleteToDo_FullMethodName = "/todo.ToDoService/DeleteToDo"
)

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for managing to do items
type ToDoServiceClient interface {
	// Create a to do item
	CreateToDo(ctx context.Context, in *CreateToDoRequest, opts ...grpc.CallOption) (*ToDo, error)
	// Update a to do item
	UpdateToDo(ctx context.Context, in *UpdateToDoRequest, opts ...grpc.CallOption) (*ToDo, error)
	// Get a to do item
	GetToDo(ctx context.Context, in *GetToDoRequest, opts ...grpc.CallOption) (*ToDo, error)
	// List all to do items
	ListToDo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListToDoResponse, error)
	// Delete a To Do item
	DeleteToDo(ctx context.Context, in *GetToDoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) CreateToDo(ctx context.Context, in *CreateToDoRequest, opts ...grpc.CallOption) (*ToDo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ToDo)
	err := c.cc.Invoke(ctx, ToDoService_CreateToDo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) UpdateToDo(ctx context.Context, in *UpdateToDoRequest, opts ...grpc.CallOption) (*ToDo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ToDo)
	err := c.cc.Invoke(ctx, ToDoService_UpdateToDo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) GetToDo(ctx context.Context, in *GetToDoRequest, opts ...grpc.CallOption) (*ToDo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ToDo)
	err := c.cc.Invoke(ctx, ToDoService_GetToDo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ListToDo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListToDoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListToDoResponse)
	err := c.cc.Invoke(ctx, ToDoService_ListToDo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) DeleteToDo(ctx context.Context, in *GetToDoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ToDoService_DeleteToDo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
// All implementations must embed UnimplementedToDoServiceServer
// for forward compatibility.
//
// Service for managing to do items
type ToDoServiceServer interface {
	// Create a to do item
	CreateToDo(context.Context, *CreateToDoRequest) (*ToDo, error)
	// Update a to do item
	UpdateToDo(context.Context, *UpdateToDoRequest) (*ToDo, error)
	// Get a to do item
	GetToDo(context.Context, *GetToDoRequest) (*ToDo, error)
	// List all to do items
	ListToDo(context.Context, *emptypb.Empty) (*ListToDoResponse, error)
	// Delete a To Do item
	DeleteToDo(context.Context, *GetToDoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedToDoServiceServer()
}

// UnimplementedToDoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedToDoServiceServer struct{}

func (UnimplementedToDoServiceServer) CreateToDo(context.Context, *CreateToDoRequest) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToDo not implemented")
}
func (UnimplementedToDoServiceServer) UpdateToDo(context.Context, *UpdateToDoRequest) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToDo not implemented")
}
func (UnimplementedToDoServiceServer) GetToDo(context.Context, *GetToDoRequest) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToDo not implemented")
}
func (UnimplementedToDoServiceServer) ListToDo(context.Context, *emptypb.Empty) (*ListToDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListToDo not implemented")
}
func (UnimplementedToDoServiceServer) DeleteToDo(context.Context, *GetToDoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToDo not implemented")
}
func (UnimplementedToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}
func (UnimplementedToDoServiceServer) testEmbeddedByValue()                     {}

// UnsafeToDoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoServiceServer will
// result in compilation errors.
type UnsafeToDoServiceServer interface {
	mustEmbedUnimplementedToDoServiceServer()
}

func RegisterToDoServiceServer(s grpc.ServiceRegistrar, srv ToDoServiceServer) {
	// If the following call pancis, it indicates UnimplementedToDoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ToDoService_ServiceDesc, srv)
}

func _ToDoService_CreateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).CreateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_CreateToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).CreateToDo(ctx, req.(*CreateToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_UpdateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).UpdateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_UpdateToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).UpdateToDo(ctx, req.(*UpdateToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_GetToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).GetToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_GetToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).GetToDo(ctx, req.(*GetToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ListToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ListToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_ListToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ListToDo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_DeleteToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).DeleteToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_DeleteToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).DeleteToDo(ctx, req.(*GetToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoService_ServiceDesc is the grpc.ServiceDesc for ToDoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToDo",
			Handler:    _ToDoService_CreateToDo_Handler,
		},
		{
			MethodName: "UpdateToDo",
			Handler:    _ToDoService_UpdateToDo_Handler,
		},
		{
			MethodName: "GetToDo",
			Handler:    _ToDoService_GetToDo_Handler,
		},
		{
			MethodName: "ListToDo",
			Handler:    _ToDoService_ListToDo_Handler,
		},
		{
			MethodName: "DeleteToDo",
			Handler:    _ToDoService_DeleteToDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/todo.proto",
}