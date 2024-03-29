// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/internal/user/user.proto

package user

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
	github_com_elojah_trax_pkg_ulid "github.com/elojah/trax/pkg/ulid"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type U struct {
	ID           github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"ID"`
	Email        string                             `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	PasswordHash []byte                             `protobuf:"bytes,3,opt,name=PasswordHash,proto3" json:"PasswordHash,omitempty"`
	PasswordSalt []byte                             `protobuf:"bytes,4,opt,name=PasswordSalt,proto3" json:"PasswordSalt,omitempty"`
	GoogleID     string                             `protobuf:"bytes,5,opt,name=GoogleID,proto3" json:"GoogleID,omitempty"`
	CreatedAt    int64                              `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    int64                              `protobuf:"varint,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *U) Reset()      { *m = U{} }
func (*U) ProtoMessage() {}
func (*U) Descriptor() ([]byte, []int) {
	return fileDescriptor_c833e8dbf581e4f7, []int{0}
}
func (m *U) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *U) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_U.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *U) XXX_Merge(src proto.Message) {
	xxx_messageInfo_U.Merge(m, src)
}
func (m *U) XXX_Size() int {
	return m.Size()
}
func (m *U) XXX_DiscardUnknown() {
	xxx_messageInfo_U.DiscardUnknown(m)
}

var xxx_messageInfo_U proto.InternalMessageInfo

type Profile struct {
	UserID    github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,1,opt,name=UserID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"UserID"`
	FirstName string                             `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string                             `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	AvatarURL string                             `protobuf:"bytes,4,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	CreatedAt int64                              `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64                              `protobuf:"varint,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *Profile) Reset()      { *m = Profile{} }
func (*Profile) ProtoMessage() {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_c833e8dbf581e4f7, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func init() {
	proto.RegisterType((*U)(nil), "user.U")
	golang_proto.RegisterType((*U)(nil), "user.U")
	proto.RegisterType((*Profile)(nil), "user.Profile")
	golang_proto.RegisterType((*Profile)(nil), "user.Profile")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/internal/user/user.proto", fileDescriptor_c833e8dbf581e4f7)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/internal/user/user.proto", fileDescriptor_c833e8dbf581e4f7)
}

var fileDescriptor_c833e8dbf581e4f7 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbd, 0xce, 0x12, 0x41,
	0x14, 0x9d, 0x0b, 0x1f, 0xfb, 0xc9, 0xe4, 0xab, 0x36, 0x16, 0x1b, 0x62, 0x2e, 0x64, 0x2b, 0x62,
	0xc1, 0x16, 0x76, 0x76, 0xe0, 0xaa, 0x6c, 0x42, 0x0c, 0x59, 0xb3, 0x0f, 0x30, 0xb8, 0xe3, 0xb2,
	0xba, 0x30, 0x64, 0x76, 0x50, 0x4b, 0x1f, 0xc1, 0xc7, 0xf0, 0x11, 0x2c, 0x29, 0x29, 0x29, 0x89,
	0x05, 0x71, 0x67, 0x1b, 0x4b, 0x1a, 0x13, 0x4b, 0xb3, 0x3f, 0x82, 0x98, 0x48, 0x61, 0x33, 0xb9,
	0xf7, 0x9c, 0x7b, 0xee, 0xcd, 0x39, 0x19, 0xea, 0x44, 0xb1, 0x9a, 0xaf, 0x67, 0x83, 0x57, 0x62,
	0xe1, 0xf0, 0x44, 0xbc, 0x61, 0x73, 0x47, 0x49, 0xf6, 0xc1, 0x89, 0x97, 0x8a, 0xcb, 0x25, 0x4b,
	0x9c, 0x75, 0xca, 0x65, 0xf9, 0x0c, 0x56, 0x52, 0x28, 0x61, 0xde, 0x14, 0x75, 0xe7, 0x5f, 0xb2,
	0xd5, 0xdb, 0xc8, 0x89, 0x44, 0x24, 0xca, 0xd9, 0xb2, 0xaa, 0x64, 0xf6, 0x0f, 0xa0, 0x10, 0x98,
	0x8f, 0x69, 0xc3, 0x73, 0x2d, 0xe8, 0x41, 0xff, 0x6e, 0xf4, 0x70, 0x7b, 0xe8, 0x92, 0xaf, 0x87,
	0xae, 0x7d, 0x65, 0xd5, 0x3a, 0x89, 0xc3, 0x81, 0xe7, 0xfa, 0x0d, 0xcf, 0x35, 0xef, 0xd3, 0xd6,
	0xd3, 0x05, 0x8b, 0x13, 0xab, 0xd1, 0x83, 0x7e, 0xdb, 0xaf, 0x1a, 0xd3, 0xa6, 0x77, 0x53, 0x96,
	0xa6, 0xef, 0x85, 0x0c, 0xc7, 0x2c, 0x9d, 0x5b, 0xcd, 0x62, 0xb7, 0x7f, 0x81, 0xfd, 0x39, 0xf3,
	0x92, 0x25, 0xca, 0xba, 0xb9, 0x9c, 0x29, 0x30, 0xb3, 0x43, 0xef, 0x3d, 0x17, 0x22, 0x4a, 0xb8,
	0xe7, 0x5a, 0xad, 0xf2, 0xc0, 0xa9, 0x37, 0x1f, 0xd0, 0xf6, 0x13, 0xc9, 0x99, 0xe2, 0xe1, 0x50,
	0x59, 0x46, 0x0f, 0xfa, 0x4d, 0xff, 0x0c, 0x14, 0x6c, 0xb0, 0x0a, 0x6b, 0xf6, 0xb6, 0x62, 0x4f,
	0x80, 0x9d, 0x01, 0xbd, 0x9d, 0x4a, 0xf1, 0x3a, 0x4e, 0xb8, 0x39, 0xa2, 0x46, 0x90, 0x72, 0xf9,
	0x5f, 0x09, 0xd4, 0xca, 0xe2, 0xda, 0xb3, 0x58, 0xa6, 0xea, 0x05, 0x5b, 0xf0, 0x3a, 0x89, 0x33,
	0x50, 0xb8, 0x98, 0xb0, 0x9a, 0x6c, 0x56, 0x2e, 0x7e, 0xf7, 0x85, 0x72, 0xf8, 0x8e, 0x29, 0x26,
	0x03, 0x7f, 0x52, 0x46, 0xd0, 0xf6, 0xcf, 0xc0, 0xa5, 0xc7, 0xd6, 0x55, 0x8f, 0xc6, 0x5f, 0x1e,
	0x47, 0xe3, 0x6d, 0x86, 0x64, 0x97, 0x21, 0xd9, 0x67, 0x48, 0x8e, 0x19, 0xc2, 0xcf, 0x0c, 0xe1,
	0xa3, 0x46, 0xf8, 0xac, 0x11, 0xbe, 0x68, 0x84, 0x8d, 0x46, 0xd8, 0x6a, 0x84, 0x9d, 0x46, 0xf8,
	0xa6, 0x11, 0xbe, 0x6b, 0x24, 0x47, 0x8d, 0xf0, 0x29, 0x47, 0xb2, 0xc9, 0x11, 0x76, 0x39, 0x92,
	0x7d, 0x8e, 0x64, 0x66, 0x94, 0x9f, 0xe5, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x90,
	0x5c, 0xd8, 0x96, 0x02, 0x00, 0x00,
}

func (this *U) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*U)
	if !ok {
		that2, ok := that.(U)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ID.Equal(that1.ID) {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	if !bytes.Equal(this.PasswordHash, that1.PasswordHash) {
		return false
	}
	if !bytes.Equal(this.PasswordSalt, that1.PasswordSalt) {
		return false
	}
	if this.GoogleID != that1.GoogleID {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	if this.UpdatedAt != that1.UpdatedAt {
		return false
	}
	return true
}
func (this *Profile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Profile)
	if !ok {
		that2, ok := that.(Profile)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UserID.Equal(that1.UserID) {
		return false
	}
	if this.FirstName != that1.FirstName {
		return false
	}
	if this.LastName != that1.LastName {
		return false
	}
	if this.AvatarURL != that1.AvatarURL {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	if this.UpdatedAt != that1.UpdatedAt {
		return false
	}
	return true
}
func (this *U) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&user.U{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "PasswordHash: "+fmt.Sprintf("%#v", this.PasswordHash)+",\n")
	s = append(s, "PasswordSalt: "+fmt.Sprintf("%#v", this.PasswordSalt)+",\n")
	s = append(s, "GoogleID: "+fmt.Sprintf("%#v", this.GoogleID)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "UpdatedAt: "+fmt.Sprintf("%#v", this.UpdatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Profile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&user.Profile{")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "FirstName: "+fmt.Sprintf("%#v", this.FirstName)+",\n")
	s = append(s, "LastName: "+fmt.Sprintf("%#v", this.LastName)+",\n")
	s = append(s, "AvatarURL: "+fmt.Sprintf("%#v", this.AvatarURL)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "UpdatedAt: "+fmt.Sprintf("%#v", this.UpdatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUser(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *U) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *U) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *U) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x38
	}
	if m.CreatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.GoogleID) > 0 {
		i -= len(m.GoogleID)
		copy(dAtA[i:], m.GoogleID)
		i = encodeVarintUser(dAtA, i, uint64(len(m.GoogleID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PasswordSalt) > 0 {
		i -= len(m.PasswordSalt)
		copy(dAtA[i:], m.PasswordSalt)
		i = encodeVarintUser(dAtA, i, uint64(len(m.PasswordSalt)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PasswordHash) > 0 {
		i -= len(m.PasswordHash)
		copy(dAtA[i:], m.PasswordHash)
		i = encodeVarintUser(dAtA, i, uint64(len(m.PasswordHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x30
	}
	if m.CreatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AvatarURL) > 0 {
		i -= len(m.AvatarURL)
		copy(dAtA[i:], m.AvatarURL)
		i = encodeVarintUser(dAtA, i, uint64(len(m.AvatarURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.UserID.Size()
		i -= size
		if _, err := m.UserID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedU(r randyUser, easy bool) *U {
	this := &U{}
	v1 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	this.Email = string(randStringUser(r))
	v2 := r.Intn(100)
	this.PasswordHash = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.PasswordHash[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.PasswordSalt = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.PasswordSalt[i] = byte(r.Intn(256))
	}
	this.GoogleID = string(randStringUser(r))
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	this.UpdatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.UpdatedAt *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedProfile(r randyUser, easy bool) *Profile {
	this := &Profile{}
	v4 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.UserID = *v4
	this.FirstName = string(randStringUser(r))
	this.LastName = string(randStringUser(r))
	this.AvatarURL = string(randStringUser(r))
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	this.UpdatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.UpdatedAt *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyUser interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUser(r randyUser) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringUser(r randyUser) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneUser(r)
	}
	return string(tmps)
}
func randUnrecognizedUser(r randyUser, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldUser(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldUser(dAtA []byte, r randyUser, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateUser(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateUser(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateUser(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *U) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovUser(uint64(l))
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.PasswordHash)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.PasswordSalt)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.GoogleID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovUser(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovUser(uint64(m.UpdatedAt))
	}
	return n
}

func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.UserID.Size()
	n += 1 + l + sovUser(uint64(l))
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.AvatarURL)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovUser(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovUser(uint64(m.UpdatedAt))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *U) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&U{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
		`PasswordHash:` + fmt.Sprintf("%v", this.PasswordHash) + `,`,
		`PasswordSalt:` + fmt.Sprintf("%v", this.PasswordSalt) + `,`,
		`GoogleID:` + fmt.Sprintf("%v", this.GoogleID) + `,`,
		`CreatedAt:` + fmt.Sprintf("%v", this.CreatedAt) + `,`,
		`UpdatedAt:` + fmt.Sprintf("%v", this.UpdatedAt) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Profile) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Profile{`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`FirstName:` + fmt.Sprintf("%v", this.FirstName) + `,`,
		`LastName:` + fmt.Sprintf("%v", this.LastName) + `,`,
		`AvatarURL:` + fmt.Sprintf("%v", this.AvatarURL) + `,`,
		`CreatedAt:` + fmt.Sprintf("%v", this.CreatedAt) + `,`,
		`UpdatedAt:` + fmt.Sprintf("%v", this.UpdatedAt) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUser(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *U) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: U: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: U: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PasswordHash = append(m.PasswordHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PasswordHash == nil {
				m.PasswordHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordSalt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PasswordSalt = append(m.PasswordSalt[:0], dAtA[iNdEx:postIndex]...)
			if m.PasswordSalt == nil {
				m.PasswordSalt = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
