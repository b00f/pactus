// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: wallet.proto

package pactus

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Wallet_CreateWallet_FullMethodName        = "/pactus.Wallet/CreateWallet"
	Wallet_RestoreWallet_FullMethodName       = "/pactus.Wallet/RestoreWallet"
	Wallet_LoadWallet_FullMethodName          = "/pactus.Wallet/LoadWallet"
	Wallet_UnloadWallet_FullMethodName        = "/pactus.Wallet/UnloadWallet"
	Wallet_GetTotalBalance_FullMethodName     = "/pactus.Wallet/GetTotalBalance"
	Wallet_SignRawTransaction_FullMethodName  = "/pactus.Wallet/SignRawTransaction"
	Wallet_GetValidatorAddress_FullMethodName = "/pactus.Wallet/GetValidatorAddress"
	Wallet_GetNewAddress_FullMethodName       = "/pactus.Wallet/GetNewAddress"
	Wallet_GetAddressHistory_FullMethodName   = "/pactus.Wallet/GetAddressHistory"
	Wallet_SignMessage_FullMethodName         = "/pactus.Wallet/SignMessage"
	Wallet_GetTotalStake_FullMethodName       = "/pactus.Wallet/GetTotalStake"
	Wallet_GetAddressInfo_FullMethodName      = "/pactus.Wallet/GetAddressInfo"
	Wallet_SetAddressLabel_FullMethodName     = "/pactus.Wallet/SetAddressLabel"
	Wallet_ListWallet_FullMethodName          = "/pactus.Wallet/ListWallet"
	Wallet_GetWalletInfo_FullMethodName       = "/pactus.Wallet/GetWalletInfo"
	Wallet_ListAddress_FullMethodName         = "/pactus.Wallet/ListAddress"
)

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Wallet service provides RPC methods for wallet management operations.
type WalletClient interface {
	// CreateWallet creates a new wallet with the specified parameters.
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	// RestoreWallet restores an existing wallet with the given mnemonic.
	RestoreWallet(ctx context.Context, in *RestoreWalletRequest, opts ...grpc.CallOption) (*RestoreWalletResponse, error)
	// LoadWallet loads an existing wallet with the given name.
	LoadWallet(ctx context.Context, in *LoadWalletRequest, opts ...grpc.CallOption) (*LoadWalletResponse, error)
	// UnloadWallet unloads a currently loaded wallet with the specified name.
	UnloadWallet(ctx context.Context, in *UnloadWalletRequest, opts ...grpc.CallOption) (*UnloadWalletResponse, error)
	// GetTotalBalance returns the total available balance of the wallet.
	GetTotalBalance(ctx context.Context, in *GetTotalBalanceRequest, opts ...grpc.CallOption) (*GetTotalBalanceResponse, error)
	// SignRawTransaction signs a raw transaction for a specified wallet.
	SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error)
	// GetValidatorAddress retrieves the validator address associated with a public key.
	GetValidatorAddress(ctx context.Context, in *GetValidatorAddressRequest, opts ...grpc.CallOption) (*GetValidatorAddressResponse, error)
	// GetNewAddress generates a new address for the specified wallet.
	GetNewAddress(ctx context.Context, in *GetNewAddressRequest, opts ...grpc.CallOption) (*GetNewAddressResponse, error)
	// GetAddressHistory retrieves the transaction history of an address.
	GetAddressHistory(ctx context.Context, in *GetAddressHistoryRequest, opts ...grpc.CallOption) (*GetAddressHistoryResponse, error)
	// SignMessage signs an arbitrary message using a wallet's private key.
	SignMessage(ctx context.Context, in *SignMessageRequest, opts ...grpc.CallOption) (*SignMessageResponse, error)
	// GetTotalStake returns the total stake amount in the wallet.
	GetTotalStake(ctx context.Context, in *GetTotalStakeRequest, opts ...grpc.CallOption) (*GetTotalStakeResponse, error)
	// GetAddressInfo returns detailed information about a specific address.
	GetAddressInfo(ctx context.Context, in *GetAddressInfoRequest, opts ...grpc.CallOption) (*GetAddressInfoResponse, error)
	// SetAddressLabel sets or updates the label for a given address.
	SetAddressLabel(ctx context.Context, in *SetAddressLabelRequest, opts ...grpc.CallOption) (*SetAddressLabelResponse, error)
	// ListWallet returns list of all available wallets.
	ListWallet(ctx context.Context, in *ListWalletRequest, opts ...grpc.CallOption) (*ListWalletResponse, error)
	// GetWalletInfo returns detailed information about a specific wallet.
	GetWalletInfo(ctx context.Context, in *GetWalletInfoRequest, opts ...grpc.CallOption) (*GetWalletInfoResponse, error)
	// ListAddress returns all addresses in the specified wallet.
	ListAddress(ctx context.Context, in *ListAddressRequest, opts ...grpc.CallOption) (*ListAddressResponse, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, Wallet_CreateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) RestoreWallet(ctx context.Context, in *RestoreWalletRequest, opts ...grpc.CallOption) (*RestoreWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestoreWalletResponse)
	err := c.cc.Invoke(ctx, Wallet_RestoreWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) LoadWallet(ctx context.Context, in *LoadWalletRequest, opts ...grpc.CallOption) (*LoadWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadWalletResponse)
	err := c.cc.Invoke(ctx, Wallet_LoadWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UnloadWallet(ctx context.Context, in *UnloadWalletRequest, opts ...grpc.CallOption) (*UnloadWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnloadWalletResponse)
	err := c.cc.Invoke(ctx, Wallet_UnloadWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetTotalBalance(ctx context.Context, in *GetTotalBalanceRequest, opts ...grpc.CallOption) (*GetTotalBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTotalBalanceResponse)
	err := c.cc.Invoke(ctx, Wallet_GetTotalBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignRawTransactionResponse)
	err := c.cc.Invoke(ctx, Wallet_SignRawTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetValidatorAddress(ctx context.Context, in *GetValidatorAddressRequest, opts ...grpc.CallOption) (*GetValidatorAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetValidatorAddressResponse)
	err := c.cc.Invoke(ctx, Wallet_GetValidatorAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetNewAddress(ctx context.Context, in *GetNewAddressRequest, opts ...grpc.CallOption) (*GetNewAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewAddressResponse)
	err := c.cc.Invoke(ctx, Wallet_GetNewAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAddressHistory(ctx context.Context, in *GetAddressHistoryRequest, opts ...grpc.CallOption) (*GetAddressHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAddressHistoryResponse)
	err := c.cc.Invoke(ctx, Wallet_GetAddressHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SignMessage(ctx context.Context, in *SignMessageRequest, opts ...grpc.CallOption) (*SignMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignMessageResponse)
	err := c.cc.Invoke(ctx, Wallet_SignMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetTotalStake(ctx context.Context, in *GetTotalStakeRequest, opts ...grpc.CallOption) (*GetTotalStakeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTotalStakeResponse)
	err := c.cc.Invoke(ctx, Wallet_GetTotalStake_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAddressInfo(ctx context.Context, in *GetAddressInfoRequest, opts ...grpc.CallOption) (*GetAddressInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAddressInfoResponse)
	err := c.cc.Invoke(ctx, Wallet_GetAddressInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SetAddressLabel(ctx context.Context, in *SetAddressLabelRequest, opts ...grpc.CallOption) (*SetAddressLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetAddressLabelResponse)
	err := c.cc.Invoke(ctx, Wallet_SetAddressLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ListWallet(ctx context.Context, in *ListWalletRequest, opts ...grpc.CallOption) (*ListWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWalletResponse)
	err := c.cc.Invoke(ctx, Wallet_ListWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetWalletInfo(ctx context.Context, in *GetWalletInfoRequest, opts ...grpc.CallOption) (*GetWalletInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletInfoResponse)
	err := c.cc.Invoke(ctx, Wallet_GetWalletInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ListAddress(ctx context.Context, in *ListAddressRequest, opts ...grpc.CallOption) (*ListAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAddressResponse)
	err := c.cc.Invoke(ctx, Wallet_ListAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations should embed UnimplementedWalletServer
// for forward compatibility.
//
// Wallet service provides RPC methods for wallet management operations.
type WalletServer interface {
	// CreateWallet creates a new wallet with the specified parameters.
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	// RestoreWallet restores an existing wallet with the given mnemonic.
	RestoreWallet(context.Context, *RestoreWalletRequest) (*RestoreWalletResponse, error)
	// LoadWallet loads an existing wallet with the given name.
	LoadWallet(context.Context, *LoadWalletRequest) (*LoadWalletResponse, error)
	// UnloadWallet unloads a currently loaded wallet with the specified name.
	UnloadWallet(context.Context, *UnloadWalletRequest) (*UnloadWalletResponse, error)
	// GetTotalBalance returns the total available balance of the wallet.
	GetTotalBalance(context.Context, *GetTotalBalanceRequest) (*GetTotalBalanceResponse, error)
	// SignRawTransaction signs a raw transaction for a specified wallet.
	SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error)
	// GetValidatorAddress retrieves the validator address associated with a public key.
	GetValidatorAddress(context.Context, *GetValidatorAddressRequest) (*GetValidatorAddressResponse, error)
	// GetNewAddress generates a new address for the specified wallet.
	GetNewAddress(context.Context, *GetNewAddressRequest) (*GetNewAddressResponse, error)
	// GetAddressHistory retrieves the transaction history of an address.
	GetAddressHistory(context.Context, *GetAddressHistoryRequest) (*GetAddressHistoryResponse, error)
	// SignMessage signs an arbitrary message using a wallet's private key.
	SignMessage(context.Context, *SignMessageRequest) (*SignMessageResponse, error)
	// GetTotalStake returns the total stake amount in the wallet.
	GetTotalStake(context.Context, *GetTotalStakeRequest) (*GetTotalStakeResponse, error)
	// GetAddressInfo returns detailed information about a specific address.
	GetAddressInfo(context.Context, *GetAddressInfoRequest) (*GetAddressInfoResponse, error)
	// SetAddressLabel sets or updates the label for a given address.
	SetAddressLabel(context.Context, *SetAddressLabelRequest) (*SetAddressLabelResponse, error)
	// ListWallet returns list of all available wallets.
	ListWallet(context.Context, *ListWalletRequest) (*ListWalletResponse, error)
	// GetWalletInfo returns detailed information about a specific wallet.
	GetWalletInfo(context.Context, *GetWalletInfoRequest) (*GetWalletInfoResponse, error)
	// ListAddress returns all addresses in the specified wallet.
	ListAddress(context.Context, *ListAddressRequest) (*ListAddressResponse, error)
}

// UnimplementedWalletServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletServer struct{}

func (UnimplementedWalletServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServer) RestoreWallet(context.Context, *RestoreWalletRequest) (*RestoreWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreWallet not implemented")
}
func (UnimplementedWalletServer) LoadWallet(context.Context, *LoadWalletRequest) (*LoadWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadWallet not implemented")
}
func (UnimplementedWalletServer) UnloadWallet(context.Context, *UnloadWalletRequest) (*UnloadWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnloadWallet not implemented")
}
func (UnimplementedWalletServer) GetTotalBalance(context.Context, *GetTotalBalanceRequest) (*GetTotalBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalBalance not implemented")
}
func (UnimplementedWalletServer) SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignRawTransaction not implemented")
}
func (UnimplementedWalletServer) GetValidatorAddress(context.Context, *GetValidatorAddressRequest) (*GetValidatorAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorAddress not implemented")
}
func (UnimplementedWalletServer) GetNewAddress(context.Context, *GetNewAddressRequest) (*GetNewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewAddress not implemented")
}
func (UnimplementedWalletServer) GetAddressHistory(context.Context, *GetAddressHistoryRequest) (*GetAddressHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressHistory not implemented")
}
func (UnimplementedWalletServer) SignMessage(context.Context, *SignMessageRequest) (*SignMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignMessage not implemented")
}
func (UnimplementedWalletServer) GetTotalStake(context.Context, *GetTotalStakeRequest) (*GetTotalStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalStake not implemented")
}
func (UnimplementedWalletServer) GetAddressInfo(context.Context, *GetAddressInfoRequest) (*GetAddressInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressInfo not implemented")
}
func (UnimplementedWalletServer) SetAddressLabel(context.Context, *SetAddressLabelRequest) (*SetAddressLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAddressLabel not implemented")
}
func (UnimplementedWalletServer) ListWallet(context.Context, *ListWalletRequest) (*ListWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallet not implemented")
}
func (UnimplementedWalletServer) GetWalletInfo(context.Context, *GetWalletInfoRequest) (*GetWalletInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletInfo not implemented")
}
func (UnimplementedWalletServer) ListAddress(context.Context, *ListAddressRequest) (*ListAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddress not implemented")
}
func (UnimplementedWalletServer) testEmbeddedByValue() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	// If the following call pancis, it indicates UnimplementedWalletServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_RestoreWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).RestoreWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_RestoreWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).RestoreWallet(ctx, req.(*RestoreWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_LoadWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).LoadWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_LoadWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).LoadWallet(ctx, req.(*LoadWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UnloadWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnloadWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UnloadWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_UnloadWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UnloadWallet(ctx, req.(*UnloadWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetTotalBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTotalBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetTotalBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTotalBalance(ctx, req.(*GetTotalBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SignRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SignRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_SignRawTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SignRawTransaction(ctx, req.(*SignRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetValidatorAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetValidatorAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetValidatorAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetValidatorAddress(ctx, req.(*GetValidatorAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetNewAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetNewAddress(ctx, req.(*GetNewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAddressHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAddressHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetAddressHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAddressHistory(ctx, req.(*GetAddressHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SignMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SignMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_SignMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SignMessage(ctx, req.(*SignMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetTotalStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalStakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTotalStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetTotalStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTotalStake(ctx, req.(*GetTotalStakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAddressInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAddressInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetAddressInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAddressInfo(ctx, req.(*GetAddressInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SetAddressLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAddressLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SetAddressLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_SetAddressLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SetAddressLabel(ctx, req.(*SetAddressLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ListWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_ListWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListWallet(ctx, req.(*ListWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetWalletInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetWalletInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetWalletInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetWalletInfo(ctx, req.(*GetWalletInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ListAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_ListAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListAddress(ctx, req.(*ListAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pactus.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _Wallet_CreateWallet_Handler,
		},
		{
			MethodName: "RestoreWallet",
			Handler:    _Wallet_RestoreWallet_Handler,
		},
		{
			MethodName: "LoadWallet",
			Handler:    _Wallet_LoadWallet_Handler,
		},
		{
			MethodName: "UnloadWallet",
			Handler:    _Wallet_UnloadWallet_Handler,
		},
		{
			MethodName: "GetTotalBalance",
			Handler:    _Wallet_GetTotalBalance_Handler,
		},
		{
			MethodName: "SignRawTransaction",
			Handler:    _Wallet_SignRawTransaction_Handler,
		},
		{
			MethodName: "GetValidatorAddress",
			Handler:    _Wallet_GetValidatorAddress_Handler,
		},
		{
			MethodName: "GetNewAddress",
			Handler:    _Wallet_GetNewAddress_Handler,
		},
		{
			MethodName: "GetAddressHistory",
			Handler:    _Wallet_GetAddressHistory_Handler,
		},
		{
			MethodName: "SignMessage",
			Handler:    _Wallet_SignMessage_Handler,
		},
		{
			MethodName: "GetTotalStake",
			Handler:    _Wallet_GetTotalStake_Handler,
		},
		{
			MethodName: "GetAddressInfo",
			Handler:    _Wallet_GetAddressInfo_Handler,
		},
		{
			MethodName: "SetAddressLabel",
			Handler:    _Wallet_SetAddressLabel_Handler,
		},
		{
			MethodName: "ListWallet",
			Handler:    _Wallet_ListWallet_Handler,
		},
		{
			MethodName: "GetWalletInfo",
			Handler:    _Wallet_GetWalletInfo_Handler,
		},
		{
			MethodName: "ListAddress",
			Handler:    _Wallet_ListAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
