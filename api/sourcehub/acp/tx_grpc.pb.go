// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sourcehub/acp/tx.proto

package acp

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
	Msg_UpdateParams_FullMethodName    = "/sourcehub.acp.Msg/UpdateParams"
	Msg_CreatePolicy_FullMethodName    = "/sourcehub.acp.Msg/CreatePolicy"
	Msg_CheckAccess_FullMethodName     = "/sourcehub.acp.Msg/CheckAccess"
	Msg_SignedPolicyCmd_FullMethodName = "/sourcehub.acp.Msg/SignedPolicyCmd"
	Msg_BearerPolicyCmd_FullMethodName = "/sourcehub.acp.Msg/BearerPolicyCmd"
	Msg_DirectPolicyCmd_FullMethodName = "/sourcehub.acp.Msg/DirectPolicyCmd"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// CreatePolicy adds a new Policy to SourceHub.
	// The Policy models an aplication's high level access control rules.
	CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error)
	// CheckAccess executes an Access Request for an User and stores the result of the evaluation in SourceHub.
	// The resulting evaluation is used to generate a cryptographic proof that the given Access Request
	// was valid at a particular block height.
	CheckAccess(ctx context.Context, in *MsgCheckAccess, opts ...grpc.CallOption) (*MsgCheckAccessResponse, error)
	// SignedPolicyCmd is a wrapper for a Command which is executed within the Context of a Policy.
	// The Command is signed by the Actor issuing it.
	SignedPolicyCmd(ctx context.Context, in *MsgSignedPolicyCmd, opts ...grpc.CallOption) (*MsgSignedPolicyCmdResponse, error)
	// The Msg authenticates the actor initiating the command through a Bearer token.
	// This token MUST be issued and signed by some DID Actor, the verification of the signature
	// is used as authentication proof.
	// Lastly, the Bearer token MUST be bound to some SourceHub account.
	BearerPolicyCmd(ctx context.Context, in *MsgBearerPolicyCmd, opts ...grpc.CallOption) (*MsgBearerPolicyCmdResponse, error)
	DirectPolicyCmd(ctx context.Context, in *MsgDirectPolicyCmd, opts ...grpc.CallOption) (*MsgDirectPolicyCmdResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error) {
	out := new(MsgCreatePolicyResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CheckAccess(ctx context.Context, in *MsgCheckAccess, opts ...grpc.CallOption) (*MsgCheckAccessResponse, error) {
	out := new(MsgCheckAccessResponse)
	err := c.cc.Invoke(ctx, Msg_CheckAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SignedPolicyCmd(ctx context.Context, in *MsgSignedPolicyCmd, opts ...grpc.CallOption) (*MsgSignedPolicyCmdResponse, error) {
	out := new(MsgSignedPolicyCmdResponse)
	err := c.cc.Invoke(ctx, Msg_SignedPolicyCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BearerPolicyCmd(ctx context.Context, in *MsgBearerPolicyCmd, opts ...grpc.CallOption) (*MsgBearerPolicyCmdResponse, error) {
	out := new(MsgBearerPolicyCmdResponse)
	err := c.cc.Invoke(ctx, Msg_BearerPolicyCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DirectPolicyCmd(ctx context.Context, in *MsgDirectPolicyCmd, opts ...grpc.CallOption) (*MsgDirectPolicyCmdResponse, error) {
	out := new(MsgDirectPolicyCmdResponse)
	err := c.cc.Invoke(ctx, Msg_DirectPolicyCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// CreatePolicy adds a new Policy to SourceHub.
	// The Policy models an aplication's high level access control rules.
	CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error)
	// CheckAccess executes an Access Request for an User and stores the result of the evaluation in SourceHub.
	// The resulting evaluation is used to generate a cryptographic proof that the given Access Request
	// was valid at a particular block height.
	CheckAccess(context.Context, *MsgCheckAccess) (*MsgCheckAccessResponse, error)
	// SignedPolicyCmd is a wrapper for a Command which is executed within the Context of a Policy.
	// The Command is signed by the Actor issuing it.
	SignedPolicyCmd(context.Context, *MsgSignedPolicyCmd) (*MsgSignedPolicyCmdResponse, error)
	// The Msg authenticates the actor initiating the command through a Bearer token.
	// This token MUST be issued and signed by some DID Actor, the verification of the signature
	// is used as authentication proof.
	// Lastly, the Bearer token MUST be bound to some SourceHub account.
	BearerPolicyCmd(context.Context, *MsgBearerPolicyCmd) (*MsgBearerPolicyCmdResponse, error)
	DirectPolicyCmd(context.Context, *MsgDirectPolicyCmd) (*MsgDirectPolicyCmdResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedMsgServer) CheckAccess(context.Context, *MsgCheckAccess) (*MsgCheckAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccess not implemented")
}
func (UnimplementedMsgServer) SignedPolicyCmd(context.Context, *MsgSignedPolicyCmd) (*MsgSignedPolicyCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignedPolicyCmd not implemented")
}
func (UnimplementedMsgServer) BearerPolicyCmd(context.Context, *MsgBearerPolicyCmd) (*MsgBearerPolicyCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BearerPolicyCmd not implemented")
}
func (UnimplementedMsgServer) DirectPolicyCmd(context.Context, *MsgDirectPolicyCmd) (*MsgDirectPolicyCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectPolicyCmd not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePolicy(ctx, req.(*MsgCreatePolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CheckAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCheckAccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CheckAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CheckAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CheckAccess(ctx, req.(*MsgCheckAccess))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SignedPolicyCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignedPolicyCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignedPolicyCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SignedPolicyCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignedPolicyCmd(ctx, req.(*MsgSignedPolicyCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BearerPolicyCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBearerPolicyCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BearerPolicyCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BearerPolicyCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BearerPolicyCmd(ctx, req.(*MsgBearerPolicyCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DirectPolicyCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDirectPolicyCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DirectPolicyCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DirectPolicyCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DirectPolicyCmd(ctx, req.(*MsgDirectPolicyCmd))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sourcehub.acp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Msg_CreatePolicy_Handler,
		},
		{
			MethodName: "CheckAccess",
			Handler:    _Msg_CheckAccess_Handler,
		},
		{
			MethodName: "SignedPolicyCmd",
			Handler:    _Msg_SignedPolicyCmd_Handler,
		},
		{
			MethodName: "BearerPolicyCmd",
			Handler:    _Msg_BearerPolicyCmd_Handler,
		},
		{
			MethodName: "DirectPolicyCmd",
			Handler:    _Msg_DirectPolicyCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sourcehub/acp/tx.proto",
}
