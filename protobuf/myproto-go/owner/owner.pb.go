package owner

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

// Owner 资产拥有者
type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId   int64  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerDesc string `protobuf:"bytes,16,opt,name=owner_desc,json=ownerDesc,proto3" json:"owner_desc,omitempty"`
	OwnerName string `protobuf:"bytes,2,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_owner_owner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_owner_owner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_owner_owner_proto_rawDescGZIP(), []int{0}
}

func (x *Owner) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Owner) GetOwnerDesc() string {
	if x != nil {
		return x.OwnerDesc
	}
	return ""
}

func (x *Owner) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

var File_owner_owner_proto protoreflect.FileDescriptor

var file_owner_owner_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x05, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61, 0x6f, 0x63,
	0x68, 0x75, 0x6e, 0x69, 0x6e, 0x68, 0x65, 0x66, 0x65, 0x69, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_owner_owner_proto_rawDescOnce sync.Once
	file_owner_owner_proto_rawDescData = file_owner_owner_proto_rawDesc
)

func file_owner_owner_proto_rawDescGZIP() []byte {
	file_owner_owner_proto_rawDescOnce.Do(func() {
		file_owner_owner_proto_rawDescData = protoimpl.X.CompressGZIP(file_owner_owner_proto_rawDescData)
	})
	return file_owner_owner_proto_rawDescData
}

var file_owner_owner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_owner_owner_proto_goTypes = []interface{}{
	(*Owner)(nil), // 0: owner.Owner
}
var file_owner_owner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_owner_owner_proto_init() }
func file_owner_owner_proto_init() {
	if File_owner_owner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_owner_owner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
			RawDescriptor: file_owner_owner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_owner_owner_proto_goTypes,
		DependencyIndexes: file_owner_owner_proto_depIdxs,
		MessageInfos:      file_owner_owner_proto_msgTypes,
	}.Build()
	File_owner_owner_proto = out.File
	file_owner_owner_proto_rawDesc = nil
	file_owner_owner_proto_goTypes = nil
	file_owner_owner_proto_depIdxs = nil
}
