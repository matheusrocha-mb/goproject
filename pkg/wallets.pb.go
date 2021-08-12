// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pkg/wallets.proto

package wlt

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

type WalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string         `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Wallet []*ModelWallet `protobuf:"bytes,2,rep,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *WalletRequest) Reset() {
	*x = WalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wallets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletRequest) ProtoMessage() {}

func (x *WalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wallets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletRequest.ProtoReflect.Descriptor instead.
func (*WalletRequest) Descriptor() ([]byte, []int) {
	return file_pkg_wallets_proto_rawDescGZIP(), []int{0}
}

func (x *WalletRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WalletRequest) GetWallet() []*ModelWallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type WalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet []*ModelWalletReply `protobuf:"bytes,1,rep,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *WalletReply) Reset() {
	*x = WalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wallets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletReply) ProtoMessage() {}

func (x *WalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wallets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletReply.ProtoReflect.Descriptor instead.
func (*WalletReply) Descriptor() ([]byte, []int) {
	return file_pkg_wallets_proto_rawDescGZIP(), []int{1}
}

func (x *WalletReply) GetWallet() []*ModelWalletReply {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type ModelWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName  string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	AssetTag  string `protobuf:"bytes,4,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	AssetName string `protobuf:"bytes,5,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	Hash      string `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	Balance   string `protobuf:"bytes,7,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ModelWallet) Reset() {
	*x = ModelWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wallets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelWallet) ProtoMessage() {}

func (x *ModelWallet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wallets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelWallet.ProtoReflect.Descriptor instead.
func (*ModelWallet) Descriptor() ([]byte, []int) {
	return file_pkg_wallets_proto_rawDescGZIP(), []int{2}
}

func (x *ModelWallet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModelWallet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ModelWallet) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ModelWallet) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *ModelWallet) GetAssetName() string {
	if x != nil {
		return x.AssetName
	}
	return ""
}

func (x *ModelWallet) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ModelWallet) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *ModelWallet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ModelWallet) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ModelWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         *ModelWallet `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	RowsAffected string       `protobuf:"bytes,3,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	Error        string       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ModelWalletReply) Reset() {
	*x = ModelWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wallets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelWalletReply) ProtoMessage() {}

func (x *ModelWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wallets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelWalletReply.ProtoReflect.Descriptor instead.
func (*ModelWalletReply) Descriptor() ([]byte, []int) {
	return file_pkg_wallets_proto_rawDescGZIP(), []int{3}
}

func (x *ModelWalletReply) GetData() *ModelWallet {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ModelWalletReply) GetRowsAffected() string {
	if x != nil {
		return x.RowsAffected
	}
	return ""
}

func (x *ModelWalletReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_wallets_proto protoreflect.FileDescriptor

var file_pkg_wallets_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x33, 0x0a, 0x06,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x22, 0x47, 0x0a, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x38, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x0b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7e, 0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x5e, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x3b, 0x77,
	0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_wallets_proto_rawDescOnce sync.Once
	file_pkg_wallets_proto_rawDescData = file_pkg_wallets_proto_rawDesc
)

func file_pkg_wallets_proto_rawDescGZIP() []byte {
	file_pkg_wallets_proto_rawDescOnce.Do(func() {
		file_pkg_wallets_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_wallets_proto_rawDescData)
	})
	return file_pkg_wallets_proto_rawDescData
}

var file_pkg_wallets_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_wallets_proto_goTypes = []interface{}{
	(*WalletRequest)(nil),    // 0: wallet_manager.WalletRequest
	(*WalletReply)(nil),      // 1: wallet_manager.WalletReply
	(*ModelWallet)(nil),      // 2: wallet_manager.modelWallet
	(*ModelWalletReply)(nil), // 3: wallet_manager.modelWalletReply
}
var file_pkg_wallets_proto_depIdxs = []int32{
	2, // 0: wallet_manager.WalletRequest.wallet:type_name -> wallet_manager.modelWallet
	3, // 1: wallet_manager.WalletReply.wallet:type_name -> wallet_manager.modelWalletReply
	2, // 2: wallet_manager.modelWalletReply.data:type_name -> wallet_manager.modelWallet
	0, // 3: wallet_manager.WalletManager.ManagerWallet:input_type -> wallet_manager.WalletRequest
	1, // 4: wallet_manager.WalletManager.ManagerWallet:output_type -> wallet_manager.WalletReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_wallets_proto_init() }
func file_pkg_wallets_proto_init() {
	if File_pkg_wallets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_wallets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletRequest); i {
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
		file_pkg_wallets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletReply); i {
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
		file_pkg_wallets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelWallet); i {
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
		file_pkg_wallets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelWalletReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_wallets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_wallets_proto_goTypes,
		DependencyIndexes: file_pkg_wallets_proto_depIdxs,
		MessageInfos:      file_pkg_wallets_proto_msgTypes,
	}.Build()
	File_pkg_wallets_proto = out.File
	file_pkg_wallets_proto_rawDesc = nil
	file_pkg_wallets_proto_goTypes = nil
	file_pkg_wallets_proto_depIdxs = nil
}
