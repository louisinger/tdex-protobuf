// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package operator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	// Returns a new derived address for the given market.
	// If market field is empty, a new Market is created and MUST be initialized.
	DepositMarket(ctx context.Context, in *DepositMarketRequest, opts ...grpc.CallOption) (*DepositMarketReply, error)
	// Returns the list of previously generated addresses for the given market.
	ListDepositMarket(ctx context.Context, in *ListDepositMarketRequest, opts ...grpc.CallOption) (*ListDepositMarketReply, error)
	// Returns a new derived address from the fee account.
	// This is only used to deposit some LBTC to subsidize blockchain fees.
	DepositFeeAccount(ctx context.Context, in *DepositFeeAccountRequest, opts ...grpc.CallOption) (*DepositFeeAccountReply, error)
	// Returns the aggregated balance of LBTC held in the fee account.
	BalanceFeeAccount(ctx context.Context, in *BalanceFeeAccountRequest, opts ...grpc.CallOption) (*BalanceFeeAccountReply, error)
	// Makes the given market tradable
	OpenMarket(ctx context.Context, in *OpenMarketRequest, opts ...grpc.CallOption) (*OpenMarketReply, error)
	// Makes the given market NOT tradabale
	CloseMarket(ctx context.Context, in *CloseMarketRequest, opts ...grpc.CallOption) (*CloseMarketReply, error)
	// Get extended details for each markets either open, closed or to be funded.
	ListMarket(ctx context.Context, in *ListMarketRequest, opts ...grpc.CallOption) (*ListMarketReply, error)
	// Changes the Liquidity Provider fee for the given market. I thsould be
	// express in basis point. To change the fee on each swap from (current) 0.25%
	// to 1% you need to pass down 100 The Market MUST be closed before doing this
	// change.
	UpdateMarketFee(ctx context.Context, in *UpdateMarketFeeRequest, opts ...grpc.CallOption) (*UpdateMarketFeeReply, error)
	// Manually updates the price for the given market
	UpdateMarketPrice(ctx context.Context, in *UpdateMarketPriceRequest, opts ...grpc.CallOption) (*UpdateMarketPriceReply, error)
	// Updates the current market making strategy, either using an automated
	// market making formula or a pluggable price feed
	UpdateMarketStrategy(ctx context.Context, in *UpdateMarketStrategyRequest, opts ...grpc.CallOption) (*UpdateMarketStrategyReply, error)
	// WithdrawMarket allows the operator to withdraw to external wallet funds
	// from a specific market. The Market MUST be closed before doing this change.
	WithdrawMarket(ctx context.Context, in *WithdrawMarketRequest, opts ...grpc.CallOption) (*WithdrawMarketReply, error)
	// Returs all the swaps processed by the daemon (both attempted and completed)
	ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsReply, error)
	// Displays a report on how much the given market is collecting in Liquidity
	// Provider fees
	ReportMarketFee(ctx context.Context, in *ReportMarketFeeRequest, opts ...grpc.CallOption) (*ReportMarketFeeReply, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) DepositMarket(ctx context.Context, in *DepositMarketRequest, opts ...grpc.CallOption) (*DepositMarketReply, error) {
	out := new(DepositMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/DepositMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListDepositMarket(ctx context.Context, in *ListDepositMarketRequest, opts ...grpc.CallOption) (*ListDepositMarketReply, error) {
	out := new(ListDepositMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/ListDepositMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) DepositFeeAccount(ctx context.Context, in *DepositFeeAccountRequest, opts ...grpc.CallOption) (*DepositFeeAccountReply, error) {
	out := new(DepositFeeAccountReply)
	err := c.cc.Invoke(ctx, "/Operator/DepositFeeAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) BalanceFeeAccount(ctx context.Context, in *BalanceFeeAccountRequest, opts ...grpc.CallOption) (*BalanceFeeAccountReply, error) {
	out := new(BalanceFeeAccountReply)
	err := c.cc.Invoke(ctx, "/Operator/BalanceFeeAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) OpenMarket(ctx context.Context, in *OpenMarketRequest, opts ...grpc.CallOption) (*OpenMarketReply, error) {
	out := new(OpenMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/OpenMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) CloseMarket(ctx context.Context, in *CloseMarketRequest, opts ...grpc.CallOption) (*CloseMarketReply, error) {
	out := new(CloseMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/CloseMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListMarket(ctx context.Context, in *ListMarketRequest, opts ...grpc.CallOption) (*ListMarketReply, error) {
	out := new(ListMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/ListMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) UpdateMarketFee(ctx context.Context, in *UpdateMarketFeeRequest, opts ...grpc.CallOption) (*UpdateMarketFeeReply, error) {
	out := new(UpdateMarketFeeReply)
	err := c.cc.Invoke(ctx, "/Operator/UpdateMarketFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) UpdateMarketPrice(ctx context.Context, in *UpdateMarketPriceRequest, opts ...grpc.CallOption) (*UpdateMarketPriceReply, error) {
	out := new(UpdateMarketPriceReply)
	err := c.cc.Invoke(ctx, "/Operator/UpdateMarketPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) UpdateMarketStrategy(ctx context.Context, in *UpdateMarketStrategyRequest, opts ...grpc.CallOption) (*UpdateMarketStrategyReply, error) {
	out := new(UpdateMarketStrategyReply)
	err := c.cc.Invoke(ctx, "/Operator/UpdateMarketStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) WithdrawMarket(ctx context.Context, in *WithdrawMarketRequest, opts ...grpc.CallOption) (*WithdrawMarketReply, error) {
	out := new(WithdrawMarketReply)
	err := c.cc.Invoke(ctx, "/Operator/WithdrawMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsReply, error) {
	out := new(ListSwapsReply)
	err := c.cc.Invoke(ctx, "/Operator/ListSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ReportMarketFee(ctx context.Context, in *ReportMarketFeeRequest, opts ...grpc.CallOption) (*ReportMarketFeeReply, error) {
	out := new(ReportMarketFeeReply)
	err := c.cc.Invoke(ctx, "/Operator/ReportMarketFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations must embed UnimplementedOperatorServer
// for forward compatibility
type OperatorServer interface {
	// Returns a new derived address for the given market.
	// If market field is empty, a new Market is created and MUST be initialized.
	DepositMarket(context.Context, *DepositMarketRequest) (*DepositMarketReply, error)
	// Returns the list of previously generated addresses for the given market.
	ListDepositMarket(context.Context, *ListDepositMarketRequest) (*ListDepositMarketReply, error)
	// Returns a new derived address from the fee account.
	// This is only used to deposit some LBTC to subsidize blockchain fees.
	DepositFeeAccount(context.Context, *DepositFeeAccountRequest) (*DepositFeeAccountReply, error)
	// Returns the aggregated balance of LBTC held in the fee account.
	BalanceFeeAccount(context.Context, *BalanceFeeAccountRequest) (*BalanceFeeAccountReply, error)
	// Makes the given market tradable
	OpenMarket(context.Context, *OpenMarketRequest) (*OpenMarketReply, error)
	// Makes the given market NOT tradabale
	CloseMarket(context.Context, *CloseMarketRequest) (*CloseMarketReply, error)
	// Get extended details for each markets either open, closed or to be funded.
	ListMarket(context.Context, *ListMarketRequest) (*ListMarketReply, error)
	// Changes the Liquidity Provider fee for the given market. I thsould be
	// express in basis point. To change the fee on each swap from (current) 0.25%
	// to 1% you need to pass down 100 The Market MUST be closed before doing this
	// change.
	UpdateMarketFee(context.Context, *UpdateMarketFeeRequest) (*UpdateMarketFeeReply, error)
	// Manually updates the price for the given market
	UpdateMarketPrice(context.Context, *UpdateMarketPriceRequest) (*UpdateMarketPriceReply, error)
	// Updates the current market making strategy, either using an automated
	// market making formula or a pluggable price feed
	UpdateMarketStrategy(context.Context, *UpdateMarketStrategyRequest) (*UpdateMarketStrategyReply, error)
	// WithdrawMarket allows the operator to withdraw to external wallet funds
	// from a specific market. The Market MUST be closed before doing this change.
	WithdrawMarket(context.Context, *WithdrawMarketRequest) (*WithdrawMarketReply, error)
	// Returs all the swaps processed by the daemon (both attempted and completed)
	ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsReply, error)
	// Displays a report on how much the given market is collecting in Liquidity
	// Provider fees
	ReportMarketFee(context.Context, *ReportMarketFeeRequest) (*ReportMarketFeeReply, error)
	mustEmbedUnimplementedOperatorServer()
}

// UnimplementedOperatorServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServer struct {
}

func (*UnimplementedOperatorServer) DepositMarket(context.Context, *DepositMarketRequest) (*DepositMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositMarket not implemented")
}
func (*UnimplementedOperatorServer) ListDepositMarket(context.Context, *ListDepositMarketRequest) (*ListDepositMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepositMarket not implemented")
}
func (*UnimplementedOperatorServer) DepositFeeAccount(context.Context, *DepositFeeAccountRequest) (*DepositFeeAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositFeeAccount not implemented")
}
func (*UnimplementedOperatorServer) BalanceFeeAccount(context.Context, *BalanceFeeAccountRequest) (*BalanceFeeAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceFeeAccount not implemented")
}
func (*UnimplementedOperatorServer) OpenMarket(context.Context, *OpenMarketRequest) (*OpenMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenMarket not implemented")
}
func (*UnimplementedOperatorServer) CloseMarket(context.Context, *CloseMarketRequest) (*CloseMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseMarket not implemented")
}
func (*UnimplementedOperatorServer) ListMarket(context.Context, *ListMarketRequest) (*ListMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMarket not implemented")
}
func (*UnimplementedOperatorServer) UpdateMarketFee(context.Context, *UpdateMarketFeeRequest) (*UpdateMarketFeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketFee not implemented")
}
func (*UnimplementedOperatorServer) UpdateMarketPrice(context.Context, *UpdateMarketPriceRequest) (*UpdateMarketPriceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketPrice not implemented")
}
func (*UnimplementedOperatorServer) UpdateMarketStrategy(context.Context, *UpdateMarketStrategyRequest) (*UpdateMarketStrategyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketStrategy not implemented")
}
func (*UnimplementedOperatorServer) WithdrawMarket(context.Context, *WithdrawMarketRequest) (*WithdrawMarketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawMarket not implemented")
}
func (*UnimplementedOperatorServer) ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSwaps not implemented")
}
func (*UnimplementedOperatorServer) ReportMarketFee(context.Context, *ReportMarketFeeRequest) (*ReportMarketFeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMarketFee not implemented")
}
func (*UnimplementedOperatorServer) mustEmbedUnimplementedOperatorServer() {}

func RegisterOperatorServer(s *grpc.Server, srv OperatorServer) {
	s.RegisterService(&_Operator_serviceDesc, srv)
}

func _Operator_DepositMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).DepositMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/DepositMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).DepositMarket(ctx, req.(*DepositMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListDepositMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepositMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListDepositMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/ListDepositMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListDepositMarket(ctx, req.(*ListDepositMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_DepositFeeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositFeeAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).DepositFeeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/DepositFeeAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).DepositFeeAccount(ctx, req.(*DepositFeeAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_BalanceFeeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceFeeAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).BalanceFeeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/BalanceFeeAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).BalanceFeeAccount(ctx, req.(*BalanceFeeAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_OpenMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).OpenMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/OpenMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).OpenMarket(ctx, req.(*OpenMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_CloseMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).CloseMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/CloseMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).CloseMarket(ctx, req.(*CloseMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/ListMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListMarket(ctx, req.(*ListMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_UpdateMarketFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).UpdateMarketFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/UpdateMarketFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).UpdateMarketFee(ctx, req.(*UpdateMarketFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_UpdateMarketPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).UpdateMarketPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/UpdateMarketPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).UpdateMarketPrice(ctx, req.(*UpdateMarketPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_UpdateMarketStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).UpdateMarketStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/UpdateMarketStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).UpdateMarketStrategy(ctx, req.(*UpdateMarketStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_WithdrawMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).WithdrawMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/WithdrawMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).WithdrawMarket(ctx, req.(*WithdrawMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/ListSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ReportMarketFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportMarketFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ReportMarketFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Operator/ReportMarketFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ReportMarketFee(ctx, req.(*ReportMarketFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DepositMarket",
			Handler:    _Operator_DepositMarket_Handler,
		},
		{
			MethodName: "ListDepositMarket",
			Handler:    _Operator_ListDepositMarket_Handler,
		},
		{
			MethodName: "DepositFeeAccount",
			Handler:    _Operator_DepositFeeAccount_Handler,
		},
		{
			MethodName: "BalanceFeeAccount",
			Handler:    _Operator_BalanceFeeAccount_Handler,
		},
		{
			MethodName: "OpenMarket",
			Handler:    _Operator_OpenMarket_Handler,
		},
		{
			MethodName: "CloseMarket",
			Handler:    _Operator_CloseMarket_Handler,
		},
		{
			MethodName: "ListMarket",
			Handler:    _Operator_ListMarket_Handler,
		},
		{
			MethodName: "UpdateMarketFee",
			Handler:    _Operator_UpdateMarketFee_Handler,
		},
		{
			MethodName: "UpdateMarketPrice",
			Handler:    _Operator_UpdateMarketPrice_Handler,
		},
		{
			MethodName: "UpdateMarketStrategy",
			Handler:    _Operator_UpdateMarketStrategy_Handler,
		},
		{
			MethodName: "WithdrawMarket",
			Handler:    _Operator_WithdrawMarket_Handler,
		},
		{
			MethodName: "ListSwaps",
			Handler:    _Operator_ListSwaps_Handler,
		},
		{
			MethodName: "ReportMarketFee",
			Handler:    _Operator_ReportMarketFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operator.proto",
}
