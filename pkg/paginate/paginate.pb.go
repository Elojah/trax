// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/pkg/paginate/paginate.proto

package paginate

import (
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
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

type Paginate struct {
	Start int64  `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
	End   int64  `protobuf:"varint,2,opt,name=End,proto3" json:"End,omitempty"`
	Sort  string `protobuf:"bytes,3,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Order bool   `protobuf:"varint,4,opt,name=Order,proto3" json:"Order,omitempty"`
}

func (m *Paginate) Reset()      { *m = Paginate{} }
func (*Paginate) ProtoMessage() {}
func (*Paginate) Descriptor() ([]byte, []int) {
	return fileDescriptor_57523d5c2e382b82, []int{0}
}
func (m *Paginate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Paginate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Paginate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Paginate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paginate.Merge(m, src)
}
func (m *Paginate) XXX_Size() int {
	return m.Size()
}
func (m *Paginate) XXX_DiscardUnknown() {
	xxx_messageInfo_Paginate.DiscardUnknown(m)
}

var xxx_messageInfo_Paginate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Paginate)(nil), "paginate.Paginate")
	golang_proto.RegisterType((*Paginate)(nil), "paginate.Paginate")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/pkg/paginate/paginate.proto", fileDescriptor_57523d5c2e382b82)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/pkg/paginate/paginate.proto", fileDescriptor_57523d5c2e382b82)
}

var fileDescriptor_57523d5c2e382b82 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x2f, 0x29,
	0x4a, 0xac, 0xd0, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0x48, 0x4c, 0xcf, 0xcc, 0x4b, 0x2c, 0x49, 0x85,
	0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60, 0x7c, 0x29, 0x7d, 0x3c, 0xba, 0xd3,
	0xf3, 0xd3, 0xf3, 0xc1, 0xea, 0xc1, 0x2c, 0x88, 0x56, 0xa5, 0x18, 0x2e, 0x8e, 0x00, 0xa8, 0x66,
	0x21, 0x11, 0x2e, 0xd6, 0xe0, 0x92, 0xc4, 0xa2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x08, 0x47, 0x48, 0x80, 0x8b, 0xd9, 0x35, 0x2f, 0x45, 0x82, 0x09, 0x2c, 0x06, 0x62, 0x0a, 0x09,
	0x71, 0xb1, 0x04, 0xe7, 0x17, 0x95, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20,
	0xbd, 0xfe, 0x45, 0x29, 0xa9, 0x45, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x10, 0x8e, 0x93,
	0xc7, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xf0, 0xe1, 0xa1,
	0x1c, 0xe3, 0x8f, 0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1,
	0x48, 0x8e, 0xf1, 0xc0, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0x38, 0xf0, 0x58, 0x8e, 0xf1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0,
	0xce, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x3d, 0x2d, 0x92, 0x1f, 0x01, 0x00, 0x00,
}

func (this *Paginate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Paginate)
	if !ok {
		that2, ok := that.(Paginate)
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
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if this.Sort != that1.Sort {
		return false
	}
	if this.Order != that1.Order {
		return false
	}
	return true
}
func (this *Paginate) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&paginate.Paginate{")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "Sort: "+fmt.Sprintf("%#v", this.Sort)+",\n")
	s = append(s, "Order: "+fmt.Sprintf("%#v", this.Order)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPaginate(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Paginate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Paginate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Paginate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order {
		i--
		if m.Order {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Sort) > 0 {
		i -= len(m.Sort)
		copy(dAtA[i:], m.Sort)
		i = encodeVarintPaginate(dAtA, i, uint64(len(m.Sort)))
		i--
		dAtA[i] = 0x1a
	}
	if m.End != 0 {
		i = encodeVarintPaginate(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintPaginate(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaginate(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaginate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedPaginate(r randyPaginate, easy bool) *Paginate {
	this := &Paginate{}
	this.Start = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Start *= -1
	}
	this.End = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.End *= -1
	}
	this.Sort = string(randStringPaginate(r))
	this.Order = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPaginate interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePaginate(r randyPaginate) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPaginate(r randyPaginate) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RunePaginate(r)
	}
	return string(tmps)
}
func randUnrecognizedPaginate(r randyPaginate, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPaginate(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPaginate(dAtA []byte, r randyPaginate, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePaginate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePaginate(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Paginate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovPaginate(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovPaginate(uint64(m.End))
	}
	l = len(m.Sort)
	if l > 0 {
		n += 1 + l + sovPaginate(uint64(l))
	}
	if m.Order {
		n += 2
	}
	return n
}

func sovPaginate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaginate(x uint64) (n int) {
	return sovPaginate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Paginate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Paginate{`,
		`Start:` + fmt.Sprintf("%v", this.Start) + `,`,
		`End:` + fmt.Sprintf("%v", this.End) + `,`,
		`Sort:` + fmt.Sprintf("%v", this.Sort) + `,`,
		`Order:` + fmt.Sprintf("%v", this.Order) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPaginate(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Paginate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaginate
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
			return fmt.Errorf("proto: Paginate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Paginate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaginate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaginate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaginate
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
				return ErrInvalidLengthPaginate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaginate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaginate
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
			m.Order = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPaginate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPaginate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPaginate
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
func skipPaginate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaginate
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
					return 0, ErrIntOverflowPaginate
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
					return 0, ErrIntOverflowPaginate
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
				return 0, ErrInvalidLengthPaginate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaginate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaginate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaginate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaginate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaginate = fmt.Errorf("proto: unexpected end of group")
)