// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auction.proto

package auctionpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PickRequest struct {
	// Name of bottle to pick
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `protobuf:"bytes,2,rep,name=varietal,proto3" json:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery               string   `protobuf:"bytes,3,opt,name=winery,proto3" json:"winery,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PickRequest) Reset()         { *m = PickRequest{} }
func (m *PickRequest) String() string { return proto.CompactTextString(m) }
func (*PickRequest) ProtoMessage()    {}
func (*PickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{0}
}

func (m *PickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PickRequest.Unmarshal(m, b)
}
func (m *PickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PickRequest.Marshal(b, m, deterministic)
}
func (m *PickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PickRequest.Merge(m, src)
}
func (m *PickRequest) XXX_Size() int {
	return xxx_messageInfo_PickRequest.Size(m)
}
func (m *PickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PickRequest proto.InternalMessageInfo

func (m *PickRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PickRequest) GetVarietal() []string {
	if m != nil {
		return m.Varietal
	}
	return nil
}

func (m *PickRequest) GetWinery() string {
	if m != nil {
		return m.Winery
	}
	return ""
}

type StoredBottleCollection struct {
	Field                []*StoredBottle `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoredBottleCollection) Reset()         { *m = StoredBottleCollection{} }
func (m *StoredBottleCollection) String() string { return proto.CompactTextString(m) }
func (*StoredBottleCollection) ProtoMessage()    {}
func (*StoredBottleCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{1}
}

func (m *StoredBottleCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredBottleCollection.Unmarshal(m, b)
}
func (m *StoredBottleCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredBottleCollection.Marshal(b, m, deterministic)
}
func (m *StoredBottleCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredBottleCollection.Merge(m, src)
}
func (m *StoredBottleCollection) XXX_Size() int {
	return xxx_messageInfo_StoredBottleCollection.Size(m)
}
func (m *StoredBottleCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredBottleCollection.DiscardUnknown(m)
}

var xxx_messageInfo_StoredBottleCollection proto.InternalMessageInfo

func (m *StoredBottleCollection) GetField() []*StoredBottle {
	if m != nil {
		return m.Field
	}
	return nil
}

// A StoredBottle describes a bottle retrieved by the storage service.
type StoredBottle struct {
	// ID is the unique id of the bottle.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name of bottle
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Winery that produces wine
	Winery *Winery `protobuf:"bytes,2,opt,name=winery,proto3" json:"winery,omitempty"`
	// Vintage of bottle
	Vintage uint32 `protobuf:"varint,3,opt,name=vintage,proto3" json:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component `protobuf:"bytes,4,rep,name=composition,proto3" json:"composition,omitempty"`
	// Description of bottle
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating               uint32   `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoredBottle) Reset()         { *m = StoredBottle{} }
func (m *StoredBottle) String() string { return proto.CompactTextString(m) }
func (*StoredBottle) ProtoMessage()    {}
func (*StoredBottle) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{2}
}

func (m *StoredBottle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredBottle.Unmarshal(m, b)
}
func (m *StoredBottle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredBottle.Marshal(b, m, deterministic)
}
func (m *StoredBottle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredBottle.Merge(m, src)
}
func (m *StoredBottle) XXX_Size() int {
	return xxx_messageInfo_StoredBottle.Size(m)
}
func (m *StoredBottle) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredBottle.DiscardUnknown(m)
}

var xxx_messageInfo_StoredBottle proto.InternalMessageInfo

func (m *StoredBottle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoredBottle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoredBottle) GetWinery() *Winery {
	if m != nil {
		return m.Winery
	}
	return nil
}

func (m *StoredBottle) GetVintage() uint32 {
	if m != nil {
		return m.Vintage
	}
	return 0
}

func (m *StoredBottle) GetComposition() []*Component {
	if m != nil {
		return m.Composition
	}
	return nil
}

func (m *StoredBottle) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StoredBottle) GetRating() uint32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type Winery struct {
	// Name of winery
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Region of winery
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Country of winery
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	// Winery website URL
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Winery) Reset()         { *m = Winery{} }
func (m *Winery) String() string { return proto.CompactTextString(m) }
func (*Winery) ProtoMessage()    {}
func (*Winery) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{3}
}

func (m *Winery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Winery.Unmarshal(m, b)
}
func (m *Winery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Winery.Marshal(b, m, deterministic)
}
func (m *Winery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Winery.Merge(m, src)
}
func (m *Winery) XXX_Size() int {
	return xxx_messageInfo_Winery.Size(m)
}
func (m *Winery) XXX_DiscardUnknown() {
	xxx_messageInfo_Winery.DiscardUnknown(m)
}

var xxx_messageInfo_Winery proto.InternalMessageInfo

func (m *Winery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Winery) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Winery) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Winery) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Component struct {
	// Grape varietal
	Varietal string `protobuf:"bytes,1,opt,name=varietal,proto3" json:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage           uint32   `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{4}
}

func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}
func (m *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(m, src)
}
func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetVarietal() string {
	if m != nil {
		return m.Varietal
	}
	return ""
}

func (m *Component) GetPercentage() uint32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PickRequest)(nil), "auction.PickRequest")
	proto.RegisterType((*StoredBottleCollection)(nil), "auction.StoredBottleCollection")
	proto.RegisterType((*StoredBottle)(nil), "auction.StoredBottle")
	proto.RegisterType((*Winery)(nil), "auction.Winery")
	proto.RegisterType((*Component)(nil), "auction.Component")
	proto.RegisterType((*GetRequest)(nil), "auction.GetRequest")
}

func init() { proto.RegisterFile("auction.proto", fileDescriptor_622f477c3a3f2896) }

var fileDescriptor_622f477c3a3f2896 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xf1, 0x9f, 0x38, 0xf1, 0x38, 0x69, 0x8b, 0xda, 0x06, 0x91, 0x43, 0x6b, 0x7c, 0xa9,
	0xa1, 0x90, 0x43, 0xda, 0x5b, 0xe9, 0xa1, 0x09, 0x25, 0xd7, 0xe2, 0x52, 0x0a, 0x7b, 0x5a, 0xc7,
	0x9e, 0x0d, 0x62, 0x1d, 0xc9, 0xab, 0xc8, 0x59, 0xf6, 0x01, 0xf6, 0x51, 0xf7, 0x3d, 0x16, 0xc9,
	0xb6, 0xe2, 0x85, 0xc0, 0xde, 0xf4, 0x8d, 0x67, 0xf4, 0xfd, 0xfc, 0x69, 0x60, 0x96, 0x37, 0x85,
	0x62, 0x82, 0x2f, 0x6b, 0x29, 0x94, 0x20, 0xe3, 0x4e, 0x26, 0xff, 0x20, 0xfa, 0xc3, 0x8a, 0xdb,
	0x0c, 0xef, 0x1a, 0x3c, 0x2a, 0x42, 0xc0, 0xe7, 0xf9, 0x01, 0xa9, 0x13, 0x3b, 0x69, 0x98, 0x99,
	0x33, 0x59, 0xc0, 0xe4, 0x94, 0x4b, 0x86, 0x2a, 0xaf, 0xa8, 0x1b, 0x7b, 0x69, 0x98, 0x59, 0x4d,
	0xe6, 0x10, 0xdc, 0x33, 0x8e, 0xf2, 0x81, 0x7a, 0x66, 0xa2, 0x53, 0xc9, 0x6f, 0x98, 0xff, 0x55,
	0x42, 0x62, 0xb9, 0x16, 0x4a, 0x55, 0xb8, 0x11, 0x55, 0x85, 0xc6, 0x90, 0x7c, 0x85, 0xd1, 0x0d,
	0xc3, 0xaa, 0xa4, 0x4e, 0xec, 0xa5, 0xd1, 0xea, 0xe3, 0xb2, 0x07, 0x1b, 0xf6, 0x67, 0x6d, 0x4f,
	0xf2, 0xe4, 0xc0, 0x74, 0x58, 0x27, 0x6f, 0xc0, 0x65, 0x25, 0x9d, 0x18, 0x2f, 0x97, 0x95, 0x17,
	0x79, 0xbf, 0x58, 0x26, 0x37, 0x76, 0xd2, 0x68, 0xf5, 0xd6, 0x5a, 0xfc, 0x37, 0xe5, 0x1e, 0x92,
	0x50, 0x18, 0x9f, 0x18, 0x57, 0xf9, 0x1e, 0x0d, 0xfd, 0x2c, 0xeb, 0x25, 0xf9, 0x0e, 0x51, 0x21,
	0x0e, 0xb5, 0x38, 0x32, 0x3d, 0x47, 0x7d, 0x83, 0x4a, 0xec, 0x3d, 0x1b, 0xfd, 0x8d, 0x23, 0x57,
	0xd9, 0xb0, 0x8d, 0xc4, 0x10, 0x95, 0x78, 0x2c, 0x24, 0xab, 0xcd, 0xd4, 0xc8, 0x30, 0x0d, 0x4b,
	0x3a, 0x2e, 0x99, 0x2b, 0xc6, 0xf7, 0x34, 0x30, 0x86, 0x9d, 0x4a, 0xae, 0x21, 0x68, 0xd9, 0x2e,
	0xfe, 0x90, 0x9e, 0xc2, 0xbd, 0xbe, 0xd2, 0x6d, 0x43, 0x6e, 0x95, 0xe6, 0x2f, 0x44, 0xc3, 0x95,
	0x4d, 0xbf, 0x97, 0xe4, 0x1d, 0x78, 0x8d, 0xac, 0xa8, 0x6f, 0xaa, 0xfa, 0x98, 0x6c, 0x21, 0xb4,
	0xd4, 0x2f, 0x5e, 0xb4, 0x35, 0x3a, 0xbf, 0xe8, 0x27, 0x80, 0x1a, 0x65, 0x81, 0x6d, 0x2e, 0xae,
	0xc1, 0x1c, 0x54, 0x92, 0x29, 0xc0, 0x16, 0x55, 0xb7, 0x2f, 0xab, 0x47, 0x07, 0xc6, 0xbf, 0xda,
	0x54, 0xc8, 0x4f, 0xf0, 0xf5, 0x2a, 0x91, 0x0f, 0x36, 0xa7, 0xc1, 0x66, 0x2d, 0x3e, 0x5f, 0x7c,
	0xe8, 0xc1, 0x62, 0xfc, 0x00, 0x6f, 0x8b, 0x8a, 0xbc, 0xb7, 0x7d, 0x67, 0x9b, 0x57, 0x87, 0xd7,
	0xd1, 0x55, 0xd8, 0x75, 0xd4, 0xbb, 0x5d, 0x60, 0x76, 0xfc, 0xdb, 0x73, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x74, 0xd4, 0x96, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuctionClient interface {
	// Pick implements pick.
	Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error)
	// get
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error)
}

type auctionClient struct {
	cc *grpc.ClientConn
}

func NewAuctionClient(cc *grpc.ClientConn) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error) {
	out := new(StoredBottleCollection)
	err := c.cc.Invoke(ctx, "/auction.Auction/Pick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error) {
	out := new(StoredBottleCollection)
	err := c.cc.Invoke(ctx, "/auction.Auction/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
type AuctionServer interface {
	// Pick implements pick.
	Pick(context.Context, *PickRequest) (*StoredBottleCollection, error)
	// get
	Get(context.Context, *GetRequest) (*StoredBottleCollection, error)
}

// UnimplementedAuctionServer can be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (*UnimplementedAuctionServer) Pick(ctx context.Context, req *PickRequest) (*StoredBottleCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pick not implemented")
}
func (*UnimplementedAuctionServer) Get(ctx context.Context, req *GetRequest) (*StoredBottleCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterAuctionServer(s *grpc.Server, srv AuctionServer) {
	s.RegisterService(&_Auction_serviceDesc, srv)
}

func _Auction_Pick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Pick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Auction/Pick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Pick(ctx, req.(*PickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Auction/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auction.Auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pick",
			Handler:    _Auction_Pick_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Auction_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction.proto",
}
