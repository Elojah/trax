// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/internal/user/dto/profile.proto

package dto

import (
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
	pbtypes "github.com/elojah/trax/pkg/pbtypes"
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

type FetchProfileReq struct {
	Me       bool                               `protobuf:"varint,1,opt,name=Me,proto3" json:"Me,omitempty"`
	UserID   github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,2,opt,name=UserID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"UserID"`
	EntityID github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,3,opt,name=EntityID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"EntityID"`
}

func (m *FetchProfileReq) Reset()      { *m = FetchProfileReq{} }
func (*FetchProfileReq) ProtoMessage() {}
func (*FetchProfileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a6d37e631abfc2, []int{0}
}
func (m *FetchProfileReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchProfileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchProfileReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchProfileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchProfileReq.Merge(m, src)
}
func (m *FetchProfileReq) XXX_Size() int {
	return m.Size()
}
func (m *FetchProfileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchProfileReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchProfileReq proto.InternalMessageInfo

type UpdateProfileReq struct {
	Firstname *pbtypes.String `protobuf:"bytes,1,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  *pbtypes.String `protobuf:"bytes,2,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	AvatarURL *pbtypes.String `protobuf:"bytes,3,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
}

func (m *UpdateProfileReq) Reset()      { *m = UpdateProfileReq{} }
func (*UpdateProfileReq) ProtoMessage() {}
func (*UpdateProfileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a6d37e631abfc2, []int{1}
}
func (m *UpdateProfileReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateProfileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateProfileReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateProfileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileReq.Merge(m, src)
}
func (m *UpdateProfileReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateProfileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FetchProfileReq)(nil), "dto.FetchProfileReq")
	golang_proto.RegisterType((*FetchProfileReq)(nil), "dto.FetchProfileReq")
	proto.RegisterType((*UpdateProfileReq)(nil), "dto.UpdateProfileReq")
	golang_proto.RegisterType((*UpdateProfileReq)(nil), "dto.UpdateProfileReq")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/internal/user/dto/profile.proto", fileDescriptor_b4a6d37e631abfc2)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/internal/user/dto/profile.proto", fileDescriptor_b4a6d37e631abfc2)
}

var fileDescriptor_b4a6d37e631abfc2 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x6f, 0xda, 0x40,
	0x18, 0x86, 0xef, 0x33, 0x12, 0xa2, 0xd7, 0xaa, 0x54, 0x9e, 0x50, 0x87, 0x0f, 0xc4, 0x84, 0x5a,
	0xd5, 0x27, 0xb5, 0x52, 0xf7, 0x22, 0x8a, 0x8a, 0x04, 0x52, 0xe5, 0x8a, 0x1f, 0x70, 0xe0, 0xab,
	0x71, 0x6b, 0x7c, 0xee, 0xf9, 0xa8, 0xc2, 0x96, 0x9f, 0x90, 0x3f, 0x90, 0x3d, 0x4b, 0xf6, 0x8c,
	0x8c, 0x8c, 0x8c, 0x28, 0x03, 0x8a, 0xcf, 0x4b, 0x46, 0xc6, 0x8c, 0x51, 0x6c, 0x44, 0x32, 0x10,
	0xa4, 0x6c, 0x37, 0x3c, 0xcf, 0xfb, 0xbd, 0xaf, 0x8e, 0x7e, 0xf5, 0x03, 0x3d, 0x99, 0x8d, 0x9c,
	0xb1, 0x9c, 0x32, 0x11, 0xca, 0x3f, 0x7c, 0xc2, 0xb4, 0xe2, 0x27, 0x2c, 0x88, 0xb4, 0x50, 0x11,
	0x0f, 0xd9, 0x2c, 0x11, 0x8a, 0x79, 0x5a, 0xb2, 0x58, 0xc9, 0xdf, 0x41, 0x28, 0x9c, 0x58, 0x49,
	0x2d, 0xed, 0x92, 0xa7, 0xe5, 0x7b, 0xf6, 0x8c, 0x1c, 0xff, 0xf5, 0x99, 0x2f, 0x7d, 0x99, 0xa3,
	0xf9, 0xab, 0xb0, 0x8e, 0x0a, 0xf1, 0x48, 0xcf, 0x63, 0x91, 0xb0, 0x44, 0xab, 0x20, 0xf2, 0x0b,
	0xa1, 0x79, 0x09, 0xb4, 0xda, 0x15, 0x7a, 0x3c, 0xf9, 0x59, 0x5c, 0x77, 0xc5, 0x3f, 0xfb, 0x2d,
	0xb5, 0x06, 0xa2, 0x06, 0x0d, 0x68, 0x55, 0x5c, 0x6b, 0x20, 0xec, 0x36, 0x2d, 0x0f, 0x13, 0xa1,
	0x7a, 0x9d, 0x9a, 0xd5, 0x80, 0xd6, 0x9b, 0xf6, 0x87, 0xe5, 0xa6, 0x4e, 0xae, 0x37, 0xf5, 0xe6,
	0x91, 0x63, 0xb3, 0x30, 0xf0, 0x9c, 0x5e, 0xc7, 0xdd, 0x99, 0x76, 0x97, 0x56, 0xbe, 0x47, 0x3a,
	0xd0, 0xf3, 0x5e, 0xa7, 0x56, 0x7a, 0x71, 0xca, 0xde, 0x6d, 0x9e, 0x03, 0x7d, 0x37, 0x8c, 0x3d,
	0xae, 0xc5, 0x93, 0xc2, 0x9f, 0xe8, 0xab, 0x6e, 0xa0, 0x12, 0x1d, 0xf1, 0x69, 0xd1, 0xfb, 0xf5,
	0xe7, 0xaa, 0xb3, 0x9b, 0xeb, 0xfc, 0xca, 0xe7, 0xba, 0x8f, 0x84, 0xfd, 0x91, 0x56, 0xfa, 0x7c,
	0x47, 0x5b, 0x87, 0xe9, 0x3d, 0xf0, 0x90, 0xfd, 0xed, 0x3f, 0xd7, 0x5c, 0x0d, 0xdd, 0x7e, 0xde,
	0xfc, 0x50, 0xf6, 0x9e, 0x68, 0xff, 0x58, 0xa6, 0x48, 0x56, 0x29, 0x92, 0x75, 0x8a, 0x64, 0x9b,
	0x22, 0xdc, 0xa5, 0x08, 0xa7, 0x06, 0xe1, 0xc2, 0x20, 0x5c, 0x19, 0x84, 0x85, 0x41, 0x58, 0x1a,
	0x84, 0x95, 0x41, 0xb8, 0x31, 0x08, 0xb7, 0x06, 0xc9, 0xd6, 0x20, 0x9c, 0x65, 0x48, 0x16, 0x19,
	0xc2, 0x2a, 0x43, 0xb2, 0xce, 0x90, 0x8c, 0xca, 0xf9, 0x07, 0x7d, 0xb9, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x76, 0xf7, 0x7e, 0x84, 0x41, 0x02, 0x00, 0x00,
}

func (this *FetchProfileReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FetchProfileReq)
	if !ok {
		that2, ok := that.(FetchProfileReq)
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
	if this.Me != that1.Me {
		return false
	}
	if !this.UserID.Equal(that1.UserID) {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	return true
}
func (this *UpdateProfileReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateProfileReq)
	if !ok {
		that2, ok := that.(UpdateProfileReq)
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
	if !this.Firstname.Equal(that1.Firstname) {
		return false
	}
	if !this.Lastname.Equal(that1.Lastname) {
		return false
	}
	if !this.AvatarURL.Equal(that1.AvatarURL) {
		return false
	}
	return true
}
func (this *FetchProfileReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dto.FetchProfileReq{")
	s = append(s, "Me: "+fmt.Sprintf("%#v", this.Me)+",\n")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdateProfileReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dto.UpdateProfileReq{")
	if this.Firstname != nil {
		s = append(s, "Firstname: "+fmt.Sprintf("%#v", this.Firstname)+",\n")
	}
	if this.Lastname != nil {
		s = append(s, "Lastname: "+fmt.Sprintf("%#v", this.Lastname)+",\n")
	}
	if this.AvatarURL != nil {
		s = append(s, "AvatarURL: "+fmt.Sprintf("%#v", this.AvatarURL)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProfile(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *FetchProfileReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchProfileReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchProfileReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProfile(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UserID.Size()
		i -= size
		if _, err := m.UserID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProfile(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Me {
		i--
		if m.Me {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateProfileReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateProfileReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateProfileReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AvatarURL != nil {
		{
			size, err := m.AvatarURL.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Lastname != nil {
		{
			size, err := m.Lastname.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Firstname != nil {
		{
			size, err := m.Firstname.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedFetchProfileReq(r randyProfile, easy bool) *FetchProfileReq {
	this := &FetchProfileReq{}
	this.Me = bool(bool(r.Intn(2) == 0))
	v1 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.UserID = *v1
	v2 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedUpdateProfileReq(r randyProfile, easy bool) *UpdateProfileReq {
	this := &UpdateProfileReq{}
	if r.Intn(5) != 0 {
		this.Firstname = pbtypes.NewPopulatedString(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Lastname = pbtypes.NewPopulatedString(r, easy)
	}
	if r.Intn(5) != 0 {
		this.AvatarURL = pbtypes.NewPopulatedString(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyProfile interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProfile(r randyProfile) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProfile(r randyProfile) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneProfile(r)
	}
	return string(tmps)
}
func randUnrecognizedProfile(r randyProfile, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProfile(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProfile(dAtA []byte, r randyProfile, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProfile(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *FetchProfileReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Me {
		n += 2
	}
	l = m.UserID.Size()
	n += 1 + l + sovProfile(uint64(l))
	l = m.EntityID.Size()
	n += 1 + l + sovProfile(uint64(l))
	return n
}

func (m *UpdateProfileReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Firstname != nil {
		l = m.Firstname.Size()
		n += 1 + l + sovProfile(uint64(l))
	}
	if m.Lastname != nil {
		l = m.Lastname.Size()
		n += 1 + l + sovProfile(uint64(l))
	}
	if m.AvatarURL != nil {
		l = m.AvatarURL.Size()
		n += 1 + l + sovProfile(uint64(l))
	}
	return n
}

func sovProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProfile(x uint64) (n int) {
	return sovProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FetchProfileReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FetchProfileReq{`,
		`Me:` + fmt.Sprintf("%v", this.Me) + `,`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateProfileReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateProfileReq{`,
		`Firstname:` + strings.Replace(fmt.Sprintf("%v", this.Firstname), "String", "pbtypes.String", 1) + `,`,
		`Lastname:` + strings.Replace(fmt.Sprintf("%v", this.Lastname), "String", "pbtypes.String", 1) + `,`,
		`AvatarURL:` + strings.Replace(fmt.Sprintf("%v", this.AvatarURL), "String", "pbtypes.String", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProfile(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FetchProfileReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfile
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
			return fmt.Errorf("proto: FetchProfileReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchProfileReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Me", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Me = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
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
				return ErrInvalidLengthProfile
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
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
				return ErrInvalidLengthProfile
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProfile
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
func (m *UpdateProfileReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProfile
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
			return fmt.Errorf("proto: UpdateProfileReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateProfileReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Firstname", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Firstname == nil {
				m.Firstname = &pbtypes.String{}
			}
			if err := m.Firstname.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lastname", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lastname == nil {
				m.Lastname = &pbtypes.String{}
			}
			if err := m.Lastname.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarURL", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AvatarURL == nil {
				m.AvatarURL = &pbtypes.String{}
			}
			if err := m.AvatarURL.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProfile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProfile
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
func skipProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProfile
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
					return 0, ErrIntOverflowProfile
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
					return 0, ErrIntOverflowProfile
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
				return 0, ErrInvalidLengthProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProfile = fmt.Errorf("proto: unexpected end of group")
)
