// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package pingtunnel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MyMsg_TYPE int32

const (
	MyMsg_DATA MyMsg_TYPE = 0
	MyMsg_PING MyMsg_TYPE = 1
)

var MyMsg_TYPE_name = map[int32]string{
	0: "DATA",
	1: "PING",
}

var MyMsg_TYPE_value = map[string]int32{
	"DATA": 0,
	"PING": 1,
}

func (x MyMsg_TYPE) String() string {
	return proto.EnumName(MyMsg_TYPE_name, int32(x))
}

func (MyMsg_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0, 0}
}

type Frame_TYPE int32

const (
	Frame_DATA Frame_TYPE = 0
	Frame_REQ  Frame_TYPE = 1
	Frame_ACK  Frame_TYPE = 2
)

var Frame_TYPE_name = map[int32]string{
	0: "DATA",
	1: "REQ",
	2: "ACK",
}

var Frame_TYPE_value = map[string]int32{
	"DATA": 0,
	"REQ":  1,
	"ACK":  2,
}

func (x Frame_TYPE) String() string {
	return proto.EnumName(Frame_TYPE_name, int32(x))
}

func (Frame_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1, 0}
}

type MyMsg struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Rproto               int32    `protobuf:"varint,5,opt,name=rproto,proto3" json:"rproto,omitempty"`
	Key                  int32    `protobuf:"varint,7,opt,name=key,proto3" json:"key,omitempty"`
	Tcpmode              int32    `protobuf:"varint,8,opt,name=tcpmode,proto3" json:"tcpmode,omitempty"`
	TcpmodeBuffersize    int32    `protobuf:"varint,9,opt,name=tcpmode_buffersize,json=tcpmodeBuffersize,proto3" json:"tcpmode_buffersize,omitempty"`
	TcpmodeMaxwin        int32    `protobuf:"varint,10,opt,name=tcpmode_maxwin,json=tcpmodeMaxwin,proto3" json:"tcpmode_maxwin,omitempty"`
	TcpmodeResendTimems  int32    `protobuf:"varint,11,opt,name=tcpmode_resend_timems,json=tcpmodeResendTimems,proto3" json:"tcpmode_resend_timems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMsg) Reset()         { *m = MyMsg{} }
func (m *MyMsg) String() string { return proto.CompactTextString(m) }
func (*MyMsg) ProtoMessage()    {}
func (*MyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *MyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMsg.Unmarshal(m, b)
}
func (m *MyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMsg.Marshal(b, m, deterministic)
}
func (m *MyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMsg.Merge(m, src)
}
func (m *MyMsg) XXX_Size() int {
	return xxx_messageInfo_MyMsg.Size(m)
}
func (m *MyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MyMsg proto.InternalMessageInfo

func (m *MyMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MyMsg) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MyMsg) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *MyMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MyMsg) GetRproto() int32 {
	if m != nil {
		return m.Rproto
	}
	return 0
}

func (m *MyMsg) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *MyMsg) GetTcpmode() int32 {
	if m != nil {
		return m.Tcpmode
	}
	return 0
}

func (m *MyMsg) GetTcpmodeBuffersize() int32 {
	if m != nil {
		return m.TcpmodeBuffersize
	}
	return 0
}

func (m *MyMsg) GetTcpmodeMaxwin() int32 {
	if m != nil {
		return m.TcpmodeMaxwin
	}
	return 0
}

func (m *MyMsg) GetTcpmodeResendTimems() int32 {
	if m != nil {
		return m.TcpmodeResendTimems
	}
	return 0
}

type Frame struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Resend               bool     `protobuf:"varint,2,opt,name=resend,proto3" json:"resend,omitempty"`
	Sendtime             int64    `protobuf:"varint,3,opt,name=sendtime,proto3" json:"sendtime,omitempty"`
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Dataid               []int32  `protobuf:"varint,6,rep,packed,name=dataid,proto3" json:"dataid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Frame) GetResend() bool {
	if m != nil {
		return m.Resend
	}
	return false
}

func (m *Frame) GetSendtime() int64 {
	if m != nil {
		return m.Sendtime
	}
	return 0
}

func (m *Frame) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Frame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Frame) GetDataid() []int32 {
	if m != nil {
		return m.Dataid
	}
	return nil
}

func init() {
	proto.RegisterEnum("MyMsg_TYPE", MyMsg_TYPE_name, MyMsg_TYPE_value)
	proto.RegisterEnum("Frame_TYPE", Frame_TYPE_name, Frame_TYPE_value)
	proto.RegisterType((*MyMsg)(nil), "MyMsg")
	proto.RegisterType((*Frame)(nil), "Frame")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x76, 0xf3, 0xd3, 0xa6, 0x63, 0x2d, 0x71, 0x45, 0x19, 0x7a, 0x0a, 0x01, 0xa1, 0x17, 0x3d,
	0xe8, 0x13, 0xb4, 0x5a, 0x45, 0xa4, 0x52, 0x97, 0x5e, 0xf4, 0x52, 0x52, 0xb3, 0x0d, 0x8b, 0x26,
	0x0d, 0xc9, 0x16, 0xad, 0x2f, 0xe5, 0x33, 0xf8, 0x66, 0xb2, 0xd3, 0x6d, 0xa0, 0x78, 0xda, 0xef,
	0x8f, 0x99, 0xe5, 0x1b, 0xe8, 0xe4, 0x75, 0x76, 0x59, 0x56, 0x2b, 0xbd, 0x8a, 0x7f, 0x1d, 0xf0,
	0x27, 0x9b, 0x49, 0x9d, 0xf1, 0x1e, 0x38, 0x2a, 0x45, 0x16, 0xb1, 0x41, 0x47, 0x38, 0x2a, 0xe5,
	0x1c, 0x3c, 0xbd, 0x29, 0x25, 0x3a, 0x11, 0x1b, 0xf8, 0x82, 0x30, 0x3f, 0x83, 0x96, 0x4e, 0xaa,
	0x4c, 0x6a, 0x74, 0x29, 0x67, 0x99, 0xc9, 0xa6, 0x89, 0x4e, 0xd0, 0x8b, 0xd8, 0xa0, 0x2b, 0x08,
	0x9b, 0x6c, 0x45, 0x3b, 0xd0, 0xa7, 0x09, 0x96, 0xf1, 0x10, 0xdc, 0x77, 0xb9, 0xc1, 0x36, 0x89,
	0x06, 0x72, 0x84, 0xb6, 0x7e, 0x2b, 0xf3, 0x55, 0x2a, 0x31, 0x20, 0x75, 0x47, 0xf9, 0x05, 0x70,
	0x0b, 0xe7, 0x8b, 0xf5, 0x72, 0x29, 0xab, 0x5a, 0x7d, 0x4b, 0xec, 0x50, 0xe8, 0xd8, 0x3a, 0xa3,
	0xc6, 0xe0, 0xe7, 0xd0, 0xdb, 0xc5, 0xf3, 0xe4, 0xeb, 0x53, 0x15, 0x08, 0x14, 0x3d, 0xb2, 0xea,
	0x84, 0x44, 0x7e, 0x05, 0xa7, 0xbb, 0x58, 0x25, 0x6b, 0x59, 0xa4, 0x73, 0xad, 0x72, 0x99, 0xd7,
	0x78, 0x48, 0xe9, 0x13, 0x6b, 0x0a, 0xf2, 0x66, 0x64, 0xc5, 0x7d, 0xf0, 0x66, 0x2f, 0xd3, 0x31,
	0x0f, 0xc0, 0xbb, 0x1d, 0xce, 0x86, 0xe1, 0x81, 0x41, 0xd3, 0x87, 0xa7, 0xfb, 0x90, 0xc5, 0x3f,
	0x0c, 0xfc, 0xbb, 0x2a, 0xc9, 0x65, 0xd3, 0x19, 0xdb, 0xef, 0x6c, 0xbb, 0x85, 0x9a, 0x0c, 0x84,
	0x65, 0xbc, 0x0f, 0x81, 0x79, 0xcd, 0x6a, 0x6a, 0xd3, 0x15, 0x0d, 0xb7, 0xb7, 0xf0, 0x68, 0x8a,
	0xbd, 0x05, 0xf5, 0xeb, 0xef, 0xf7, 0x6b, 0x5e, 0x95, 0x62, 0x2b, 0x72, 0x4d, 0xbf, 0x5b, 0x16,
	0xc7, 0xff, 0x7e, 0xda, 0x06, 0x57, 0x8c, 0x9f, 0x43, 0x66, 0xc0, 0xf0, 0xe6, 0x31, 0x74, 0x46,
	0xdd, 0x57, 0x28, 0x55, 0x91, 0xe9, 0x75, 0x51, 0xc8, 0x8f, 0x45, 0x8b, 0x0e, 0x73, 0xfd, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0xfe, 0x5f, 0xc5, 0x17, 0x02, 0x00, 0x00,
}