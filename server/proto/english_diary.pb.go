// Code generated by protoc-gen-go. DO NOT EDIT.
// source: english_diary.proto

package english_diary

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Diary struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Diary) Reset()         { *m = Diary{} }
func (m *Diary) String() string { return proto.CompactTextString(m) }
func (*Diary) ProtoMessage()    {}
func (*Diary) Descriptor() ([]byte, []int) {
	return fileDescriptor_english_diary_4ac71dc90fc8c34a, []int{0}
}
func (m *Diary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Diary.Unmarshal(m, b)
}
func (m *Diary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Diary.Marshal(b, m, deterministic)
}
func (dst *Diary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Diary.Merge(dst, src)
}
func (m *Diary) XXX_Size() int {
	return xxx_messageInfo_Diary.Size(m)
}
func (m *Diary) XXX_DiscardUnknown() {
	xxx_messageInfo_Diary.DiscardUnknown(m)
}

var xxx_messageInfo_Diary proto.InternalMessageInfo

func (m *Diary) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Diary) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Diary) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type GetDiaryRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiaryRequest) Reset()         { *m = GetDiaryRequest{} }
func (m *GetDiaryRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiaryRequest) ProtoMessage()    {}
func (*GetDiaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_english_diary_4ac71dc90fc8c34a, []int{1}
}
func (m *GetDiaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiaryRequest.Unmarshal(m, b)
}
func (m *GetDiaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiaryRequest.Marshal(b, m, deterministic)
}
func (dst *GetDiaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiaryRequest.Merge(dst, src)
}
func (m *GetDiaryRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiaryRequest.Size(m)
}
func (m *GetDiaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiaryRequest proto.InternalMessageInfo

func (m *GetDiaryRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Diary)(nil), "english_diary.Diary")
	proto.RegisterType((*GetDiaryRequest)(nil), "english_diary.GetDiaryRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EnglishDiaryClient is the client API for EnglishDiary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnglishDiaryClient interface {
	GetDiary(ctx context.Context, in *GetDiaryRequest, opts ...grpc.CallOption) (*Diary, error)
}

type englishDiaryClient struct {
	cc *grpc.ClientConn
}

func NewEnglishDiaryClient(cc *grpc.ClientConn) EnglishDiaryClient {
	return &englishDiaryClient{cc}
}

func (c *englishDiaryClient) GetDiary(ctx context.Context, in *GetDiaryRequest, opts ...grpc.CallOption) (*Diary, error) {
	out := new(Diary)
	err := c.cc.Invoke(ctx, "/english_diary.EnglishDiary/GetDiary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnglishDiaryServer is the server API for EnglishDiary service.
type EnglishDiaryServer interface {
	GetDiary(context.Context, *GetDiaryRequest) (*Diary, error)
}

func RegisterEnglishDiaryServer(s *grpc.Server, srv EnglishDiaryServer) {
	s.RegisterService(&_EnglishDiary_serviceDesc, srv)
}

func _EnglishDiary_GetDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnglishDiaryServer).GetDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/english_diary.EnglishDiary/GetDiary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnglishDiaryServer).GetDiary(ctx, req.(*GetDiaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnglishDiary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "english_diary.EnglishDiary",
	HandlerType: (*EnglishDiaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiary",
			Handler:    _EnglishDiary_GetDiary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "english_diary.proto",
}

func init() { proto.RegisterFile("english_diary.proto", fileDescriptor_english_diary_4ac71dc90fc8c34a) }

var fileDescriptor_english_diary_4ac71dc90fc8c34a = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcd, 0x4b, 0xcf,
	0xc9, 0x2c, 0xce, 0x88, 0x4f, 0xc9, 0x4c, 0x2c, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0x72, 0xe4, 0x62, 0x75, 0x01, 0x31, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0x32,
	0x4b, 0x72, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96,
	0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x66, 0xb0, 0x20, 0x98, 0xad, 0xa4, 0xc8, 0xc5, 0xef, 0x9e, 0x5a,
	0x02, 0x36, 0x25, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x04, 0xdd, 0x30, 0xa3, 0x20, 0x2e, 0x1e,
	0x57, 0x88, 0xb5, 0x10, 0xcb, 0x9c, 0xb8, 0x38, 0x60, 0x5a, 0x84, 0xe4, 0xf4, 0x50, 0x9d, 0x89,
	0x66, 0x96, 0x94, 0x08, 0x9a, 0x3c, 0x58, 0x52, 0x89, 0xc1, 0x89, 0x3f, 0x0a, 0xd5, 0x2b, 0x49,
	0x6c, 0x60, 0x0f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x7b, 0x14, 0x8a, 0xf7, 0x00,
	0x00, 0x00,
}