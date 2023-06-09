// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: logstream1.proto

package logstream1_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamLogEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *StreamLogEventsRequest) Reset() {
	*x = StreamLogEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLogEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLogEventsRequest) ProtoMessage() {}

func (x *StreamLogEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logstream1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLogEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamLogEventsRequest) Descriptor() ([]byte, []int) {
	return file_logstream1_proto_rawDescGZIP(), []int{0}
}

func (x *StreamLogEventsRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type LogEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*LogEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *LogEvents) Reset() {
	*x = LogEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvents) ProtoMessage() {}

func (x *LogEvents) ProtoReflect() protoreflect.Message {
	mi := &file_logstream1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvents.ProtoReflect.Descriptor instead.
func (*LogEvents) Descriptor() ([]byte, []int) {
	return file_logstream1_proto_rawDescGZIP(), []int{1}
}

func (x *LogEvents) GetEvents() []*LogEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg       string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_logstream1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_logstream1_proto_rawDescGZIP(), []int{2}
}

func (x *LogEvent) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogEvent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LogEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_logstream1_proto protoreflect.FileDescriptor

var file_logstream1_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c,
	0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x5e,
	0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logstream1_proto_rawDescOnce sync.Once
	file_logstream1_proto_rawDescData = file_logstream1_proto_rawDesc
)

func file_logstream1_proto_rawDescGZIP() []byte {
	file_logstream1_proto_rawDescOnce.Do(func() {
		file_logstream1_proto_rawDescData = protoimpl.X.CompressGZIP(file_logstream1_proto_rawDescData)
	})
	return file_logstream1_proto_rawDescData
}

var file_logstream1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_logstream1_proto_goTypes = []interface{}{
	(*StreamLogEventsRequest)(nil), // 0: logevents.StreamLogEventsRequest
	(*LogEvents)(nil),              // 1: logevents.LogEvents
	(*LogEvent)(nil),               // 2: logevents.LogEvent
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_logstream1_proto_depIdxs = []int32{
	2, // 0: logevents.LogEvents.events:type_name -> logevents.LogEvent
	3, // 1: logevents.LogEvent.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: logevents.LogEventService.StreamLogEvents:input_type -> logevents.StreamLogEventsRequest
	2, // 3: logevents.LogEventService.StreamLogEvents:output_type -> logevents.LogEvent
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logstream1_proto_init() }
func file_logstream1_proto_init() {
	if File_logstream1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logstream1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamLogEventsRequest); i {
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
		file_logstream1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEvents); i {
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
		file_logstream1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEvent); i {
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
			RawDescriptor: file_logstream1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logstream1_proto_goTypes,
		DependencyIndexes: file_logstream1_proto_depIdxs,
		MessageInfos:      file_logstream1_proto_msgTypes,
	}.Build()
	File_logstream1_proto = out.File
	file_logstream1_proto_rawDesc = nil
	file_logstream1_proto_goTypes = nil
	file_logstream1_proto_depIdxs = nil
}
