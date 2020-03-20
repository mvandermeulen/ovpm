// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vpn.proto

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

type VPNProto int32

const (
	VPNProto_NOPREF VPNProto = 0
	VPNProto_UDP    VPNProto = 1
	VPNProto_TCP    VPNProto = 2
)

var VPNProto_name = map[int32]string{
	0: "NOPREF",
	1: "UDP",
	2: "TCP",
}

var VPNProto_value = map[string]int32{
	"NOPREF": 0,
	"UDP":    1,
	"TCP":    2,
}

func (x VPNProto) String() string {
	return proto.EnumName(VPNProto_name, int32(x))
}

func (VPNProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{0}
}

type VPNStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNStatusRequest) Reset()         { *m = VPNStatusRequest{} }
func (m *VPNStatusRequest) String() string { return proto.CompactTextString(m) }
func (*VPNStatusRequest) ProtoMessage()    {}
func (*VPNStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{0}
}

func (m *VPNStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNStatusRequest.Unmarshal(m, b)
}
func (m *VPNStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNStatusRequest.Marshal(b, m, deterministic)
}
func (m *VPNStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNStatusRequest.Merge(m, src)
}
func (m *VPNStatusRequest) XXX_Size() int {
	return xxx_messageInfo_VPNStatusRequest.Size(m)
}
func (m *VPNStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNStatusRequest proto.InternalMessageInfo

type VPNInitRequest struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ProtoPref            VPNProto `protobuf:"varint,3,opt,name=proto_pref,json=protoPref,proto3,enum=pb.VPNProto" json:"proto_pref,omitempty"`
	IpBlock              string   `protobuf:"bytes,4,opt,name=ip_block,json=ipBlock,proto3" json:"ip_block,omitempty"`
	Dns                  string   `protobuf:"bytes,5,opt,name=dns,proto3" json:"dns,omitempty"`
	KeepalivePeriod      string   `protobuf:"bytes,6,opt,name=keepalive_period,json=keepalivePeriod,proto3" json:"keepalive_period,omitempty"`
	KeepaliveTimeout     string   `protobuf:"bytes,7,opt,name=keepalive_timeout,json=keepaliveTimeout,proto3" json:"keepalive_timeout,omitempty"`
	UseLzo               bool     `protobuf:"varint,8,opt,name=use_lzo,json=useLzo,proto3" json:"use_lzo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNInitRequest) Reset()         { *m = VPNInitRequest{} }
func (m *VPNInitRequest) String() string { return proto.CompactTextString(m) }
func (*VPNInitRequest) ProtoMessage()    {}
func (*VPNInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{1}
}

func (m *VPNInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNInitRequest.Unmarshal(m, b)
}
func (m *VPNInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNInitRequest.Marshal(b, m, deterministic)
}
func (m *VPNInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNInitRequest.Merge(m, src)
}
func (m *VPNInitRequest) XXX_Size() int {
	return xxx_messageInfo_VPNInitRequest.Size(m)
}
func (m *VPNInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNInitRequest proto.InternalMessageInfo

func (m *VPNInitRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNInitRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNInitRequest) GetProtoPref() VPNProto {
	if m != nil {
		return m.ProtoPref
	}
	return VPNProto_NOPREF
}

func (m *VPNInitRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNInitRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

func (m *VPNInitRequest) GetKeepalivePeriod() string {
	if m != nil {
		return m.KeepalivePeriod
	}
	return ""
}

func (m *VPNInitRequest) GetKeepaliveTimeout() string {
	if m != nil {
		return m.KeepaliveTimeout
	}
	return ""
}

func (m *VPNInitRequest) GetUseLzo() bool {
	if m != nil {
		return m.UseLzo
	}
	return false
}

type VPNUpdateRequest struct {
	IpBlock              string   `protobuf:"bytes,1,opt,name=ip_block,json=ipBlock,proto3" json:"ip_block,omitempty"`
	Dns                  string   `protobuf:"bytes,2,opt,name=dns,proto3" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNUpdateRequest) Reset()         { *m = VPNUpdateRequest{} }
func (m *VPNUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*VPNUpdateRequest) ProtoMessage()    {}
func (*VPNUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{2}
}

func (m *VPNUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNUpdateRequest.Unmarshal(m, b)
}
func (m *VPNUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNUpdateRequest.Marshal(b, m, deterministic)
}
func (m *VPNUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNUpdateRequest.Merge(m, src)
}
func (m *VPNUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_VPNUpdateRequest.Size(m)
}
func (m *VPNUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNUpdateRequest proto.InternalMessageInfo

func (m *VPNUpdateRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNUpdateRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNRestartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNRestartRequest) Reset()         { *m = VPNRestartRequest{} }
func (m *VPNRestartRequest) String() string { return proto.CompactTextString(m) }
func (*VPNRestartRequest) ProtoMessage()    {}
func (*VPNRestartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{3}
}

func (m *VPNRestartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNRestartRequest.Unmarshal(m, b)
}
func (m *VPNRestartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNRestartRequest.Marshal(b, m, deterministic)
}
func (m *VPNRestartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNRestartRequest.Merge(m, src)
}
func (m *VPNRestartRequest) XXX_Size() int {
	return xxx_messageInfo_VPNRestartRequest.Size(m)
}
func (m *VPNRestartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNRestartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNRestartRequest proto.InternalMessageInfo

type VPNStatusResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SerialNumber         string   `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port                 string   `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Cert                 string   `protobuf:"bytes,5,opt,name=cert,proto3" json:"cert,omitempty"`
	CaCert               string   `protobuf:"bytes,6,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	Net                  string   `protobuf:"bytes,7,opt,name=net,proto3" json:"net,omitempty"`
	Mask                 string   `protobuf:"bytes,8,opt,name=mask,proto3" json:"mask,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Proto                string   `protobuf:"bytes,10,opt,name=proto,proto3" json:"proto,omitempty"`
	Dns                  string   `protobuf:"bytes,11,opt,name=dns,proto3" json:"dns,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,12,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CaExpiresAt          string   `protobuf:"bytes,13,opt,name=ca_expires_at,json=caExpiresAt,proto3" json:"ca_expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNStatusResponse) Reset()         { *m = VPNStatusResponse{} }
func (m *VPNStatusResponse) String() string { return proto.CompactTextString(m) }
func (*VPNStatusResponse) ProtoMessage()    {}
func (*VPNStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{4}
}

func (m *VPNStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNStatusResponse.Unmarshal(m, b)
}
func (m *VPNStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNStatusResponse.Marshal(b, m, deterministic)
}
func (m *VPNStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNStatusResponse.Merge(m, src)
}
func (m *VPNStatusResponse) XXX_Size() int {
	return xxx_messageInfo_VPNStatusResponse.Size(m)
}
func (m *VPNStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNStatusResponse proto.InternalMessageInfo

func (m *VPNStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VPNStatusResponse) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *VPNStatusResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNStatusResponse) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNStatusResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *VPNStatusResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *VPNStatusResponse) GetNet() string {
	if m != nil {
		return m.Net
	}
	return ""
}

func (m *VPNStatusResponse) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *VPNStatusResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *VPNStatusResponse) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *VPNStatusResponse) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

func (m *VPNStatusResponse) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *VPNStatusResponse) GetCaExpiresAt() string {
	if m != nil {
		return m.CaExpiresAt
	}
	return ""
}

type VPNInitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNInitResponse) Reset()         { *m = VPNInitResponse{} }
func (m *VPNInitResponse) String() string { return proto.CompactTextString(m) }
func (*VPNInitResponse) ProtoMessage()    {}
func (*VPNInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{5}
}

func (m *VPNInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNInitResponse.Unmarshal(m, b)
}
func (m *VPNInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNInitResponse.Marshal(b, m, deterministic)
}
func (m *VPNInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNInitResponse.Merge(m, src)
}
func (m *VPNInitResponse) XXX_Size() int {
	return xxx_messageInfo_VPNInitResponse.Size(m)
}
func (m *VPNInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNInitResponse proto.InternalMessageInfo

type VPNUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNUpdateResponse) Reset()         { *m = VPNUpdateResponse{} }
func (m *VPNUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*VPNUpdateResponse) ProtoMessage()    {}
func (*VPNUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{6}
}

func (m *VPNUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNUpdateResponse.Unmarshal(m, b)
}
func (m *VPNUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNUpdateResponse.Marshal(b, m, deterministic)
}
func (m *VPNUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNUpdateResponse.Merge(m, src)
}
func (m *VPNUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_VPNUpdateResponse.Size(m)
}
func (m *VPNUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNUpdateResponse proto.InternalMessageInfo

type VPNRestartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNRestartResponse) Reset()         { *m = VPNRestartResponse{} }
func (m *VPNRestartResponse) String() string { return proto.CompactTextString(m) }
func (*VPNRestartResponse) ProtoMessage()    {}
func (*VPNRestartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75c5922f473942e1, []int{7}
}

func (m *VPNRestartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNRestartResponse.Unmarshal(m, b)
}
func (m *VPNRestartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNRestartResponse.Marshal(b, m, deterministic)
}
func (m *VPNRestartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNRestartResponse.Merge(m, src)
}
func (m *VPNRestartResponse) XXX_Size() int {
	return xxx_messageInfo_VPNRestartResponse.Size(m)
}
func (m *VPNRestartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNRestartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNRestartResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("pb.VPNProto", VPNProto_name, VPNProto_value)
	proto.RegisterType((*VPNStatusRequest)(nil), "pb.VPNStatusRequest")
	proto.RegisterType((*VPNInitRequest)(nil), "pb.VPNInitRequest")
	proto.RegisterType((*VPNUpdateRequest)(nil), "pb.VPNUpdateRequest")
	proto.RegisterType((*VPNRestartRequest)(nil), "pb.VPNRestartRequest")
	proto.RegisterType((*VPNStatusResponse)(nil), "pb.VPNStatusResponse")
	proto.RegisterType((*VPNInitResponse)(nil), "pb.VPNInitResponse")
	proto.RegisterType((*VPNUpdateResponse)(nil), "pb.VPNUpdateResponse")
	proto.RegisterType((*VPNRestartResponse)(nil), "pb.VPNRestartResponse")
}

func init() { proto.RegisterFile("vpn.proto", fileDescriptor_75c5922f473942e1) }

var fileDescriptor_75c5922f473942e1 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0x5f, 0x3b, 0xa9, 0x93, 0x4c, 0xff, 0xb9, 0xd3, 0xf4, 0xad, 0x09, 0x54, 0xaa, 0xcc,
	0x25, 0xb4, 0x52, 0x23, 0xca, 0x8d, 0x0b, 0x2a, 0xa5, 0x48, 0x48, 0x55, 0x30, 0xa6, 0xcd, 0xd5,
	0xda, 0x38, 0xdb, 0x62, 0x35, 0xf1, 0x2e, 0xbb, 0xeb, 0x08, 0xf5, 0xc8, 0x95, 0x23, 0x1f, 0x8d,
	0x13, 0x9c, 0xf9, 0x20, 0xc8, 0x63, 0x3b, 0x75, 0x22, 0xb8, 0x8d, 0x7f, 0x33, 0x79, 0x34, 0x79,
	0x9e, 0x59, 0xe8, 0xcc, 0x65, 0x7a, 0x22, 0x95, 0x30, 0x02, 0x6d, 0x39, 0xee, 0x3d, 0xb9, 0x15,
	0xe2, 0x76, 0xca, 0x07, 0x4c, 0x26, 0x03, 0x96, 0xa6, 0xc2, 0x30, 0x93, 0x88, 0x54, 0x17, 0x13,
	0x3e, 0x82, 0x3b, 0x0a, 0x86, 0x1f, 0x0d, 0x33, 0x99, 0x0e, 0xf9, 0xe7, 0x8c, 0x6b, 0xe3, 0x7f,
	0xb3, 0x61, 0x6b, 0x14, 0x0c, 0xdf, 0xa5, 0x89, 0x29, 0x11, 0xf6, 0xa0, 0xfd, 0x49, 0x68, 0x93,
	0xb2, 0x19, 0xf7, 0xac, 0x43, 0xab, 0xdf, 0x09, 0x17, 0xdf, 0x88, 0xd0, 0x94, 0x42, 0x19, 0xcf,
	0x26, 0x4e, 0x35, 0x1e, 0x03, 0x90, 0x7e, 0x24, 0x15, 0xbf, 0xf1, 0x1a, 0x87, 0x56, 0x7f, 0xeb,
	0x74, 0xe3, 0x44, 0x8e, 0x4f, 0x46, 0xc1, 0x30, 0xc8, 0x1b, 0x61, 0x87, 0xfa, 0x81, 0xe2, 0x37,
	0xf8, 0x08, 0xda, 0x89, 0x8c, 0xc6, 0x53, 0x11, 0xdf, 0x79, 0x4d, 0x12, 0x69, 0x25, 0xf2, 0x75,
	0xfe, 0x89, 0x2e, 0x34, 0x26, 0xa9, 0xf6, 0xd6, 0x88, 0xe6, 0x25, 0x3e, 0x03, 0xf7, 0x8e, 0x73,
	0xc9, 0xa6, 0xc9, 0x9c, 0x47, 0x92, 0xab, 0x44, 0x4c, 0x3c, 0x87, 0xda, 0xdb, 0x0b, 0x1e, 0x10,
	0xc6, 0x63, 0xd8, 0x79, 0x18, 0x35, 0xc9, 0x8c, 0x8b, 0xcc, 0x78, 0x2d, 0x9a, 0x7d, 0xd0, 0xb8,
	0x2a, 0x38, 0xee, 0x43, 0x2b, 0xd3, 0x3c, 0x9a, 0xde, 0x0b, 0xaf, 0x7d, 0x68, 0xf5, 0xdb, 0xa1,
	0x93, 0x69, 0x7e, 0x79, 0x2f, 0xfc, 0x57, 0xe4, 0xd0, 0xb5, 0x9c, 0x30, 0xc3, 0x2b, 0x3b, 0xea,
	0x1b, 0x5b, 0x7f, 0xdd, 0xd8, 0x5e, 0x6c, 0xec, 0xef, 0xc2, 0xce, 0x28, 0x18, 0x86, 0x5c, 0x1b,
	0xa6, 0x2a, 0x43, 0xfd, 0x9f, 0x36, 0xd1, 0xca, 0x78, 0x2d, 0x45, 0xaa, 0xc9, 0xca, 0x9a, 0xc5,
	0x54, 0xe3, 0x53, 0xd8, 0xd4, 0x5c, 0x25, 0x6c, 0x1a, 0xa5, 0xd9, 0x6c, 0xcc, 0x55, 0x29, 0xbd,
	0x51, 0xc0, 0x21, 0xb1, 0xa5, 0x7c, 0x1a, 0xff, 0xc8, 0xa7, 0x59, 0xcb, 0x07, 0xa1, 0x19, 0x73,
	0x65, 0x4a, 0x63, 0xa9, 0xce, 0x1d, 0x88, 0x59, 0x44, 0xb8, 0x30, 0xd4, 0x89, 0xd9, 0x79, 0xde,
	0x70, 0xa1, 0x91, 0xf2, 0xca, 0xb9, 0xbc, 0xcc, 0x7f, 0x3e, 0x63, 0xfa, 0x8e, 0x9c, 0xea, 0x84,
	0x54, 0xe3, 0x01, 0x40, 0xac, 0x38, 0x33, 0x7c, 0x12, 0x31, 0xe3, 0x75, 0xa8, 0xd3, 0x29, 0xc9,
	0x99, 0xc1, 0x2e, 0xac, 0x51, 0xe2, 0x1e, 0x50, 0xa7, 0xf8, 0xa8, 0xdc, 0x5a, 0x7f, 0xc8, 0xf7,
	0x00, 0x80, 0x7f, 0x91, 0x89, 0xe2, 0x3a, 0x97, 0xd9, 0x28, 0x64, 0x4a, 0x72, 0x66, 0xd0, 0x87,
	0xcd, 0x98, 0x45, 0xb5, 0x89, 0x4d, 0x9a, 0x58, 0x8f, 0xd9, 0x45, 0x35, 0xe3, 0xef, 0xc0, 0xf6,
	0xe2, 0x7c, 0x0b, 0x63, 0xcb, 0x0c, 0xaa, 0x10, 0x4b, 0xd8, 0x05, 0xac, 0x07, 0x53, 0xd0, 0xa3,
	0x3e, 0xb4, 0xab, 0x23, 0x45, 0x00, 0x67, 0xf8, 0x3e, 0x08, 0x2f, 0xde, 0xba, 0xff, 0x61, 0x0b,
	0x1a, 0xd7, 0x6f, 0x02, 0xd7, 0xca, 0x8b, 0xab, 0xf3, 0xc0, 0xb5, 0x4f, 0x7f, 0xd9, 0x00, 0x79,
	0x86, 0x5c, 0xcd, 0x93, 0x98, 0xe3, 0x07, 0x70, 0x8a, 0x38, 0xb1, 0x5b, 0x5e, 0xfa, 0xd2, 0xb3,
	0xea, 0xed, 0xad, 0xd0, 0x72, 0x8b, 0xde, 0xd7, 0x1f, 0xbf, 0xbf, 0xdb, 0x5d, 0x44, 0x7a, 0xa1,
	0xf3, 0xe7, 0x83, 0xb9, 0x4c, 0x07, 0xba, 0x10, 0xba, 0x84, 0x66, 0xfe, 0x37, 0x10, 0xcb, 0x9f,
	0xd6, 0x9e, 0x64, 0x6f, 0x77, 0x89, 0x95, 0x62, 0x8f, 0x49, 0x6c, 0xcf, 0x77, 0xeb, 0x62, 0x49,
	0x9a, 0x98, 0x97, 0xd6, 0x11, 0x5e, 0x81, 0x53, 0x38, 0xb0, 0x58, 0x70, 0xe9, 0xaa, 0x17, 0x0b,
	0xae, 0xd8, 0x74, 0x40, 0x9a, 0xfb, 0xfe, 0xd2, 0x82, 0x19, 0xcd, 0xe4, 0xaa, 0xd7, 0xd0, 0x2a,
	0x2d, 0xc4, 0x4a, 0x60, 0xf9, 0xd6, 0x7b, 0xff, 0xaf, 0xe2, 0x95, 0x65, 0x77, 0xeb, 0xc2, 0xaa,
	0x18, 0x1a, 0x3b, 0x74, 0x20, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xba, 0x04, 0x57, 0xbd,
	0xce, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VPNServiceClient is the client API for VPNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VPNServiceClient interface {
	Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error)
	Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error)
	Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error)
	Restart(ctx context.Context, in *VPNRestartRequest, opts ...grpc.CallOption) (*VPNRestartResponse, error)
}

type vPNServiceClient struct {
	cc *grpc.ClientConn
}

func NewVPNServiceClient(cc *grpc.ClientConn) VPNServiceClient {
	return &vPNServiceClient{cc}
}

func (c *vPNServiceClient) Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error) {
	out := new(VPNStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.VPNService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error) {
	out := new(VPNInitResponse)
	err := c.cc.Invoke(ctx, "/pb.VPNService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error) {
	out := new(VPNUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.VPNService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Restart(ctx context.Context, in *VPNRestartRequest, opts ...grpc.CallOption) (*VPNRestartResponse, error) {
	out := new(VPNRestartResponse)
	err := c.cc.Invoke(ctx, "/pb.VPNService/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VPNServiceServer is the server API for VPNService service.
type VPNServiceServer interface {
	Status(context.Context, *VPNStatusRequest) (*VPNStatusResponse, error)
	Init(context.Context, *VPNInitRequest) (*VPNInitResponse, error)
	Update(context.Context, *VPNUpdateRequest) (*VPNUpdateResponse, error)
	Restart(context.Context, *VPNRestartRequest) (*VPNRestartResponse, error)
}

// UnimplementedVPNServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVPNServiceServer struct {
}

func (*UnimplementedVPNServiceServer) Status(ctx context.Context, req *VPNStatusRequest) (*VPNStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedVPNServiceServer) Init(ctx context.Context, req *VPNInitRequest) (*VPNInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedVPNServiceServer) Update(ctx context.Context, req *VPNUpdateRequest) (*VPNUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedVPNServiceServer) Restart(ctx context.Context, req *VPNRestartRequest) (*VPNRestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}

func RegisterVPNServiceServer(s *grpc.Server, srv VPNServiceServer) {
	s.RegisterService(&_VPNService_serviceDesc, srv)
}

func _VPNService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Status(ctx, req.(*VPNStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Init(ctx, req.(*VPNInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Update(ctx, req.(*VPNUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Restart(ctx, req.(*VPNRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VPNService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VPNService",
	HandlerType: (*VPNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _VPNService_Status_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _VPNService_Init_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VPNService_Update_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _VPNService_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vpn.proto",
}
