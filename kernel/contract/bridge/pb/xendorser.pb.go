// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xendorser.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 请求参数
type EndorserRequest struct {
	Header               *Header      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	RequestName          string       `protobuf:"bytes,2,opt,name=RequestName,proto3" json:"RequestName,omitempty"`
	BcName               string       `protobuf:"bytes,3,opt,name=BcName,proto3" json:"BcName,omitempty"`
	Fee                  *Transaction `protobuf:"bytes,4,opt,name=Fee,proto3" json:"Fee,omitempty"`
	RequestData          []byte       `protobuf:"bytes,5,opt,name=RequestData,proto3" json:"RequestData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EndorserRequest) Reset()         { *m = EndorserRequest{} }
func (m *EndorserRequest) String() string { return proto.CompactTextString(m) }
func (*EndorserRequest) ProtoMessage()    {}
func (*EndorserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf870ebd3b57e1, []int{0}
}

func (m *EndorserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorserRequest.Unmarshal(m, b)
}
func (m *EndorserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorserRequest.Marshal(b, m, deterministic)
}
func (m *EndorserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorserRequest.Merge(m, src)
}
func (m *EndorserRequest) XXX_Size() int {
	return xxx_messageInfo_EndorserRequest.Size(m)
}
func (m *EndorserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndorserRequest proto.InternalMessageInfo

func (m *EndorserRequest) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EndorserRequest) GetRequestName() string {
	if m != nil {
		return m.RequestName
	}
	return ""
}

func (m *EndorserRequest) GetBcName() string {
	if m != nil {
		return m.BcName
	}
	return ""
}

func (m *EndorserRequest) GetFee() *Transaction {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *EndorserRequest) GetRequestData() []byte {
	if m != nil {
		return m.RequestData
	}
	return nil
}

type EndorserResponse struct {
	Header               *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ResponseName         string         `protobuf:"bytes,2,opt,name=ResponseName,proto3" json:"ResponseName,omitempty"`
	EndorserAddress      string         `protobuf:"bytes,3,opt,name=EndorserAddress,proto3" json:"EndorserAddress,omitempty"`
	EndorserSign         *SignatureInfo `protobuf:"bytes,4,opt,name=EndorserSign,proto3" json:"EndorserSign,omitempty"`
	ResponseData         []byte         `protobuf:"bytes,5,opt,name=ResponseData,proto3" json:"ResponseData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EndorserResponse) Reset()         { *m = EndorserResponse{} }
func (m *EndorserResponse) String() string { return proto.CompactTextString(m) }
func (*EndorserResponse) ProtoMessage()    {}
func (*EndorserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeaf870ebd3b57e1, []int{1}
}

func (m *EndorserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorserResponse.Unmarshal(m, b)
}
func (m *EndorserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorserResponse.Marshal(b, m, deterministic)
}
func (m *EndorserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorserResponse.Merge(m, src)
}
func (m *EndorserResponse) XXX_Size() int {
	return xxx_messageInfo_EndorserResponse.Size(m)
}
func (m *EndorserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndorserResponse proto.InternalMessageInfo

func (m *EndorserResponse) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EndorserResponse) GetResponseName() string {
	if m != nil {
		return m.ResponseName
	}
	return ""
}

func (m *EndorserResponse) GetEndorserAddress() string {
	if m != nil {
		return m.EndorserAddress
	}
	return ""
}

func (m *EndorserResponse) GetEndorserSign() *SignatureInfo {
	if m != nil {
		return m.EndorserSign
	}
	return nil
}

func (m *EndorserResponse) GetResponseData() []byte {
	if m != nil {
		return m.ResponseData
	}
	return nil
}

func init() {
	// proto.RegisterType((*EndorserRequest)(nil), "pb.EndorserRequest")
	// proto.RegisterType((*EndorserResponse)(nil), "pb.EndorserResponse")
}

// func init() { proto.RegisterFile("xendorser.proto", fileDescriptor_eeaf870ebd3b57e1) }

var fileDescriptor_eeaf870ebd3b57e1 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xe5, 0x16, 0x2a, 0xd5, 0x8d, 0xd4, 0x62, 0xfe, 0x28, 0x2a, 0x0c, 0x21, 0x53, 0xc4,
	0xd0, 0x88, 0x22, 0x16, 0x36, 0xfe, 0x0a, 0x16, 0x86, 0x80, 0xd8, 0xaf, 0xc9, 0xd1, 0x46, 0x0a,
	0x76, 0xb0, 0x5d, 0xd4, 0x99, 0x57, 0xe0, 0x55, 0x78, 0x13, 0x1e, 0x80, 0x85, 0x07, 0x41, 0x4e,
	0x5d, 0xe2, 0x74, 0x62, 0xbc, 0xdf, 0x77, 0xfa, 0xee, 0x3b, 0x9f, 0x69, 0x7f, 0x81, 0x3c, 0x13,
	0x52, 0xa1, 0x1c, 0x95, 0x52, 0x68, 0xc1, 0x5a, 0xe5, 0x64, 0xe8, 0x2d, 0xd2, 0x19, 0xe4, 0x7c,
	0x49, 0x86, 0x07, 0x53, 0x21, 0xa6, 0x05, 0xc6, 0x50, 0xe6, 0x31, 0x70, 0x2e, 0x34, 0xe8, 0x5c,
	0x70, 0xb5, 0x54, 0xc3, 0x4f, 0x42, 0xfb, 0xd7, 0xd6, 0x22, 0xc1, 0xd7, 0x39, 0x2a, 0xcd, 0x42,
	0xda, 0x99, 0x21, 0x64, 0x28, 0x7d, 0x12, 0x90, 0xa8, 0x37, 0xa6, 0xa3, 0x72, 0x32, 0xba, 0xad,
	0x48, 0x62, 0x15, 0x16, 0xd0, 0x9e, 0x6d, 0xbf, 0x87, 0x17, 0xf4, 0x5b, 0x01, 0x89, 0xba, 0x89,
	0x8b, 0xd8, 0x1e, 0xed, 0x5c, 0xa4, 0x95, 0xd8, 0xae, 0x44, 0x5b, 0xb1, 0x43, 0xda, 0xbe, 0x41,
	0xf4, 0x37, 0x2a, 0xeb, 0xbe, 0xb1, 0x7e, 0x94, 0xc0, 0x15, 0xa4, 0x26, 0x56, 0x62, 0x34, 0xc7,
	0xfc, 0x0a, 0x34, 0xf8, 0x9b, 0x01, 0x89, 0xbc, 0xc4, 0x45, 0xe1, 0x37, 0xa1, 0x83, 0x3a, 0xb6,
	0x2a, 0x05, 0x57, 0xf8, 0xaf, 0xdc, 0x21, 0xf5, 0x56, 0xfd, 0x4e, 0xf0, 0x06, 0x63, 0x51, 0xfd,
	0x24, 0xe7, 0x59, 0x26, 0x51, 0x29, 0xbb, 0xc2, 0x3a, 0x66, 0xa7, 0xd4, 0x5b, 0xa1, 0x87, 0x7c,
	0xca, 0xed, 0x52, 0x5b, 0x66, 0xae, 0xa9, 0x41, 0xcf, 0x25, 0xde, 0xf1, 0x67, 0x91, 0x34, 0xda,
	0xdc, 0x10, 0xce, 0x82, 0x0d, 0x36, 0x4e, 0x69, 0xf7, 0xef, 0xb6, 0xec, 0xa9, 0x9e, 0x73, 0x09,
	0x45, 0xc1, 0xb6, 0xcd, 0x84, 0xb5, 0xb3, 0x0d, 0x77, 0x9a, 0x70, 0xe9, 0x15, 0xee, 0xbf, 0x7f,
	0xfd, 0x7c, 0xb4, 0x76, 0xc3, 0x41, 0xfc, 0x76, 0x1c, 0xaf, 0x0c, 0x53, 0x28, 0x8a, 0x33, 0x72,
	0x34, 0xe9, 0x54, 0x9f, 0xe0, 0xe4, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x40, 0x0c, 0x04, 0x47,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// XendorserClient is the client API for Xendorser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type XendorserClient interface {
	EndorserCall(ctx context.Context, in *EndorserRequest, opts ...grpc.CallOption) (*EndorserResponse, error)
}

type xendorserClient struct {
	cc grpc.ClientConnInterface
}

func NewXendorserClient(cc grpc.ClientConnInterface) XendorserClient {
	return &xendorserClient{cc}
}

func (c *xendorserClient) EndorserCall(ctx context.Context, in *EndorserRequest, opts ...grpc.CallOption) (*EndorserResponse, error) {
	out := new(EndorserResponse)
	err := c.cc.Invoke(ctx, "/pb.xendorser/EndorserCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XendorserServer is the server API for Xendorser service.
type XendorserServer interface {
	EndorserCall(context.Context, *EndorserRequest) (*EndorserResponse, error)
}

// UnimplementedXendorserServer can be embedded to have forward compatible implementations.
type UnimplementedXendorserServer struct {
}

func (*UnimplementedXendorserServer) EndorserCall(ctx context.Context, req *EndorserRequest) (*EndorserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndorserCall not implemented")
}

func RegisterXendorserServer(s *grpc.Server, srv XendorserServer) {
	s.RegisterService(&_Xendorser_serviceDesc, srv)
}

func _Xendorser_EndorserCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndorserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XendorserServer).EndorserCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.xendorser/EndorserCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XendorserServer).EndorserCall(ctx, req.(*EndorserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Xendorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.xendorser",
	HandlerType: (*XendorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EndorserCall",
			Handler:    _Xendorser_EndorserCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xendorser.proto",
}
