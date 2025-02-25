// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: protobuf/proto/agent.proto

package agent

import (
	context "context"
	common "github.com/dyrector-io/dyrectorio/protobuf/go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	//*
	// Subscribe with pre-assigned AgentID, waiting for incoming
	// deploy requests and prefix status requests.
	// In both cases, separate, shorter-living channels are opened.
	// For deployment status reports, closed when ended.
	// For prefix state reports, should be closed by the server.
	Connect(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (Agent_ConnectClient, error)
	DeploymentStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_DeploymentStatusClient, error)
	ContainerState(ctx context.Context, opts ...grpc.CallOption) (Agent_ContainerStateClient, error)
	SecretsList(ctx context.Context, in *common.ListSecretsResponse, opts ...grpc.CallOption) (*Empty, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Connect(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (Agent_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/agent.Agent/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_ConnectClient interface {
	Recv() (*AgentCommand, error)
	grpc.ClientStream
}

type agentConnectClient struct {
	grpc.ClientStream
}

func (x *agentConnectClient) Recv() (*AgentCommand, error) {
	m := new(AgentCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) DeploymentStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_DeploymentStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[1], "/agent.Agent/DeploymentStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentDeploymentStatusClient{stream}
	return x, nil
}

type Agent_DeploymentStatusClient interface {
	Send(*common.DeploymentStatusMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type agentDeploymentStatusClient struct {
	grpc.ClientStream
}

func (x *agentDeploymentStatusClient) Send(m *common.DeploymentStatusMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentDeploymentStatusClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) ContainerState(ctx context.Context, opts ...grpc.CallOption) (Agent_ContainerStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[2], "/agent.Agent/ContainerState", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentContainerStateClient{stream}
	return x, nil
}

type Agent_ContainerStateClient interface {
	Send(*common.ContainerStateListMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type agentContainerStateClient struct {
	grpc.ClientStream
}

func (x *agentContainerStateClient) Send(m *common.ContainerStateListMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentContainerStateClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) SecretsList(ctx context.Context, in *common.ListSecretsResponse, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agent.Agent/SecretsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	//*
	// Subscribe with pre-assigned AgentID, waiting for incoming
	// deploy requests and prefix status requests.
	// In both cases, separate, shorter-living channels are opened.
	// For deployment status reports, closed when ended.
	// For prefix state reports, should be closed by the server.
	Connect(*AgentInfo, Agent_ConnectServer) error
	DeploymentStatus(Agent_DeploymentStatusServer) error
	ContainerState(Agent_ContainerStateServer) error
	SecretsList(context.Context, *common.ListSecretsResponse) (*Empty, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Connect(*AgentInfo, Agent_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedAgentServer) DeploymentStatus(Agent_DeploymentStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method DeploymentStatus not implemented")
}
func (UnimplementedAgentServer) ContainerState(Agent_ContainerStateServer) error {
	return status.Errorf(codes.Unimplemented, "method ContainerState not implemented")
}
func (UnimplementedAgentServer) SecretsList(context.Context, *common.ListSecretsResponse) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecretsList not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).Connect(m, &agentConnectServer{stream})
}

type Agent_ConnectServer interface {
	Send(*AgentCommand) error
	grpc.ServerStream
}

type agentConnectServer struct {
	grpc.ServerStream
}

func (x *agentConnectServer) Send(m *AgentCommand) error {
	return x.ServerStream.SendMsg(m)
}

func _Agent_DeploymentStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).DeploymentStatus(&agentDeploymentStatusServer{stream})
}

type Agent_DeploymentStatusServer interface {
	SendAndClose(*Empty) error
	Recv() (*common.DeploymentStatusMessage, error)
	grpc.ServerStream
}

type agentDeploymentStatusServer struct {
	grpc.ServerStream
}

func (x *agentDeploymentStatusServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentDeploymentStatusServer) Recv() (*common.DeploymentStatusMessage, error) {
	m := new(common.DeploymentStatusMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_ContainerState_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).ContainerState(&agentContainerStateServer{stream})
}

type Agent_ContainerStateServer interface {
	SendAndClose(*Empty) error
	Recv() (*common.ContainerStateListMessage, error)
	grpc.ServerStream
}

type agentContainerStateServer struct {
	grpc.ServerStream
}

func (x *agentContainerStateServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentContainerStateServer) Recv() (*common.ContainerStateListMessage, error) {
	m := new(common.ContainerStateListMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_SecretsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ListSecretsResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SecretsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/SecretsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SecretsList(ctx, req.(*common.ListSecretsResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SecretsList",
			Handler:    _Agent_SecretsList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Agent_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeploymentStatus",
			Handler:       _Agent_DeploymentStatus_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ContainerState",
			Handler:       _Agent_ContainerState_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/proto/agent.proto",
}
