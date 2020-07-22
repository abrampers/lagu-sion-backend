// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lagu_sion.proto

package lagusion

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

type SongBook int32

const (
	SongBook_ALL                     SongBook = 0
	SongBook_LAGU_SION               SongBook = 1
	SongBook_LAGU_SION_EDISI_LENGKAP SongBook = 2
)

var SongBook_name = map[int32]string{
	0: "ALL",
	1: "LAGU_SION",
	2: "LAGU_SION_EDISI_LENGKAP",
}

var SongBook_value = map[string]int32{
	"ALL":                     0,
	"LAGU_SION":               1,
	"LAGU_SION_EDISI_LENGKAP": 2,
}

func (x SongBook) String() string {
	return proto.EnumName(SongBook_name, int32(x))
}

func (SongBook) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{0}
}

type SortOptions int32

const (
	SortOptions_NUMBER   SortOptions = 0
	SortOptions_ALPHABET SortOptions = 1
)

var SortOptions_name = map[int32]string{
	0: "NUMBER",
	1: "ALPHABET",
}

var SortOptions_value = map[string]int32{
	"NUMBER":   0,
	"ALPHABET": 1,
}

func (x SortOptions) String() string {
	return proto.EnumName(SortOptions_name, int32(x))
}

func (SortOptions) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{1}
}

type Verse struct {
	Contents             []string `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Verse) Reset()         { *m = Verse{} }
func (m *Verse) String() string { return proto.CompactTextString(m) }
func (*Verse) ProtoMessage()    {}
func (*Verse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{0}
}

func (m *Verse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verse.Unmarshal(m, b)
}
func (m *Verse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verse.Marshal(b, m, deterministic)
}
func (m *Verse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verse.Merge(m, src)
}
func (m *Verse) XXX_Size() int {
	return xxx_messageInfo_Verse.Size(m)
}
func (m *Verse) XXX_DiscardUnknown() {
	xxx_messageInfo_Verse.DiscardUnknown(m)
}

var xxx_messageInfo_Verse proto.InternalMessageInfo

func (m *Verse) GetContents() []string {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Song struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Verses               []*Verse `protobuf:"bytes,4,rep,name=verses,proto3" json:"verses,omitempty"`
	Reff                 *Verse   `protobuf:"bytes,5,opt,name=reff,proto3" json:"reff,omitempty"`
	SongBook             SongBook `protobuf:"varint,6,opt,name=song_book,json=songBook,proto3,enum=lagusion.SongBook" json:"song_book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{1}
}

func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Song) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Song) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Song) GetVerses() []*Verse {
	if m != nil {
		return m.Verses
	}
	return nil
}

func (m *Song) GetReff() *Verse {
	if m != nil {
		return m.Reff
	}
	return nil
}

func (m *Song) GetSongBook() SongBook {
	if m != nil {
		return m.SongBook
	}
	return SongBook_ALL
}

type ListSongRequest struct {
	SongBook             SongBook    `protobuf:"varint,1,opt,name=song_book,json=songBook,proto3,enum=lagusion.SongBook" json:"song_book,omitempty"`
	SortOptions          SortOptions `protobuf:"varint,2,opt,name=sort_options,json=sortOptions,proto3,enum=lagusion.SortOptions" json:"sort_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListSongRequest) Reset()         { *m = ListSongRequest{} }
func (m *ListSongRequest) String() string { return proto.CompactTextString(m) }
func (*ListSongRequest) ProtoMessage()    {}
func (*ListSongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{2}
}

func (m *ListSongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSongRequest.Unmarshal(m, b)
}
func (m *ListSongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSongRequest.Marshal(b, m, deterministic)
}
func (m *ListSongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSongRequest.Merge(m, src)
}
func (m *ListSongRequest) XXX_Size() int {
	return xxx_messageInfo_ListSongRequest.Size(m)
}
func (m *ListSongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSongRequest proto.InternalMessageInfo

func (m *ListSongRequest) GetSongBook() SongBook {
	if m != nil {
		return m.SongBook
	}
	return SongBook_ALL
}

func (m *ListSongRequest) GetSortOptions() SortOptions {
	if m != nil {
		return m.SortOptions
	}
	return SortOptions_NUMBER
}

type ListSongResponse struct {
	Songs                []*Song  `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSongResponse) Reset()         { *m = ListSongResponse{} }
func (m *ListSongResponse) String() string { return proto.CompactTextString(m) }
func (*ListSongResponse) ProtoMessage()    {}
func (*ListSongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac555336d8700ef, []int{3}
}

func (m *ListSongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSongResponse.Unmarshal(m, b)
}
func (m *ListSongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSongResponse.Marshal(b, m, deterministic)
}
func (m *ListSongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSongResponse.Merge(m, src)
}
func (m *ListSongResponse) XXX_Size() int {
	return xxx_messageInfo_ListSongResponse.Size(m)
}
func (m *ListSongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSongResponse proto.InternalMessageInfo

func (m *ListSongResponse) GetSongs() []*Song {
	if m != nil {
		return m.Songs
	}
	return nil
}

func init() {
	proto.RegisterEnum("lagusion.SongBook", SongBook_name, SongBook_value)
	proto.RegisterEnum("lagusion.SortOptions", SortOptions_name, SortOptions_value)
	proto.RegisterType((*Verse)(nil), "lagusion.Verse")
	proto.RegisterType((*Song)(nil), "lagusion.Song")
	proto.RegisterType((*ListSongRequest)(nil), "lagusion.ListSongRequest")
	proto.RegisterType((*ListSongResponse)(nil), "lagusion.ListSongResponse")
}

func init() { proto.RegisterFile("lagu_sion.proto", fileDescriptor_4ac555336d8700ef) }

var fileDescriptor_4ac555336d8700ef = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6f, 0x9b, 0x40,
	0x10, 0xc5, 0xbd, 0xb6, 0xa1, 0x30, 0x24, 0x80, 0x46, 0xfd, 0x43, 0xdd, 0x0b, 0x22, 0x95, 0x82,
	0x72, 0x70, 0x25, 0x7a, 0xc9, 0xad, 0xc2, 0x0a, 0x4a, 0xad, 0x52, 0x27, 0x5a, 0x9a, 0xf6, 0x88,
	0xe2, 0x64, 0x83, 0x56, 0x49, 0x77, 0x5d, 0x76, 0xc9, 0xa9, 0xdf, 0xae, 0x5f, 0xac, 0x02, 0xec,
	0x90, 0x54, 0x91, 0x7a, 0xdb, 0xb7, 0x6f, 0x76, 0xde, 0x6f, 0x46, 0x0b, 0xde, 0xdd, 0x65, 0xd5,
	0x94, 0x8a, 0x4b, 0x31, 0xdf, 0xd4, 0x52, 0x4b, 0xb4, 0xda, 0x8b, 0x56, 0x47, 0x07, 0x60, 0x7c,
	0x67, 0xb5, 0x62, 0x38, 0x03, 0xeb, 0x4a, 0x0a, 0xcd, 0x84, 0x56, 0x01, 0x09, 0x27, 0xb1, 0x4d,
	0x1f, 0x74, 0xf4, 0x87, 0xc0, 0xb4, 0x90, 0xa2, 0x42, 0x17, 0xc6, 0xfc, 0x3a, 0x20, 0x21, 0x89,
	0xf7, 0xe9, 0x98, 0x5f, 0xe3, 0x6b, 0x30, 0x45, 0xf3, 0x73, 0xcd, 0xea, 0x60, 0x1c, 0x92, 0xd8,
	0xa0, 0x5b, 0x85, 0x2f, 0xc1, 0xd0, 0x5c, 0xdf, 0xb1, 0x60, 0x12, 0x92, 0xd8, 0xa6, 0xbd, 0xc0,
	0x43, 0x30, 0xef, 0xdb, 0x2c, 0x15, 0x4c, 0xc3, 0x49, 0xec, 0x24, 0xde, 0x7c, 0x87, 0x31, 0xef,
	0x18, 0xe8, 0xd6, 0xc6, 0x03, 0x98, 0xd6, 0xec, 0xe6, 0x26, 0x30, 0x42, 0xf2, 0x5c, 0x59, 0x67,
	0xe2, 0x07, 0xb0, 0x95, 0x14, 0x55, 0xb9, 0x96, 0xf2, 0x36, 0x30, 0x43, 0x12, 0xbb, 0x09, 0x0e,
	0x95, 0x2d, 0xee, 0x42, 0xca, 0x5b, 0x6a, 0xa9, 0xed, 0x29, 0xfa, 0x0d, 0x5e, 0xce, 0x95, 0x6e,
	0x1d, 0xca, 0x7e, 0x35, 0x4c, 0xe9, 0xa7, 0x3d, 0xc8, 0xff, 0x7b, 0xe0, 0x31, 0xec, 0x29, 0x59,
	0xeb, 0x52, 0x6e, 0x34, 0x97, 0x42, 0x75, 0x63, 0xbb, 0xc9, 0xab, 0xc7, 0x6f, 0x6a, 0x7d, 0xd6,
	0x9b, 0xd4, 0x51, 0x83, 0x88, 0x8e, 0xc1, 0x1f, 0xd2, 0xd5, 0x46, 0x0a, 0xc5, 0xf0, 0x3d, 0x18,
	0x6d, 0xe7, 0x7e, 0xe1, 0x4e, 0xe2, 0x3e, 0x8d, 0xa6, 0xbd, 0x79, 0xf4, 0x09, 0xac, 0x1d, 0x09,
	0xbe, 0x80, 0x49, 0x9a, 0xe7, 0xfe, 0x08, 0xf7, 0xc1, 0xce, 0xd3, 0xd3, 0x8b, 0xb2, 0x58, 0x9e,
	0xad, 0x7c, 0x82, 0xef, 0xe0, 0xcd, 0x83, 0x2c, 0xb3, 0x93, 0x65, 0xb1, 0x2c, 0xf3, 0x6c, 0x75,
	0xfa, 0x25, 0x3d, 0xf7, 0xc7, 0x47, 0x87, 0xe0, 0x3c, 0xc2, 0x42, 0x00, 0x73, 0x75, 0xf1, 0x75,
	0x91, 0x51, 0x7f, 0x84, 0x7b, 0x60, 0xa5, 0xf9, 0xf9, 0xe7, 0x74, 0x91, 0x7d, 0xf3, 0x49, 0xf2,
	0x03, 0xbc, 0xfc, 0xb2, 0x6a, 0x0a, 0x2e, 0x45, 0xc1, 0xea, 0x7b, 0x7e, 0xc5, 0xf0, 0x04, 0xec,
	0x1d, 0xb6, 0xc2, 0xb7, 0x03, 0xe0, 0x3f, 0x9b, 0x9c, 0xcd, 0x9e, 0xb3, 0xfa, 0x31, 0xa3, 0xd1,
	0xda, 0xec, 0xbe, 0xdd, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0xa7, 0x0b, 0xfa, 0x89,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaguSionServiceClient is the client API for LaguSionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaguSionServiceClient interface {
	ListSongs(ctx context.Context, in *ListSongRequest, opts ...grpc.CallOption) (*ListSongResponse, error)
}

type laguSionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaguSionServiceClient(cc grpc.ClientConnInterface) LaguSionServiceClient {
	return &laguSionServiceClient{cc}
}

func (c *laguSionServiceClient) ListSongs(ctx context.Context, in *ListSongRequest, opts ...grpc.CallOption) (*ListSongResponse, error) {
	out := new(ListSongResponse)
	err := c.cc.Invoke(ctx, "/lagusion.LaguSionService/ListSongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaguSionServiceServer is the server API for LaguSionService service.
type LaguSionServiceServer interface {
	ListSongs(context.Context, *ListSongRequest) (*ListSongResponse, error)
}

// UnimplementedLaguSionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaguSionServiceServer struct {
}

func (*UnimplementedLaguSionServiceServer) ListSongs(ctx context.Context, req *ListSongRequest) (*ListSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSongs not implemented")
}

func RegisterLaguSionServiceServer(s *grpc.Server, srv LaguSionServiceServer) {
	s.RegisterService(&_LaguSionService_serviceDesc, srv)
}

func _LaguSionService_ListSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaguSionServiceServer).ListSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lagusion.LaguSionService/ListSongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaguSionServiceServer).ListSongs(ctx, req.(*ListSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaguSionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lagusion.LaguSionService",
	HandlerType: (*LaguSionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSongs",
			Handler:    _LaguSionService_ListSongs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lagu_sion.proto",
}
