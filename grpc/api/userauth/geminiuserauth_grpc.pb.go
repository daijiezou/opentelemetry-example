// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: proto/geminiuserauth/geminiuserauth.proto

package geminiuserauth

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

// GeminiAuthClient is the client API for GeminiAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeminiAuthClient interface {
	GeminiAuth(ctx context.Context, in *GeminiAuthRequest, opts ...grpc.CallOption) (*GeminiAuthReply, error)
	//  rpc GeminiTokenAuth (GeminiTokenAuthRequest) returns
	//  (GeminiTokenAuthReply) {}
	GeminiLoginAuth(ctx context.Context, in *GeminiLoginAuthRequest, opts ...grpc.CallOption) (*GeminiLoginAuthReplay, error)
	PullImageAuth(ctx context.Context, in *PullImageAuthRequest, opts ...grpc.CallOption) (*PullImageAuthReplay, error)
	PushImageAuth(ctx context.Context, in *PushImageAuthRequest, opts ...grpc.CallOption) (*PushImageAuthReplay, error)
	RefreshLicense(ctx context.Context, in *RefreshLicenseRequest, opts ...grpc.CallOption) (*RefreshLicenseReplay, error)
	GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdReply, error)
	GetUsersByRoleId(ctx context.Context, in *GetUsersByRoleIdRequest, opts ...grpc.CallOption) (*GetUsersByRoleIdReply, error)
	GetRolesInfoById(ctx context.Context, in *GetRolesInfoByIdRequest, opts ...grpc.CallOption) (*GetRolesInfoByIdReply, error)
	CheckKey(ctx context.Context, in *CheckKeyRequest, opts ...grpc.CallOption) (*CheckKeyReply, error)
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyReply, error)
	// 获取用户侧user信息
	UserInfoList(ctx context.Context, in *UserInfoListReq, opts ...grpc.CallOption) (*UserInfoListReply, error)
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginReply, error)
	ReFreshToken(ctx context.Context, in *ReFreshTokenReq, opts ...grpc.CallOption) (*ReFreshTokenReply, error)
}

type geminiAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewGeminiAuthClient(cc grpc.ClientConnInterface) GeminiAuthClient {
	return &geminiAuthClient{cc}
}

func (c *geminiAuthClient) GeminiAuth(ctx context.Context, in *GeminiAuthRequest, opts ...grpc.CallOption) (*GeminiAuthReply, error) {
	out := new(GeminiAuthReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GeminiAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) GeminiLoginAuth(ctx context.Context, in *GeminiLoginAuthRequest, opts ...grpc.CallOption) (*GeminiLoginAuthReplay, error) {
	out := new(GeminiLoginAuthReplay)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GeminiLoginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) PullImageAuth(ctx context.Context, in *PullImageAuthRequest, opts ...grpc.CallOption) (*PullImageAuthReplay, error) {
	out := new(PullImageAuthReplay)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/PullImageAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) PushImageAuth(ctx context.Context, in *PushImageAuthRequest, opts ...grpc.CallOption) (*PushImageAuthReplay, error) {
	out := new(PushImageAuthReplay)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/PushImageAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) RefreshLicense(ctx context.Context, in *RefreshLicenseRequest, opts ...grpc.CallOption) (*RefreshLicenseReplay, error) {
	out := new(RefreshLicenseReplay)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/RefreshLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdReply, error) {
	out := new(GetUserInfoByIdReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GetUserInfoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) GetUsersByRoleId(ctx context.Context, in *GetUsersByRoleIdRequest, opts ...grpc.CallOption) (*GetUsersByRoleIdReply, error) {
	out := new(GetUsersByRoleIdReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GetUsersByRoleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) GetRolesInfoById(ctx context.Context, in *GetRolesInfoByIdRequest, opts ...grpc.CallOption) (*GetRolesInfoByIdReply, error) {
	out := new(GetRolesInfoByIdReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GetRolesInfoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) CheckKey(ctx context.Context, in *CheckKeyRequest, opts ...grpc.CallOption) (*CheckKeyReply, error) {
	out := new(CheckKeyReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/CheckKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyReply, error) {
	out := new(GetKeyReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) UserInfoList(ctx context.Context, in *UserInfoListReq, opts ...grpc.CallOption) (*UserInfoListReply, error) {
	out := new(UserInfoListReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/UserInfoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginReply, error) {
	out := new(UserLoginReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geminiAuthClient) ReFreshToken(ctx context.Context, in *ReFreshTokenReq, opts ...grpc.CallOption) (*ReFreshTokenReply, error) {
	out := new(ReFreshTokenReply)
	err := c.cc.Invoke(ctx, "/geminiuserauth.GeminiAuth/ReFreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeminiAuthServer is the server API for GeminiAuth service.
// All implementations must embed UnimplementedGeminiAuthServer
// for forward compatibility
type GeminiAuthServer interface {
	GeminiAuth(context.Context, *GeminiAuthRequest) (*GeminiAuthReply, error)
	//  rpc GeminiTokenAuth (GeminiTokenAuthRequest) returns
	//  (GeminiTokenAuthReply) {}
	GeminiLoginAuth(context.Context, *GeminiLoginAuthRequest) (*GeminiLoginAuthReplay, error)
	PullImageAuth(context.Context, *PullImageAuthRequest) (*PullImageAuthReplay, error)
	PushImageAuth(context.Context, *PushImageAuthRequest) (*PushImageAuthReplay, error)
	RefreshLicense(context.Context, *RefreshLicenseRequest) (*RefreshLicenseReplay, error)
	GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdReply, error)
	GetUsersByRoleId(context.Context, *GetUsersByRoleIdRequest) (*GetUsersByRoleIdReply, error)
	GetRolesInfoById(context.Context, *GetRolesInfoByIdRequest) (*GetRolesInfoByIdReply, error)
	CheckKey(context.Context, *CheckKeyRequest) (*CheckKeyReply, error)
	GetKey(context.Context, *GetKeyRequest) (*GetKeyReply, error)
	// 获取用户侧user信息
	UserInfoList(context.Context, *UserInfoListReq) (*UserInfoListReply, error)
	UserLogin(context.Context, *UserLoginReq) (*UserLoginReply, error)
	ReFreshToken(context.Context, *ReFreshTokenReq) (*ReFreshTokenReply, error)
	mustEmbedUnimplementedGeminiAuthServer()
}

// UnimplementedGeminiAuthServer must be embedded to have forward compatible implementations.
type UnimplementedGeminiAuthServer struct {
}

func (UnimplementedGeminiAuthServer) GeminiAuth(context.Context, *GeminiAuthRequest) (*GeminiAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeminiAuth not implemented")
}
func (UnimplementedGeminiAuthServer) GeminiLoginAuth(context.Context, *GeminiLoginAuthRequest) (*GeminiLoginAuthReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeminiLoginAuth not implemented")
}
func (UnimplementedGeminiAuthServer) PullImageAuth(context.Context, *PullImageAuthRequest) (*PullImageAuthReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullImageAuth not implemented")
}
func (UnimplementedGeminiAuthServer) PushImageAuth(context.Context, *PushImageAuthRequest) (*PushImageAuthReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushImageAuth not implemented")
}
func (UnimplementedGeminiAuthServer) RefreshLicense(context.Context, *RefreshLicenseRequest) (*RefreshLicenseReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshLicense not implemented")
}
func (UnimplementedGeminiAuthServer) GetUserInfoById(context.Context, *GetUserInfoByIdRequest) (*GetUserInfoByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoById not implemented")
}
func (UnimplementedGeminiAuthServer) GetUsersByRoleId(context.Context, *GetUsersByRoleIdRequest) (*GetUsersByRoleIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByRoleId not implemented")
}
func (UnimplementedGeminiAuthServer) GetRolesInfoById(context.Context, *GetRolesInfoByIdRequest) (*GetRolesInfoByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesInfoById not implemented")
}
func (UnimplementedGeminiAuthServer) CheckKey(context.Context, *CheckKeyRequest) (*CheckKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckKey not implemented")
}
func (UnimplementedGeminiAuthServer) GetKey(context.Context, *GetKeyRequest) (*GetKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedGeminiAuthServer) UserInfoList(context.Context, *UserInfoListReq) (*UserInfoListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoList not implemented")
}
func (UnimplementedGeminiAuthServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedGeminiAuthServer) ReFreshToken(context.Context, *ReFreshTokenReq) (*ReFreshTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReFreshToken not implemented")
}
func (UnimplementedGeminiAuthServer) mustEmbedUnimplementedGeminiAuthServer() {}

// UnsafeGeminiAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeminiAuthServer will
// result in compilation errors.
type UnsafeGeminiAuthServer interface {
	mustEmbedUnimplementedGeminiAuthServer()
}

func RegisterGeminiAuthServer(s grpc.ServiceRegistrar, srv GeminiAuthServer) {
	s.RegisterService(&GeminiAuth_ServiceDesc, srv)
}

func _GeminiAuth_GeminiAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeminiAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GeminiAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GeminiAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GeminiAuth(ctx, req.(*GeminiAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_GeminiLoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeminiLoginAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GeminiLoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GeminiLoginAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GeminiLoginAuth(ctx, req.(*GeminiLoginAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_PullImageAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).PullImageAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/PullImageAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).PullImageAuth(ctx, req.(*PullImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_PushImageAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).PushImageAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/PushImageAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).PushImageAuth(ctx, req.(*PushImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_RefreshLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).RefreshLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/RefreshLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).RefreshLicense(ctx, req.(*RefreshLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GetUserInfoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GetUserInfoById(ctx, req.(*GetUserInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_GetUsersByRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByRoleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GetUsersByRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GetUsersByRoleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GetUsersByRoleId(ctx, req.(*GetUsersByRoleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_GetRolesInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GetRolesInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GetRolesInfoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GetRolesInfoById(ctx, req.(*GetRolesInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_CheckKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).CheckKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/CheckKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).CheckKey(ctx, req.(*CheckKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).GetKey(ctx, req.(*GetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_UserInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).UserInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/UserInfoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).UserInfoList(ctx, req.(*UserInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeminiAuth_ReFreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReFreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeminiAuthServer).ReFreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geminiuserauth.GeminiAuth/ReFreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeminiAuthServer).ReFreshToken(ctx, req.(*ReFreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GeminiAuth_ServiceDesc is the grpc.ServiceDesc for GeminiAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeminiAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geminiuserauth.GeminiAuth",
	HandlerType: (*GeminiAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeminiAuth",
			Handler:    _GeminiAuth_GeminiAuth_Handler,
		},
		{
			MethodName: "GeminiLoginAuth",
			Handler:    _GeminiAuth_GeminiLoginAuth_Handler,
		},
		{
			MethodName: "PullImageAuth",
			Handler:    _GeminiAuth_PullImageAuth_Handler,
		},
		{
			MethodName: "PushImageAuth",
			Handler:    _GeminiAuth_PushImageAuth_Handler,
		},
		{
			MethodName: "RefreshLicense",
			Handler:    _GeminiAuth_RefreshLicense_Handler,
		},
		{
			MethodName: "GetUserInfoById",
			Handler:    _GeminiAuth_GetUserInfoById_Handler,
		},
		{
			MethodName: "GetUsersByRoleId",
			Handler:    _GeminiAuth_GetUsersByRoleId_Handler,
		},
		{
			MethodName: "GetRolesInfoById",
			Handler:    _GeminiAuth_GetRolesInfoById_Handler,
		},
		{
			MethodName: "CheckKey",
			Handler:    _GeminiAuth_CheckKey_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _GeminiAuth_GetKey_Handler,
		},
		{
			MethodName: "UserInfoList",
			Handler:    _GeminiAuth_UserInfoList_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _GeminiAuth_UserLogin_Handler,
		},
		{
			MethodName: "ReFreshToken",
			Handler:    _GeminiAuth_ReFreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/geminiuserauth/geminiuserauth.proto",
}