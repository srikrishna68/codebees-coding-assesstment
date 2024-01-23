// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

package pb

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

type PostResponse_Status int32

const (
	PostResponse_UNKNOWN PostResponse_Status = 0
	PostResponse_SUCCESS PostResponse_Status = 1
	PostResponse_ERROR   PostResponse_Status = 2
)

var PostResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "ERROR",
}

var PostResponse_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"SUCCESS": 1,
	"ERROR":   2,
}

func (x PostResponse_Status) String() string {
	return proto.EnumName(PostResponse_Status_name, int32(x))
}

func (PostResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{2, 0}
}

type DeleteResponse_Status int32

const (
	DeleteResponse_UNKNOWN DeleteResponse_Status = 0
	DeleteResponse_SUCCESS DeleteResponse_Status = 1
	DeleteResponse_ERROR   DeleteResponse_Status = 2
)

var DeleteResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "ERROR",
}

var DeleteResponse_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"SUCCESS": 1,
	"ERROR":   2,
}

func (x DeleteResponse_Status) String() string {
	return proto.EnumName(DeleteResponse_Status_name, int32(x))
}

func (DeleteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{3, 0}
}

type Post struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublicationDate      string   `protobuf:"bytes,4,opt,name=publication_date,json=publicationDate,proto3" json:"publication_date,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Post) GetPublicationDate() string {
	if m != nil {
		return m.PublicationDate
	}
	return ""
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PostID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostID) Reset()         { *m = PostID{} }
func (m *PostID) String() string { return proto.CompactTextString(m) }
func (*PostID) ProtoMessage()    {}
func (*PostID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{1}
}

func (m *PostID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostID.Unmarshal(m, b)
}
func (m *PostID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostID.Marshal(b, m, deterministic)
}
func (m *PostID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostID.Merge(m, src)
}
func (m *PostID) XXX_Size() int {
	return xxx_messageInfo_PostID.Size(m)
}
func (m *PostID) XXX_DiscardUnknown() {
	xxx_messageInfo_PostID.DiscardUnknown(m)
}

var xxx_messageInfo_PostID proto.InternalMessageInfo

func (m *PostID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PostResponse struct {
	Status               PostResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.PostResponse_Status" json:"status,omitempty"`
	PostId               string              `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Post                 *Post               `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Error                string              `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{2}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetStatus() PostResponse_Status {
	if m != nil {
		return m.Status
	}
	return PostResponse_UNKNOWN
}

func (m *PostResponse) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *PostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *PostResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeleteResponse struct {
	Status               DeleteResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.DeleteResponse_Status" json:"status,omitempty"`
	Error                string                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5203cdc85000bc, []int{3}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() DeleteResponse_Status {
	if m != nil {
		return m.Status
	}
	return DeleteResponse_UNKNOWN
}

func (m *DeleteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.PostResponse_Status", PostResponse_Status_name, PostResponse_Status_value)
	proto.RegisterEnum("proto.DeleteResponse_Status", DeleteResponse_Status_name, DeleteResponse_Status_value)
	proto.RegisterType((*Post)(nil), "proto.Post")
	proto.RegisterType((*PostID)(nil), "proto.PostID")
	proto.RegisterType((*PostResponse)(nil), "proto.PostResponse")
	proto.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
}

func init() {
	proto.RegisterFile("proto/blog.proto", fileDescriptor_fc5203cdc85000bc)
}

var fileDescriptor_fc5203cdc85000bc = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0x9a, 0xa4, 0x76, 0xa2, 0x35, 0xac, 0xff, 0x42, 0x11, 0x2c, 0x39, 0xd5, 0x83,
	0x11, 0xa2, 0x37, 0x6f, 0x6d, 0x7a, 0x28, 0x42, 0x2b, 0x1b, 0x8a, 0xe0, 0xa5, 0x6c, 0x9a, 0xa5,
	0x06, 0x42, 0x36, 0x24, 0x53, 0x5f, 0x42, 0x7c, 0x28, 0x1f, 0xc1, 0x37, 0x92, 0x6c, 0x62, 0x89,
	0xa5, 0x82, 0x78, 0xca, 0x7c, 0x3b, 0x3f, 0x66, 0xe6, 0xfb, 0x08, 0x58, 0x59, 0x2e, 0x50, 0xdc,
	0x84, 0x89, 0x58, 0xb9, 0xb2, 0x24, 0xba, 0xfc, 0x38, 0xef, 0x0a, 0x68, 0x8f, 0xa2, 0x40, 0x72,
	0x02, 0x3a, 0xc6, 0x98, 0x70, 0x5b, 0xe9, 0x2b, 0x83, 0x0e, 0xad, 0x04, 0xb1, 0xa1, 0xbd, 0x14,
	0x29, 0xf2, 0x14, 0x6d, 0x55, 0xbe, 0x7f, 0x4b, 0x72, 0x06, 0x06, 0x5b, 0xe3, 0x8b, 0xc8, 0xed,
	0x96, 0x6c, 0xd4, 0x8a, 0x5c, 0x81, 0x95, 0xad, 0xc3, 0x24, 0x5e, 0x32, 0x8c, 0x45, 0xba, 0x88,
	0x18, 0x72, 0x5b, 0x93, 0xc4, 0x51, 0xe3, 0xdd, 0x67, 0xc8, 0x09, 0x01, 0x0d, 0xd9, 0xaa, 0xb0,
	0xf5, 0x7e, 0x6b, 0xd0, 0xa1, 0xb2, 0x76, 0x6c, 0x30, 0xca, 0x73, 0x26, 0x3e, 0xe9, 0x82, 0x1a,
	0x47, 0xf5, 0x35, 0x6a, 0x1c, 0x39, 0x1f, 0x0a, 0x1c, 0x94, 0x2d, 0xca, 0x8b, 0x4c, 0xa4, 0x05,
	0x27, 0x1e, 0x18, 0x05, 0x32, 0x5c, 0x17, 0x12, 0xea, 0x7a, 0xbd, 0xca, 0x99, 0xdb, 0x84, 0xdc,
	0x40, 0x12, 0xb4, 0x26, 0xc9, 0x39, 0xb4, 0x33, 0x51, 0xe0, 0x22, 0x8e, 0x6a, 0x3f, 0x46, 0x29,
	0x27, 0x11, 0xb9, 0x04, 0xad, 0xac, 0xa4, 0x19, 0xd3, 0x33, 0x9b, 0xa3, 0x64, 0xa3, 0xcc, 0x87,
	0xe7, 0xb9, 0xc8, 0x6b, 0x33, 0x95, 0x70, 0xae, 0xc1, 0xa8, 0x36, 0x10, 0x13, 0xda, 0xf3, 0xe9,
	0xc3, 0x74, 0xf6, 0x34, 0xb5, 0xf6, 0x4a, 0x11, 0xcc, 0x47, 0xa3, 0x71, 0x10, 0x58, 0x0a, 0xe9,
	0x80, 0x3e, 0xa6, 0x74, 0x46, 0x2d, 0xd5, 0x79, 0x53, 0xa0, 0xeb, 0xf3, 0x84, 0x23, 0xdf, 0xb8,
	0xb8, 0xdb, 0x72, 0x71, 0x51, 0xaf, 0xfe, 0x89, 0x6d, 0xfb, 0xd8, 0x5c, 0xa3, 0xfe, 0xff, 0x1a,
	0xef, 0x53, 0x01, 0x73, 0x98, 0x88, 0x55, 0xc0, 0xf3, 0xd7, 0x78, 0xc9, 0x89, 0x0b, 0x30, 0xca,
	0x39, 0x43, 0x2e, 0x7f, 0x88, 0x66, 0x06, 0xbd, 0xe3, 0x1d, 0xd9, 0x12, 0x17, 0xf6, 0x29, 0x67,
	0x91, 0xa4, 0x0f, 0x1b, 0xc0, 0xc4, 0xff, 0x8d, 0x87, 0x79, 0x16, 0xfd, 0x7d, 0xbe, 0x07, 0x50,
	0xa5, 0xb0, 0x6b, 0xc3, 0xe9, 0xce, 0x9c, 0x86, 0xc6, 0xb3, 0xe6, 0xde, 0x67, 0x61, 0x68, 0xc8,
	0xee, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xaa, 0x4c, 0xd5, 0xf9, 0x02, 0x00, 0x00,
}