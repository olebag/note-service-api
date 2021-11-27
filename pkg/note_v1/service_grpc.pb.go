// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package note_v1

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

// NoteV1Client is the client API for NoteV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteV1Client interface {
	// Adds a note to the database, giving it a specific Id.
	CreateNoteV1(ctx context.Context, in *CreateNoteV1Request, opts ...grpc.CallOption) (*CreateNoteV1Response, error)
}

type noteV1Client struct {
	cc grpc.ClientConnInterface
}

func NewNoteV1Client(cc grpc.ClientConnInterface) NoteV1Client {
	return &noteV1Client{cc}
}

func (c *noteV1Client) CreateNoteV1(ctx context.Context, in *CreateNoteV1Request, opts ...grpc.CallOption) (*CreateNoteV1Response, error) {
	out := new(CreateNoteV1Response)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/CreateNoteV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteV1Server is the server API for NoteV1 service.
// All implementations must embed UnimplementedNoteV1Server
// for forward compatibility
type NoteV1Server interface {
	// Adds a note to the database, giving it a specific Id.
	CreateNoteV1(context.Context, *CreateNoteV1Request) (*CreateNoteV1Response, error)
	mustEmbedUnimplementedNoteV1Server()
}

// UnimplementedNoteV1Server must be embedded to have forward compatible implementations.
type UnimplementedNoteV1Server struct {
}

func (UnimplementedNoteV1Server) CreateNoteV1(context.Context, *CreateNoteV1Request) (*CreateNoteV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNoteV1 not implemented")
}
func (UnimplementedNoteV1Server) mustEmbedUnimplementedNoteV1Server() {}

// UnsafeNoteV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteV1Server will
// result in compilation errors.
type UnsafeNoteV1Server interface {
	mustEmbedUnimplementedNoteV1Server()
}

func RegisterNoteV1Server(s grpc.ServiceRegistrar, srv NoteV1Server) {
	s.RegisterService(&NoteV1_ServiceDesc, srv)
}

func _NoteV1_CreateNoteV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).CreateNoteV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/CreateNoteV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).CreateNoteV1(ctx, req.(*CreateNoteV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteV1_ServiceDesc is the grpc.ServiceDesc for NoteV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.note_v1.NoteV1",
	HandlerType: (*NoteV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNoteV1",
			Handler:    _NoteV1_CreateNoteV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
