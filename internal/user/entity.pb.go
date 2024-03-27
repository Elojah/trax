// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/internal/user/entity.proto

package user

import (
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

type Entity struct {
	ID        github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"ID"`
	Name      string                             `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	AvatarURL string                             `protobuf:"bytes,3,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	CreatedAt int64                              `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64                              `protobuf:"varint,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *Entity) Reset()      { *m = Entity{} }
func (*Entity) ProtoMessage() {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71e0ddd634719b5, []int{0}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return m.Size()
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Entity)(nil), "user.Entity")
	golang_proto.RegisterType((*Entity)(nil), "user.Entity")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/internal/user/entity.proto", fileDescriptor_a71e0ddd634719b5)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/internal/user/entity.proto", fileDescriptor_a71e0ddd634719b5)
}

var fileDescriptor_a71e0ddd634719b5 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0xef, 0x6b, 0x6b, 0xa1, 0xc1, 0x29, 0x53, 0x10, 0xf9, 0x5a, 0x3a, 0x15, 0x87, 0x1c,
	0xe2, 0xe6, 0xd6, 0x1a, 0xc1, 0x80, 0x38, 0x04, 0xfa, 0x03, 0xae, 0xe6, 0x48, 0xa3, 0x69, 0x2e,
	0x9c, 0x17, 0xd1, 0xcd, 0x9f, 0xe0, 0xcf, 0x70, 0x76, 0x72, 0xec, 0x98, 0x31, 0x63, 0x71, 0x28,
	0xe6, 0xb2, 0x38, 0x76, 0x74, 0x94, 0x5c, 0xc5, 0x4e, 0xba, 0xbd, 0x3c, 0xcf, 0xf7, 0xbe, 0xc3,
	0x67, 0x1d, 0x47, 0xb1, 0x9a, 0xe7, 0x33, 0xf7, 0x5a, 0x2c, 0x28, 0x4f, 0xc4, 0x0d, 0x9b, 0x53,
	0x25, 0xd9, 0x03, 0x8d, 0x53, 0xc5, 0x65, 0xca, 0x12, 0x9a, 0xdf, 0x71, 0x49, 0x79, 0xaa, 0x62,
	0xf5, 0xe8, 0x66, 0x52, 0x28, 0x61, 0x77, 0x1a, 0x74, 0x40, 0xff, 0x28, 0x66, 0xb7, 0x11, 0x8d,
	0x44, 0x24, 0xcc, 0xad, 0x49, 0xdb, 0xda, 0xf0, 0x15, 0xac, 0xee, 0xb9, 0xd9, 0xb1, 0x4f, 0xad,
	0x96, 0xef, 0x39, 0x30, 0x80, 0xd1, 0xfe, 0xe4, 0xa8, 0x58, 0xf7, 0xc9, 0xfb, 0xba, 0x3f, 0xfc,
	0x67, 0x2f, 0x4f, 0xe2, 0xd0, 0xf5, 0xbd, 0xa0, 0xe5, 0x7b, 0xb6, 0x6d, 0x75, 0xae, 0xd8, 0x82,
	0x3b, 0xad, 0x01, 0x8c, 0x7a, 0x81, 0xc9, 0xf6, 0xa1, 0xd5, 0x1b, 0xdf, 0x33, 0xc5, 0xe4, 0x34,
	0xb8, 0x74, 0xda, 0x46, 0xec, 0x40, 0x63, 0xcf, 0x24, 0x67, 0x8a, 0x87, 0x63, 0xe5, 0x74, 0x06,
	0x30, 0x6a, 0x07, 0x3b, 0xd0, 0xd8, 0x69, 0x16, 0xfe, 0xd8, 0xbd, 0xad, 0xfd, 0x05, 0x93, 0x8b,
	0xa2, 0x42, 0x52, 0x56, 0x48, 0x56, 0x15, 0x92, 0x4d, 0x85, 0xf0, 0x55, 0x21, 0x3c, 0x69, 0x84,
	0x17, 0x8d, 0xf0, 0xa6, 0x11, 0x96, 0x1a, 0xa1, 0xd0, 0x08, 0xa5, 0x46, 0xf8, 0xd0, 0x08, 0x9f,
	0x1a, 0xc9, 0x46, 0x23, 0x3c, 0xd7, 0x48, 0x96, 0x35, 0x42, 0x59, 0x23, 0x59, 0xd5, 0x48, 0x66,
	0x5d, 0xf3, 0x85, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xa8, 0xaf, 0x0f, 0x71, 0x01,
	0x00, 0x00,
}

func (this *Entity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Entity)
	if !ok {
		that2, ok := that.(Entity)
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
	if this.Name != that1.Name {
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
func (this *Entity) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&user.Entity{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "AvatarURL: "+fmt.Sprintf("%#v", this.AvatarURL)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "UpdatedAt: "+fmt.Sprintf("%#v", this.UpdatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEntity(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Entity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x28
	}
	if m.CreatedAt != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AvatarURL) > 0 {
		i -= len(m.AvatarURL)
		copy(dAtA[i:], m.AvatarURL)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.AvatarURL)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEntity(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedEntity(r randyEntity, easy bool) *Entity {
	this := &Entity{}
	v1 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	this.Name = string(randStringEntity(r))
	this.AvatarURL = string(randStringEntity(r))
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

type randyEntity interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEntity(r randyEntity) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEntity(r randyEntity) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneEntity(r)
	}
	return string(tmps)
}
func randUnrecognizedEntity(r randyEntity, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEntity(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEntity(dAtA []byte, r randyEntity, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEntity(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Entity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	l = len(m.AvatarURL)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovEntity(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovEntity(uint64(m.UpdatedAt))
	}
	return n
}

func sovEntity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntity(x uint64) (n int) {
	return sovEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Entity) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Entity{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`AvatarURL:` + fmt.Sprintf("%v", this.AvatarURL) + `,`,
		`CreatedAt:` + fmt.Sprintf("%v", this.CreatedAt) + `,`,
		`UpdatedAt:` + fmt.Sprintf("%v", this.UpdatedAt) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEntity(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Entity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
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
			return fmt.Errorf("proto: Entity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
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
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEntity
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
func skipEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
					return 0, ErrIntOverflowEntity
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
				return 0, ErrInvalidLengthEntity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntity = fmt.Errorf("proto: unexpected end of group")
)