// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: coinmeca_grpc.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VaultTradeType int32

const (
	VaultTradeType_DEPOSIT  VaultTradeType = 0
	VaultTradeType_WITHDRAW VaultTradeType = 1
)

// Enum value maps for VaultTradeType.
var (
	VaultTradeType_name = map[int32]string{
		0: "DEPOSIT",
		1: "WITHDRAW",
	}
	VaultTradeType_value = map[string]int32{
		"DEPOSIT":  0,
		"WITHDRAW": 1,
	}
)

func (x VaultTradeType) Enum() *VaultTradeType {
	p := new(VaultTradeType)
	*p = x
	return p
}

func (x VaultTradeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VaultTradeType) Descriptor() protoreflect.EnumDescriptor {
	return file_coinmeca_grpc_proto_enumTypes[0].Descriptor()
}

func (VaultTradeType) Type() protoreflect.EnumType {
	return &file_coinmeca_grpc_proto_enumTypes[0]
}

func (x VaultTradeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VaultTradeType.Descriptor instead.
func (VaultTradeType) EnumDescriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{0}
}

type GeneralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*GeneralRequest_Transaction
	//	*GeneralRequest_VaultRecent
	Request isGeneralRequest_Request `protobuf_oneof:"request"`
}

func (x *GeneralRequest) Reset() {
	*x = GeneralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralRequest) ProtoMessage() {}

func (x *GeneralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralRequest.ProtoReflect.Descriptor instead.
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{0}
}

func (m *GeneralRequest) GetRequest() isGeneralRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *GeneralRequest) GetTransaction() *TransactionRequest {
	if x, ok := x.GetRequest().(*GeneralRequest_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (x *GeneralRequest) GetVaultRecent() *VaultRecentRequest {
	if x, ok := x.GetRequest().(*GeneralRequest_VaultRecent); ok {
		return x.VaultRecent
	}
	return nil
}

type isGeneralRequest_Request interface {
	isGeneralRequest_Request()
}

type GeneralRequest_Transaction struct {
	Transaction *TransactionRequest `protobuf:"bytes,1,opt,name=transaction,proto3,oneof"`
}

type GeneralRequest_VaultRecent struct {
	VaultRecent *VaultRecentRequest `protobuf:"bytes,2,opt,name=vaultRecent,proto3,oneof"`
}

func (*GeneralRequest_Transaction) isGeneralRequest_Request() {}

func (*GeneralRequest_VaultRecent) isGeneralRequest_Request() {}

type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ====================================
// AddTransaction
// ====================================
type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Transaction `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionRequest) GetData() *Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	BlockHash   string   `protobuf:"bytes,2,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	BlockNumber string   `protobuf:"bytes,3,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	Hash        string   `protobuf:"bytes,4,opt,name=Hash,proto3" json:"Hash,omitempty"`
	AccessList  []string `protobuf:"bytes,5,rep,name=AccessList,proto3" json:"AccessList,omitempty"`
	ChainId     string   `protobuf:"bytes,6,opt,name=ChainId,proto3" json:"ChainId,omitempty"`
	From        string   `protobuf:"bytes,7,opt,name=From,proto3" json:"From,omitempty"`
	To          string   `protobuf:"bytes,8,opt,name=To,proto3" json:"To,omitempty"`
	Gas         string   `protobuf:"bytes,9,opt,name=Gas,proto3" json:"Gas,omitempty"`
	GasPrice    string   `protobuf:"bytes,10,opt,name=GasPrice,proto3" json:"GasPrice,omitempty"`
	Input       string   `protobuf:"bytes,11,opt,name=Input,proto3" json:"Input,omitempty"`
	Cate        string   `protobuf:"bytes,12,opt,name=Cate,proto3" json:"Cate,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetAccessList() []string {
	if x != nil {
		return x.AccessList
	}
	return nil
}

func (x *Transaction) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Transaction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Transaction) GetCate() string {
	if x != nil {
		return x.Cate
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ====================================
// Vault
// ====================================
type VaultRecentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *VaultRecent `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *VaultRecentRequest) Reset() {
	*x = VaultRecentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultRecentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultRecentRequest) ProtoMessage() {}

func (x *VaultRecentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultRecentRequest.ProtoReflect.Descriptor instead.
func (*VaultRecentRequest) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *VaultRecentRequest) GetData() *VaultRecent {
	if x != nil {
		return x.Data
	}
	return nil
}

type VaultRecent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId      int64          `protobuf:"varint,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	TokenAddress string         `protobuf:"bytes,2,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	OwnerAddress string         `protobuf:"bytes,3,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	Type         VaultTradeType `protobuf:"varint,4,opt,name=type,proto3,enum=grpcmodule.VaultTradeType" json:"type,omitempty"`
	Time         string         `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Amount       float64        `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Meca         float64        `protobuf:"fixed64,7,opt,name=meca,proto3" json:"meca,omitempty"`
	Share        float64        `protobuf:"fixed64,8,opt,name=share,proto3" json:"share,omitempty"`
	TxHash       string         `protobuf:"bytes,9,opt,name=txHash,proto3" json:"txHash,omitempty"`
	UpdateAt     string         `protobuf:"bytes,10,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *VaultRecent) Reset() {
	*x = VaultRecent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultRecent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultRecent) ProtoMessage() {}

func (x *VaultRecent) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultRecent.ProtoReflect.Descriptor instead.
func (*VaultRecent) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *VaultRecent) GetChainId() int64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *VaultRecent) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *VaultRecent) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *VaultRecent) GetType() VaultTradeType {
	if x != nil {
		return x.Type
	}
	return VaultTradeType_DEPOSIT
}

func (x *VaultRecent) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *VaultRecent) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *VaultRecent) GetMeca() float64 {
	if x != nil {
		return x.Meca
	}
	return 0
}

func (x *VaultRecent) GetShare() float64 {
	if x != nil {
		return x.Share
	}
	return 0
}

func (x *VaultRecent) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *VaultRecent) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type VaultRecentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VaultRecentResponse) Reset() {
	*x = VaultRecentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultRecentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultRecentResponse) ProtoMessage() {}

func (x *VaultRecentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultRecentResponse.ProtoReflect.Descriptor instead.
func (*VaultRecentResponse) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *VaultRecentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAllFromVaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GetAllFromVault `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllFromVaultRequest) Reset() {
	*x = GetAllFromVaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFromVaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFromVaultRequest) ProtoMessage() {}

func (x *GetAllFromVaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFromVaultRequest.ProtoReflect.Descriptor instead.
func (*GetAllFromVaultRequest) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllFromVaultRequest) GetData() *GetAllFromVault {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllFromVault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Decimals int64  `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Exchange string `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Key      bool   `protobuf:"varint,4,opt,name=key,proto3" json:"key,omitempty"`
	Locked   string `protobuf:"bytes,5,opt,name=locked,proto3" json:"locked,omitempty"`
	Name     string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Need     string `protobuf:"bytes,7,opt,name=need,proto3" json:"need,omitempty"`
	Rate     string `protobuf:"bytes,8,opt,name=rate,proto3" json:"rate,omitempty"`
	Symbol   string `protobuf:"bytes,9,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Weight   string `protobuf:"bytes,10,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *GetAllFromVault) Reset() {
	*x = GetAllFromVault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinmeca_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFromVault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFromVault) ProtoMessage() {}

func (x *GetAllFromVault) ProtoReflect() protoreflect.Message {
	mi := &file_coinmeca_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFromVault.ProtoReflect.Descriptor instead.
func (*GetAllFromVault) Descriptor() ([]byte, []int) {
	return file_coinmeca_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllFromVault) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetAllFromVault) GetDecimals() int64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *GetAllFromVault) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *GetAllFromVault) GetKey() bool {
	if x != nil {
		return x.Key
	}
	return false
}

func (x *GetAllFromVault) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *GetAllFromVault) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllFromVault) GetNeed() string {
	if x != nil {
		return x.Need
	}
	return ""
}

func (x *GetAllFromVault) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *GetAllFromVault) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetAllFromVault) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

var File_coinmeca_grpc_proto protoreflect.FileDescriptor

var file_coinmeca_grpc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x69, 0x6e, 0x6d, 0x65, 0x63, 0x61, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x41, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x61, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x47,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x61, 0x74,
	0x65, 0x22, 0x2f, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x41, 0x0a, 0x12, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa9, 0x02, 0x0a, 0x0b, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x63, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x6d, 0x65, 0x63, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x22, 0x2f, 0x0a, 0x13, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x49, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72,
	0x6f, 0x6d, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf9, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x56, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2a, 0x2b, 0x0a, 0x0e, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x49, 0x54, 0x48,
	0x44, 0x52, 0x41, 0x57, 0x10, 0x01, 0x32, 0x5a, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x6d, 0x65,
	0x63, 0x61, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinmeca_grpc_proto_rawDescOnce sync.Once
	file_coinmeca_grpc_proto_rawDescData = file_coinmeca_grpc_proto_rawDesc
)

func file_coinmeca_grpc_proto_rawDescGZIP() []byte {
	file_coinmeca_grpc_proto_rawDescOnce.Do(func() {
		file_coinmeca_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinmeca_grpc_proto_rawDescData)
	})
	return file_coinmeca_grpc_proto_rawDescData
}

var file_coinmeca_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coinmeca_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_coinmeca_grpc_proto_goTypes = []interface{}{
	(VaultTradeType)(0),            // 0: grpcmodule.VaultTradeType
	(*GeneralRequest)(nil),         // 1: grpcmodule.GeneralRequest
	(*GeneralResponse)(nil),        // 2: grpcmodule.GeneralResponse
	(*TransactionRequest)(nil),     // 3: grpcmodule.TransactionRequest
	(*Transaction)(nil),            // 4: grpcmodule.Transaction
	(*TransactionResponse)(nil),    // 5: grpcmodule.TransactionResponse
	(*VaultRecentRequest)(nil),     // 6: grpcmodule.VaultRecentRequest
	(*VaultRecent)(nil),            // 7: grpcmodule.VaultRecent
	(*VaultRecentResponse)(nil),    // 8: grpcmodule.VaultRecentResponse
	(*GetAllFromVaultRequest)(nil), // 9: grpcmodule.GetAllFromVaultRequest
	(*GetAllFromVault)(nil),        // 10: grpcmodule.GetAllFromVault
}
var file_coinmeca_grpc_proto_depIdxs = []int32{
	3,  // 0: grpcmodule.GeneralRequest.transaction:type_name -> grpcmodule.TransactionRequest
	6,  // 1: grpcmodule.GeneralRequest.vaultRecent:type_name -> grpcmodule.VaultRecentRequest
	4,  // 2: grpcmodule.TransactionRequest.data:type_name -> grpcmodule.Transaction
	7,  // 3: grpcmodule.VaultRecentRequest.data:type_name -> grpcmodule.VaultRecent
	0,  // 4: grpcmodule.VaultRecent.type:type_name -> grpcmodule.VaultTradeType
	10, // 5: grpcmodule.GetAllFromVaultRequest.data:type_name -> grpcmodule.GetAllFromVault
	1,  // 6: grpcmodule.CoinmecaGrpcModule.AddData:input_type -> grpcmodule.GeneralRequest
	2,  // 7: grpcmodule.CoinmecaGrpcModule.AddData:output_type -> grpcmodule.GeneralResponse
	7,  // [7:8] is the sub-list for method output_type
	6,  // [6:7] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_coinmeca_grpc_proto_init() }
func file_coinmeca_grpc_proto_init() {
	if File_coinmeca_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinmeca_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultRecentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultRecent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultRecentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFromVaultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinmeca_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFromVault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_coinmeca_grpc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GeneralRequest_Transaction)(nil),
		(*GeneralRequest_VaultRecent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinmeca_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coinmeca_grpc_proto_goTypes,
		DependencyIndexes: file_coinmeca_grpc_proto_depIdxs,
		EnumInfos:         file_coinmeca_grpc_proto_enumTypes,
		MessageInfos:      file_coinmeca_grpc_proto_msgTypes,
	}.Build()
	File_coinmeca_grpc_proto = out.File
	file_coinmeca_grpc_proto_rawDesc = nil
	file_coinmeca_grpc_proto_goTypes = nil
	file_coinmeca_grpc_proto_depIdxs = nil
}
