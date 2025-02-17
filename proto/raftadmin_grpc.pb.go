// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: raftadmin.proto

package proto

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
	RaftAdmin_AddNonvoter_FullMethodName                = "/RaftAdmin/AddNonvoter"
	RaftAdmin_AddVoter_FullMethodName                   = "/RaftAdmin/AddVoter"
	RaftAdmin_AppliedIndex_FullMethodName               = "/RaftAdmin/AppliedIndex"
	RaftAdmin_ApplyLog_FullMethodName                   = "/RaftAdmin/ApplyLog"
	RaftAdmin_Barrier_FullMethodName                    = "/RaftAdmin/Barrier"
	RaftAdmin_DemoteVoter_FullMethodName                = "/RaftAdmin/DemoteVoter"
	RaftAdmin_GetConfiguration_FullMethodName           = "/RaftAdmin/GetConfiguration"
	RaftAdmin_LastContact_FullMethodName                = "/RaftAdmin/LastContact"
	RaftAdmin_LastIndex_FullMethodName                  = "/RaftAdmin/LastIndex"
	RaftAdmin_Leader_FullMethodName                     = "/RaftAdmin/Leader"
	RaftAdmin_LeadershipTransfer_FullMethodName         = "/RaftAdmin/LeadershipTransfer"
	RaftAdmin_LeadershipTransferToServer_FullMethodName = "/RaftAdmin/LeadershipTransferToServer"
	RaftAdmin_RemoveServer_FullMethodName               = "/RaftAdmin/RemoveServer"
	RaftAdmin_Shutdown_FullMethodName                   = "/RaftAdmin/Shutdown"
	RaftAdmin_Snapshot_FullMethodName                   = "/RaftAdmin/Snapshot"
	RaftAdmin_State_FullMethodName                      = "/RaftAdmin/State"
	RaftAdmin_Stats_FullMethodName                      = "/RaftAdmin/Stats"
	RaftAdmin_VerifyLeader_FullMethodName               = "/RaftAdmin/VerifyLeader"
	RaftAdmin_Await_FullMethodName                      = "/RaftAdmin/Await"
	RaftAdmin_Forget_FullMethodName                     = "/RaftAdmin/Forget"
)

// RaftAdminClient is the client API for RaftAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftAdminClient interface {
	AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*Future, error)
	AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*Future, error)
	AppliedIndex(ctx context.Context, in *AppliedIndexRequest, opts ...grpc.CallOption) (*AppliedIndexResponse, error)
	ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*Future, error)
	Barrier(ctx context.Context, in *BarrierRequest, opts ...grpc.CallOption) (*Future, error)
	DemoteVoter(ctx context.Context, in *DemoteVoterRequest, opts ...grpc.CallOption) (*Future, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	LastContact(ctx context.Context, in *LastContactRequest, opts ...grpc.CallOption) (*LastContactResponse, error)
	LastIndex(ctx context.Context, in *LastIndexRequest, opts ...grpc.CallOption) (*LastIndexResponse, error)
	Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error)
	LeadershipTransfer(ctx context.Context, in *LeadershipTransferRequest, opts ...grpc.CallOption) (*Future, error)
	LeadershipTransferToServer(ctx context.Context, in *LeadershipTransferToServerRequest, opts ...grpc.CallOption) (*Future, error)
	RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*Future, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*Future, error)
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*Future, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	VerifyLeader(ctx context.Context, in *VerifyLeaderRequest, opts ...grpc.CallOption) (*Future, error)
	Await(ctx context.Context, in *Future, opts ...grpc.CallOption) (*AwaitResponse, error)
	Forget(ctx context.Context, in *Future, opts ...grpc.CallOption) (*ForgetResponse, error)
}

type raftAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftAdminClient(cc grpc.ClientConnInterface) RaftAdminClient {
	return &raftAdminClient{cc}
}

func (c *raftAdminClient) AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_AddNonvoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_AddVoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) AppliedIndex(ctx context.Context, in *AppliedIndexRequest, opts ...grpc.CallOption) (*AppliedIndexResponse, error) {
	out := new(AppliedIndexResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_AppliedIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_ApplyLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Barrier(ctx context.Context, in *BarrierRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_Barrier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) DemoteVoter(ctx context.Context, in *DemoteVoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_DemoteVoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_GetConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) LastContact(ctx context.Context, in *LastContactRequest, opts ...grpc.CallOption) (*LastContactResponse, error) {
	out := new(LastContactResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_LastContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) LastIndex(ctx context.Context, in *LastIndexRequest, opts ...grpc.CallOption) (*LastIndexResponse, error) {
	out := new(LastIndexResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_LastIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error) {
	out := new(LeaderResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_Leader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) LeadershipTransfer(ctx context.Context, in *LeadershipTransferRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_LeadershipTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) LeadershipTransferToServer(ctx context.Context, in *LeadershipTransferToServerRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_LeadershipTransferToServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_RemoveServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_Shutdown_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_Snapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_State_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_Stats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) VerifyLeader(ctx context.Context, in *VerifyLeaderRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftAdmin_VerifyLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Await(ctx context.Context, in *Future, opts ...grpc.CallOption) (*AwaitResponse, error) {
	out := new(AwaitResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_Await_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftAdminClient) Forget(ctx context.Context, in *Future, opts ...grpc.CallOption) (*ForgetResponse, error) {
	out := new(ForgetResponse)
	err := c.cc.Invoke(ctx, RaftAdmin_Forget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftAdminServer is the server API for RaftAdmin service.
// All implementations must embed UnimplementedRaftAdminServer
// for forward compatibility
type RaftAdminServer interface {
	AddNonvoter(context.Context, *AddNonvoterRequest) (*Future, error)
	AddVoter(context.Context, *AddVoterRequest) (*Future, error)
	AppliedIndex(context.Context, *AppliedIndexRequest) (*AppliedIndexResponse, error)
	ApplyLog(context.Context, *ApplyLogRequest) (*Future, error)
	Barrier(context.Context, *BarrierRequest) (*Future, error)
	DemoteVoter(context.Context, *DemoteVoterRequest) (*Future, error)
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	LastContact(context.Context, *LastContactRequest) (*LastContactResponse, error)
	LastIndex(context.Context, *LastIndexRequest) (*LastIndexResponse, error)
	Leader(context.Context, *LeaderRequest) (*LeaderResponse, error)
	LeadershipTransfer(context.Context, *LeadershipTransferRequest) (*Future, error)
	LeadershipTransferToServer(context.Context, *LeadershipTransferToServerRequest) (*Future, error)
	RemoveServer(context.Context, *RemoveServerRequest) (*Future, error)
	Shutdown(context.Context, *ShutdownRequest) (*Future, error)
	Snapshot(context.Context, *SnapshotRequest) (*Future, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	VerifyLeader(context.Context, *VerifyLeaderRequest) (*Future, error)
	Await(context.Context, *Future) (*AwaitResponse, error)
	Forget(context.Context, *Future) (*ForgetResponse, error)
	mustEmbedUnimplementedRaftAdminServer()
}

// UnimplementedRaftAdminServer must be embedded to have forward compatible implementations.
type UnimplementedRaftAdminServer struct {
}

func (UnimplementedRaftAdminServer) AddNonvoter(context.Context, *AddNonvoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNonvoter not implemented")
}
func (UnimplementedRaftAdminServer) AddVoter(context.Context, *AddVoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVoter not implemented")
}
func (UnimplementedRaftAdminServer) AppliedIndex(context.Context, *AppliedIndexRequest) (*AppliedIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedIndex not implemented")
}
func (UnimplementedRaftAdminServer) ApplyLog(context.Context, *ApplyLogRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLog not implemented")
}
func (UnimplementedRaftAdminServer) Barrier(context.Context, *BarrierRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Barrier not implemented")
}
func (UnimplementedRaftAdminServer) DemoteVoter(context.Context, *DemoteVoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteVoter not implemented")
}
func (UnimplementedRaftAdminServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedRaftAdminServer) LastContact(context.Context, *LastContactRequest) (*LastContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastContact not implemented")
}
func (UnimplementedRaftAdminServer) LastIndex(context.Context, *LastIndexRequest) (*LastIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastIndex not implemented")
}
func (UnimplementedRaftAdminServer) Leader(context.Context, *LeaderRequest) (*LeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leader not implemented")
}
func (UnimplementedRaftAdminServer) LeadershipTransfer(context.Context, *LeadershipTransferRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeadershipTransfer not implemented")
}
func (UnimplementedRaftAdminServer) LeadershipTransferToServer(context.Context, *LeadershipTransferToServerRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeadershipTransferToServer not implemented")
}
func (UnimplementedRaftAdminServer) RemoveServer(context.Context, *RemoveServerRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServer not implemented")
}
func (UnimplementedRaftAdminServer) Shutdown(context.Context, *ShutdownRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedRaftAdminServer) Snapshot(context.Context, *SnapshotRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedRaftAdminServer) State(context.Context, *StateRequest) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedRaftAdminServer) Stats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedRaftAdminServer) VerifyLeader(context.Context, *VerifyLeaderRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyLeader not implemented")
}
func (UnimplementedRaftAdminServer) Await(context.Context, *Future) (*AwaitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Await not implemented")
}
func (UnimplementedRaftAdminServer) Forget(context.Context, *Future) (*ForgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forget not implemented")
}
func (UnimplementedRaftAdminServer) mustEmbedUnimplementedRaftAdminServer() {}

// UnsafeRaftAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftAdminServer will
// result in compilation errors.
type UnsafeRaftAdminServer interface {
	mustEmbedUnimplementedRaftAdminServer()
}

func RegisterRaftAdminServer(s grpc.ServiceRegistrar, srv RaftAdminServer) {
	s.RegisterService(&RaftAdmin_ServiceDesc, srv)
}

func _RaftAdmin_AddNonvoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNonvoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).AddNonvoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_AddNonvoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).AddNonvoter(ctx, req.(*AddNonvoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_AddVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).AddVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_AddVoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).AddVoter(ctx, req.(*AddVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_AppliedIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppliedIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).AppliedIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_AppliedIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).AppliedIndex(ctx, req.(*AppliedIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_ApplyLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).ApplyLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_ApplyLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).ApplyLog(ctx, req.(*ApplyLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Barrier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarrierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Barrier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Barrier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Barrier(ctx, req.(*BarrierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_DemoteVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoteVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).DemoteVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_DemoteVoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).DemoteVoter(ctx, req.(*DemoteVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_GetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_LastContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).LastContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_LastContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).LastContact(ctx, req.(*LastContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_LastIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).LastIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_LastIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).LastIndex(ctx, req.(*LastIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Leader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Leader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Leader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Leader(ctx, req.(*LeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_LeadershipTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadershipTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).LeadershipTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_LeadershipTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).LeadershipTransfer(ctx, req.(*LeadershipTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_LeadershipTransferToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadershipTransferToServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).LeadershipTransferToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_LeadershipTransferToServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).LeadershipTransferToServer(ctx, req.(*LeadershipTransferToServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_RemoveServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).RemoveServer(ctx, req.(*RemoveServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Shutdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Snapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Snapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_State_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Stats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_VerifyLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).VerifyLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_VerifyLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).VerifyLeader(ctx, req.(*VerifyLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Await_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Future)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Await(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Await_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Await(ctx, req.(*Future))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftAdmin_Forget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Future)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftAdminServer).Forget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftAdmin_Forget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftAdminServer).Forget(ctx, req.(*Future))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftAdmin_ServiceDesc is the grpc.ServiceDesc for RaftAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RaftAdmin",
	HandlerType: (*RaftAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNonvoter",
			Handler:    _RaftAdmin_AddNonvoter_Handler,
		},
		{
			MethodName: "AddVoter",
			Handler:    _RaftAdmin_AddVoter_Handler,
		},
		{
			MethodName: "AppliedIndex",
			Handler:    _RaftAdmin_AppliedIndex_Handler,
		},
		{
			MethodName: "ApplyLog",
			Handler:    _RaftAdmin_ApplyLog_Handler,
		},
		{
			MethodName: "Barrier",
			Handler:    _RaftAdmin_Barrier_Handler,
		},
		{
			MethodName: "DemoteVoter",
			Handler:    _RaftAdmin_DemoteVoter_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _RaftAdmin_GetConfiguration_Handler,
		},
		{
			MethodName: "LastContact",
			Handler:    _RaftAdmin_LastContact_Handler,
		},
		{
			MethodName: "LastIndex",
			Handler:    _RaftAdmin_LastIndex_Handler,
		},
		{
			MethodName: "Leader",
			Handler:    _RaftAdmin_Leader_Handler,
		},
		{
			MethodName: "LeadershipTransfer",
			Handler:    _RaftAdmin_LeadershipTransfer_Handler,
		},
		{
			MethodName: "LeadershipTransferToServer",
			Handler:    _RaftAdmin_LeadershipTransferToServer_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _RaftAdmin_RemoveServer_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _RaftAdmin_Shutdown_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _RaftAdmin_Snapshot_Handler,
		},
		{
			MethodName: "State",
			Handler:    _RaftAdmin_State_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _RaftAdmin_Stats_Handler,
		},
		{
			MethodName: "VerifyLeader",
			Handler:    _RaftAdmin_VerifyLeader_Handler,
		},
		{
			MethodName: "Await",
			Handler:    _RaftAdmin_Await_Handler,
		},
		{
			MethodName: "Forget",
			Handler:    _RaftAdmin_Forget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raftadmin.proto",
}
