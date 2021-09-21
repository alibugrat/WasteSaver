// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_product_grpc

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

// ProductManagementClient is the client API for ProductManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManagementClient interface {
	//Unary
	GetProducts(ctx context.Context, in *GetProductParams, opts ...grpc.CallOption) (*ProductList, error)
	GetProduct(ctx context.Context, in *GetProductParamsWithId, opts ...grpc.CallOption) (*Product, error)
}

type productManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManagementClient(cc grpc.ClientConnInterface) ProductManagementClient {
	return &productManagementClient{cc}
}

func (c *productManagementClient) GetProducts(ctx context.Context, in *GetProductParams, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/product.ProductManagement/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementClient) GetProduct(ctx context.Context, in *GetProductParamsWithId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.ProductManagement/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductManagementServer is the servers API for ProductManagement service.
// All implementations must embed UnimplementedProductManagementServer
// for forward compatibility
type ProductManagementServer interface {
	//Unary
	GetProducts(context.Context, *GetProductParams) (*ProductList, error)
	GetProduct(context.Context, *GetProductParamsWithId) (*Product, error)
	mustEmbedUnimplementedProductManagementServer()
}

// UnimplementedProductManagementServer must be embedded to have forward compatible implementations.
type UnimplementedProductManagementServer struct {
}

func (UnimplementedProductManagementServer) GetProducts(context.Context, *GetProductParams) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductManagementServer) GetProduct(context.Context, *GetProductParamsWithId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductManagementServer) mustEmbedUnimplementedProductManagementServer() {}

// UnsafeProductManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManagementServer will
// result in compilation errors.
type UnsafeProductManagementServer interface {
	mustEmbedUnimplementedProductManagementServer()
}

func RegisterProductManagementServer(s grpc.ServiceRegistrar, srv ProductManagementServer) {
	s.RegisterService(&ProductManagement_ServiceDesc, srv)
}

func _ProductManagement_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagement/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServer).GetProducts(ctx, req.(*GetProductParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagement_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductParamsWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagement/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServer).GetProduct(ctx, req.(*GetProductParamsWithId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductManagement_ServiceDesc is the grpc.ServiceDesc for ProductManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductManagement",
	HandlerType: (*ProductManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductManagement_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductManagement_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}