// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: queue.proto

package types

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

// Idea: add loadtype, and it will processed by any queue that supports that loadtype.
type QueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant string       `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Queue  string       `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Items  []*QueueItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *QueueRequest) Reset() {
	*x = QueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueRequest) ProtoMessage() {}

func (x *QueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueRequest.ProtoReflect.Descriptor instead.
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{0}
}

func (x *QueueRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *QueueRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *QueueRequest) GetItems() []*QueueItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DequeueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant string       `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Queue  string       `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Items  []*QueueItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DequeueRequest) Reset() {
	*x = DequeueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueRequest) ProtoMessage() {}

func (x *DequeueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueRequest.ProtoReflect.Descriptor instead.
func (*DequeueRequest) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{1}
}

func (x *DequeueRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *DequeueRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *DequeueRequest) GetItems() []*QueueItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DequeueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*QueueItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DequeueReply) Reset() {
	*x = DequeueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueReply) ProtoMessage() {}

func (x *DequeueReply) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueReply.ProtoReflect.Descriptor instead.
func (*DequeueReply) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{2}
}

func (x *DequeueReply) GetItems() []*QueueItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueueItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueueItem) Reset() {
	*x = QueueItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueItem) ProtoMessage() {}

func (x *QueueItem) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueItem.ProtoReflect.Descriptor instead.
func (*QueueItem) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{3}
}

func (x *QueueItem) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *QueueItem) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_queue_proto protoreflect.FileDescriptor

var file_queue_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0c, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x60, 0x0a, 0x0e, 0x44,
	0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x30, 0x0a,
	0x0c, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x35, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x55, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x1f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0d, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x2b, 0x0a, 0x07, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0f, 0x2e, 0x44, 0x65,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x44,
	0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_proto_rawDescOnce sync.Once
	file_queue_proto_rawDescData = file_queue_proto_rawDesc
)

func file_queue_proto_rawDescGZIP() []byte {
	file_queue_proto_rawDescOnce.Do(func() {
		file_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_proto_rawDescData)
	})
	return file_queue_proto_rawDescData
}

var file_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_queue_proto_goTypes = []interface{}{
	(*QueueRequest)(nil),   // 0: QueueRequest
	(*DequeueRequest)(nil), // 1: DequeueRequest
	(*DequeueReply)(nil),   // 2: DequeueReply
	(*QueueItem)(nil),      // 3: QueueItem
	(*Void)(nil),           // 4: Void
}
var file_queue_proto_depIdxs = []int32{
	3, // 0: QueueRequest.items:type_name -> QueueItem
	3, // 1: DequeueRequest.items:type_name -> QueueItem
	3, // 2: DequeueReply.items:type_name -> QueueItem
	0, // 3: Queue.Queue:input_type -> QueueRequest
	1, // 4: Queue.Dequeue:input_type -> DequeueRequest
	4, // 5: Queue.Queue:output_type -> Void
	2, // 6: Queue.Dequeue:output_type -> DequeueReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_queue_proto_init() }
func file_queue_proto_init() {
	if File_queue_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueRequest); i {
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
		file_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueRequest); i {
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
		file_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueReply); i {
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
		file_queue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueItem); i {
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
			RawDescriptor: file_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queue_proto_goTypes,
		DependencyIndexes: file_queue_proto_depIdxs,
		MessageInfos:      file_queue_proto_msgTypes,
	}.Build()
	File_queue_proto = out.File
	file_queue_proto_rawDesc = nil
	file_queue_proto_goTypes = nil
	file_queue_proto_depIdxs = nil
}
