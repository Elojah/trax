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
	FirstName    string                             `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName     string                             `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
	AvatarURL    string                             `protobuf:"bytes,8,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	CreatedAt    int64                              `protobuf:"varint,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    int64                              `protobuf:"varint,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
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

func init() {
	proto.RegisterType((*U)(nil), "user.U")
	golang_proto.RegisterType((*U)(nil), "user.U")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/internal/user/user.proto", fileDescriptor_c833e8dbf581e4f7)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/internal/user/user.proto", fileDescriptor_c833e8dbf581e4f7)
}

var fileDescriptor_c833e8dbf581e4f7 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x6f, 0xda, 0x40,
	0x18, 0xc6, 0xef, 0x35, 0x7f, 0x0a, 0x27, 0x26, 0xab, 0xc3, 0x09, 0x55, 0x2f, 0x88, 0x09, 0x75,
	0xc0, 0x43, 0xb7, 0x6e, 0x50, 0xb7, 0xc5, 0x12, 0xaa, 0x2a, 0x57, 0x7c, 0x80, 0xa3, 0xb6, 0x8c,
	0x5b, 0xc3, 0xa1, 0xf3, 0xd1, 0x66, 0xcc, 0x47, 0xc8, 0xc7, 0xc8, 0x47, 0xc8, 0xc8, 0xc8, 0x12,
	0x89, 0x11, 0x65, 0x40, 0xf1, 0x79, 0xc9, 0xc8, 0x98, 0x31, 0xb2, 0x4d, 0x20, 0x0c, 0xc9, 0x62,
	0xbd, 0xef, 0xf3, 0x7b, 0x7e, 0xaf, 0x87, 0xa3, 0x56, 0x10, 0xaa, 0xe9, 0x72, 0xd2, 0xfb, 0x2d,
	0x66, 0x96, 0x1f, 0x89, 0x3f, 0x7c, 0x6a, 0x29, 0xc9, 0x2f, 0xac, 0x70, 0xae, 0x7c, 0x39, 0xe7,
	0x91, 0xb5, 0x8c, 0x7d, 0x99, 0x7f, 0x7a, 0x0b, 0x29, 0x94, 0x30, 0xcb, 0xd9, 0xdc, 0x7c, 0x4d,
	0x5b, 0xfc, 0x0d, 0xac, 0x40, 0x04, 0x22, 0xef, 0xe6, 0x53, 0xa1, 0x75, 0x6e, 0x0d, 0x0a, 0x63,
	0xf3, 0x33, 0x35, 0x1c, 0x9b, 0x41, 0x1b, 0xba, 0x8d, 0xc1, 0xc7, 0xf5, 0xae, 0x45, 0xee, 0x76,
	0xad, 0xce, 0x1b, 0xa7, 0x96, 0x51, 0xe8, 0xf5, 0x1c, 0xdb, 0x35, 0x1c, 0xdb, 0x7c, 0x4f, 0x2b,
	0x5f, 0x67, 0x3c, 0x8c, 0x98, 0xd1, 0x86, 0x6e, 0xdd, 0x2d, 0x16, 0xb3, 0x43, 0x1b, 0x3f, 0x79,
	0x1c, 0xff, 0x17, 0xd2, 0x1b, 0xf2, 0x78, 0xca, 0x4a, 0xd9, 0x6d, 0xf7, 0x2c, 0x7b, 0xd9, 0xf9,
	0xc5, 0x23, 0xc5, 0xca, 0xe7, 0x9d, 0x2c, 0x33, 0x9b, 0xb4, 0xf6, 0x5d, 0x88, 0x20, 0xf2, 0x1d,
	0x9b, 0x55, 0xf2, 0x1f, 0x1c, 0x77, 0xf3, 0x03, 0xad, 0x7f, 0x0b, 0x65, 0xac, 0x7e, 0xf0, 0x99,
	0xcf, 0xaa, 0x39, 0x3c, 0x05, 0x99, 0x39, 0xe2, 0x07, 0xf8, 0xae, 0x30, 0x9f, 0xf7, 0xcc, 0xec,
	0xff, 0xe3, 0x8a, 0xcb, 0xb1, 0x3b, 0x62, 0xb5, 0xc2, 0x3c, 0x06, 0x19, 0xfd, 0x22, 0x7d, 0xae,
	0x7c, 0xaf, 0xaf, 0x58, 0xbd, 0x0d, 0xdd, 0x92, 0x7b, 0x0a, 0x32, 0x3a, 0x5e, 0x78, 0x07, 0x4a,
	0x0b, 0x7a, 0x0c, 0x06, 0xc3, 0x75, 0x82, 0x64, 0x93, 0x20, 0xd9, 0x26, 0x48, 0xf6, 0x09, 0xc2,
	0x63, 0x82, 0x70, 0xa9, 0x11, 0xae, 0x35, 0xc2, 0x8d, 0x46, 0x58, 0x69, 0x84, 0xb5, 0x46, 0xd8,
	0x68, 0x84, 0x7b, 0x8d, 0xf0, 0xa0, 0x91, 0xec, 0x35, 0xc2, 0x55, 0x8a, 0x64, 0x95, 0x22, 0x6c,
	0x52, 0x24, 0xdb, 0x14, 0xc9, 0xa4, 0x9a, 0x3f, 0xd0, 0xa7, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4d, 0xcd, 0xfd, 0xfc, 0x0a, 0x02, 0x00, 0x00,
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
	s := make([]string, 0, 14)
	s = append(s, "&user.U{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "PasswordHash: "+fmt.Sprintf("%#v", this.PasswordHash)+",\n")
	s = append(s, "PasswordSalt: "+fmt.Sprintf("%#v", this.PasswordSalt)+",\n")
	s = append(s, "GoogleID: "+fmt.Sprintf("%#v", this.GoogleID)+",\n")
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
		dAtA[i] = 0x50
	}
	if m.CreatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x48
	}
	if len(m.AvatarURL) > 0 {
		i -= len(m.AvatarURL)
		copy(dAtA[i:], m.AvatarURL)
		i = encodeVarintUser(dAtA, i, uint64(len(m.AvatarURL)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x32
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
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
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
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateUser(dAtA, uint64(v5))
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
		case 7:
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
		case 8:
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
		case 9:
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
		case 10:
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
