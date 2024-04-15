// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/trax/internal/user/claims.proto

package user

import (
	fmt "fmt"
	_ "github.com/elojah/trax/pkg/gogoproto"
	github_com_elojah_trax_pkg_ulid "github.com/elojah/trax/pkg/ulid"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
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

type ClaimCommands struct {
	Commands []Command `protobuf:"varint,1,rep,packed,name=Commands,proto3,enum=user.Command" json:"Commands,omitempty"`
}

func (m *ClaimCommands) Reset()      { *m = ClaimCommands{} }
func (*ClaimCommands) ProtoMessage() {}
func (*ClaimCommands) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ce1bb4770ff0ca, []int{0}
}
func (m *ClaimCommands) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimCommands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimCommands.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimCommands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimCommands.Merge(m, src)
}
func (m *ClaimCommands) XXX_Size() int {
	return m.Size()
}
func (m *ClaimCommands) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimCommands.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimCommands proto.InternalMessageInfo

type ClaimEntity struct {
	ID github_com_elojah_trax_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/trax/pkg/ulid.ID" json:"ID"`
	// Permissions key must be a user.Resource
	Commands map[string]ClaimCommands `protobuf:"bytes,2,rep,name=Commands,proto3" json:"Commands" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ClaimEntity) Reset()      { *m = ClaimEntity{} }
func (*ClaimEntity) ProtoMessage() {}
func (*ClaimEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ce1bb4770ff0ca, []int{1}
}
func (m *ClaimEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimEntity.Merge(m, src)
}
func (m *ClaimEntity) XXX_Size() int {
	return m.Size()
}
func (m *ClaimEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimEntity proto.InternalMessageInfo

type ClaimAuth struct {
	// Permissions key must be a entity.ID
	Entities map[string]ClaimEntity `protobuf:"bytes,1,rep,name=Entities,proto3" json:"Entities" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ClaimAuth) Reset()      { *m = ClaimAuth{} }
func (*ClaimAuth) ProtoMessage() {}
func (*ClaimAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ce1bb4770ff0ca, []int{2}
}
func (m *ClaimAuth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimAuth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimAuth.Merge(m, src)
}
func (m *ClaimAuth) XXX_Size() int {
	return m.Size()
}
func (m *ClaimAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimAuth proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClaimCommands)(nil), "user.ClaimCommands")
	golang_proto.RegisterType((*ClaimCommands)(nil), "user.ClaimCommands")
	proto.RegisterType((*ClaimEntity)(nil), "user.ClaimEntity")
	golang_proto.RegisterType((*ClaimEntity)(nil), "user.ClaimEntity")
	proto.RegisterMapType((map[string]ClaimCommands)(nil), "user.ClaimEntity.CommandsEntry")
	golang_proto.RegisterMapType((map[string]ClaimCommands)(nil), "user.ClaimEntity.CommandsEntry")
	proto.RegisterType((*ClaimAuth)(nil), "user.ClaimAuth")
	golang_proto.RegisterType((*ClaimAuth)(nil), "user.ClaimAuth")
	proto.RegisterMapType((map[string]ClaimEntity)(nil), "user.ClaimAuth.EntitiesEntry")
	golang_proto.RegisterMapType((map[string]ClaimEntity)(nil), "user.ClaimAuth.EntitiesEntry")
}

func init() {
	proto.RegisterFile("github.com/elojah/trax/internal/user/claims.proto", fileDescriptor_f8ce1bb4770ff0ca)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/trax/internal/user/claims.proto", fileDescriptor_f8ce1bb4770ff0ca)
}

var fileDescriptor_f8ce1bb4770ff0ca = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0x93, 0x50,
	0x14, 0x86, 0xef, 0xa1, 0xa3, 0x99, 0xb9, 0x88, 0x51, 0xdc, 0x34, 0x4d, 0x3c, 0x25, 0x6c, 0x64,
	0x5c, 0x40, 0xc4, 0x8d, 0xe9, 0xc6, 0x4c, 0xa7, 0x93, 0xd8, 0x8d, 0x31, 0xbc, 0x01, 0x33, 0x43,
	0x28, 0x0e, 0x70, 0x1b, 0xb8, 0x18, 0xbb, 0xf3, 0x11, 0x7c, 0x00, 0x1f, 0xc0, 0x47, 0x70, 0xd9,
	0x65, 0x97, 0x5d, 0x36, 0x2e, 0x1a, 0xb9, 0x6c, 0x5c, 0x76, 0xe9, 0xd2, 0x00, 0xc5, 0x42, 0xb4,
	0xee, 0x4e, 0xce, 0xf9, 0xbf, 0xfc, 0x1f, 0x37, 0xd0, 0x17, 0x7e, 0xc0, 0x67, 0xd9, 0xb5, 0x79,
	0xc3, 0x22, 0xcb, 0x0b, 0xd9, 0x7b, 0x77, 0x66, 0xf1, 0xc4, 0xfd, 0x68, 0x05, 0x31, 0xf7, 0x92,
	0xd8, 0x0d, 0xad, 0x2c, 0xf5, 0x12, 0xeb, 0x26, 0x74, 0x83, 0x28, 0x35, 0xe7, 0x09, 0xe3, 0x4c,
	0x3d, 0x29, 0x57, 0x03, 0xeb, 0x08, 0x38, 0xbf, 0xf3, 0x2d, 0x9f, 0xf9, 0xac, 0xca, 0x56, 0x53,
	0x8d, 0x1d, 0x05, 0xba, 0x4d, 0x09, 0x0b, 0xbd, 0x1a, 0xd0, 0x47, 0x54, 0xb9, 0x2c, 0x7b, 0x2f,
	0x59, 0x14, 0xb9, 0xf1, 0x6d, 0xaa, 0x9e, 0xd3, 0xd3, 0x66, 0xee, 0x83, 0xd6, 0x33, 0x1e, 0xda,
	0x8a, 0x59, 0x42, 0xe6, 0x7e, 0xeb, 0xfc, 0x39, 0xeb, 0x5b, 0xa0, 0x72, 0x05, 0x5f, 0xc5, 0x3c,
	0xe0, 0x0b, 0x75, 0x44, 0xa5, 0xe9, 0xa4, 0x0f, 0x1a, 0x18, 0x0f, 0xc6, 0xcf, 0x57, 0xdb, 0x21,
	0xf9, 0xbe, 0x1d, 0xea, 0xff, 0xf9, 0x82, 0x2c, 0x0c, 0x6e, 0xcd, 0xe9, 0xc4, 0x91, 0xa6, 0x13,
	0xf5, 0xa2, 0x55, 0x2b, 0x69, 0x3d, 0x43, 0xb6, 0x87, 0xfb, 0xda, 0x43, 0x41, 0xa3, 0x90, 0x5e,
	0xc5, 0x3c, 0x59, 0x8c, 0x4f, 0xca, 0x8a, 0x83, 0xce, 0xe0, 0x1d, 0x55, 0x3a, 0x01, 0xf5, 0x11,
	0xed, 0xdd, 0x79, 0x8b, 0x4a, 0xe8, 0xcc, 0x29, 0x47, 0xf5, 0x9c, 0xde, 0xfb, 0xe0, 0x86, 0x99,
	0xd7, 0x97, 0x34, 0x30, 0x64, 0xfb, 0x49, 0xab, 0xa2, 0x41, 0x9d, 0x3a, 0x31, 0x92, 0x5e, 0x81,
	0xfe, 0x05, 0xe8, 0x59, 0x75, 0xbc, 0xc8, 0xf8, 0x4c, 0x7d, 0x4d, 0x4f, 0x2b, 0x8f, 0xc0, 0xab,
	0x5f, 0x46, 0xb6, 0x9f, 0xb6, 0xf8, 0x32, 0x62, 0x36, 0xf7, 0x8e, 0x60, 0xb3, 0x1c, 0xbc, 0xa5,
	0x4a, 0x27, 0xf0, 0x0f, 0xc1, 0x67, 0x5d, 0xc1, 0xc7, 0x7f, 0xbd, 0x41, 0x4b, 0x6f, 0xfc, 0x66,
	0x95, 0x23, 0x59, 0xe7, 0x48, 0x36, 0x39, 0x92, 0x5d, 0x8e, 0xf0, 0x2b, 0x47, 0xf8, 0x24, 0x10,
	0xbe, 0x0a, 0x84, 0x6f, 0x02, 0x61, 0x29, 0x10, 0x56, 0x02, 0x61, 0x2d, 0x10, 0x7e, 0x08, 0x84,
	0x9f, 0x02, 0xc9, 0x4e, 0x20, 0x7c, 0x2e, 0x90, 0x2c, 0x0b, 0x84, 0x75, 0x81, 0x64, 0x53, 0x20,
	0xb9, 0xbe, 0x5f, 0xfd, 0x0c, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x89, 0x87, 0x36, 0xda,
	0xa9, 0x02, 0x00, 0x00,
}

func (this *ClaimCommands) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimCommands)
	if !ok {
		that2, ok := that.(ClaimCommands)
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
	if len(this.Commands) != len(that1.Commands) {
		return false
	}
	for i := range this.Commands {
		if this.Commands[i] != that1.Commands[i] {
			return false
		}
	}
	return true
}
func (this *ClaimEntity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimEntity)
	if !ok {
		that2, ok := that.(ClaimEntity)
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
	if len(this.Commands) != len(that1.Commands) {
		return false
	}
	for i := range this.Commands {
		a := this.Commands[i]
		b := that1.Commands[i]
		if !(&a).Equal(&b) {
			return false
		}
	}
	return true
}
func (this *ClaimAuth) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimAuth)
	if !ok {
		that2, ok := that.(ClaimAuth)
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
	if len(this.Entities) != len(that1.Entities) {
		return false
	}
	for i := range this.Entities {
		a := this.Entities[i]
		b := that1.Entities[i]
		if !(&a).Equal(&b) {
			return false
		}
	}
	return true
}
func (this *ClaimCommands) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&user.ClaimCommands{")
	s = append(s, "Commands: "+fmt.Sprintf("%#v", this.Commands)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClaimEntity) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&user.ClaimEntity{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	keysForCommands := make([]string, 0, len(this.Commands))
	for k, _ := range this.Commands {
		keysForCommands = append(keysForCommands, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForCommands)
	mapStringForCommands := "map[string]ClaimCommands{"
	for _, k := range keysForCommands {
		mapStringForCommands += fmt.Sprintf("%#v: %#v,", k, this.Commands[k])
	}
	mapStringForCommands += "}"
	if this.Commands != nil {
		s = append(s, "Commands: "+mapStringForCommands+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClaimAuth) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&user.ClaimAuth{")
	keysForEntities := make([]string, 0, len(this.Entities))
	for k, _ := range this.Entities {
		keysForEntities = append(keysForEntities, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForEntities)
	mapStringForEntities := "map[string]ClaimEntity{"
	for _, k := range keysForEntities {
		mapStringForEntities += fmt.Sprintf("%#v: %#v,", k, this.Entities[k])
	}
	mapStringForEntities += "}"
	if this.Entities != nil {
		s = append(s, "Entities: "+mapStringForEntities+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClaims(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ClaimCommands) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimCommands) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimCommands) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commands) > 0 {
		dAtA2 := make([]byte, len(m.Commands)*10)
		var j1 int
		for _, num := range m.Commands {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintClaims(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commands) > 0 {
		for k := range m.Commands {
			v := m.Commands[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaims(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintClaims(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintClaims(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClaimAuth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimAuth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimAuth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entities) > 0 {
		for k := range m.Entities {
			v := m.Entities[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaims(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintClaims(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintClaims(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaims(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaims(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedClaimCommands(r randyClaims, easy bool) *ClaimCommands {
	this := &ClaimCommands{}
	v1 := r.Intn(10)
	this.Commands = make([]Command, v1)
	for i := 0; i < v1; i++ {
		this.Commands[i] = Command([]int32{0, 1, 2, 3}[r.Intn(4)])
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClaimEntity(r randyClaims, easy bool) *ClaimEntity {
	this := &ClaimEntity{}
	v2 := github_com_elojah_trax_pkg_ulid.NewPopulatedID(r)
	this.ID = *v2
	if r.Intn(5) != 0 {
		v3 := r.Intn(10)
		this.Commands = make(map[string]ClaimCommands)
		for i := 0; i < v3; i++ {
			this.Commands[randStringClaims(r)] = *NewPopulatedClaimCommands(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClaimAuth(r randyClaims, easy bool) *ClaimAuth {
	this := &ClaimAuth{}
	if r.Intn(5) != 0 {
		v4 := r.Intn(10)
		this.Entities = make(map[string]ClaimEntity)
		for i := 0; i < v4; i++ {
			this.Entities[randStringClaims(r)] = *NewPopulatedClaimEntity(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyClaims interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneClaims(r randyClaims) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringClaims(r randyClaims) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneClaims(r)
	}
	return string(tmps)
}
func randUnrecognizedClaims(r randyClaims, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldClaims(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldClaims(dAtA []byte, r randyClaims, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateClaims(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateClaims(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ClaimCommands) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Commands) > 0 {
		l = 0
		for _, e := range m.Commands {
			l += sovClaims(uint64(e))
		}
		n += 1 + sovClaims(uint64(l)) + l
	}
	return n
}

func (m *ClaimEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovClaims(uint64(l))
	if len(m.Commands) > 0 {
		for k, v := range m.Commands {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovClaims(uint64(len(k))) + 1 + l + sovClaims(uint64(l))
			n += mapEntrySize + 1 + sovClaims(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ClaimAuth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entities) > 0 {
		for k, v := range m.Entities {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovClaims(uint64(len(k))) + 1 + l + sovClaims(uint64(l))
			n += mapEntrySize + 1 + sovClaims(uint64(mapEntrySize))
		}
	}
	return n
}

func sovClaims(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaims(x uint64) (n int) {
	return sovClaims(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClaimCommands) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClaimCommands{`,
		`Commands:` + fmt.Sprintf("%v", this.Commands) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClaimEntity) String() string {
	if this == nil {
		return "nil"
	}
	keysForCommands := make([]string, 0, len(this.Commands))
	for k, _ := range this.Commands {
		keysForCommands = append(keysForCommands, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForCommands)
	mapStringForCommands := "map[string]ClaimCommands{"
	for _, k := range keysForCommands {
		mapStringForCommands += fmt.Sprintf("%v: %v,", k, this.Commands[k])
	}
	mapStringForCommands += "}"
	s := strings.Join([]string{`&ClaimEntity{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Commands:` + mapStringForCommands + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClaimAuth) String() string {
	if this == nil {
		return "nil"
	}
	keysForEntities := make([]string, 0, len(this.Entities))
	for k, _ := range this.Entities {
		keysForEntities = append(keysForEntities, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForEntities)
	mapStringForEntities := "map[string]ClaimEntity{"
	for _, k := range keysForEntities {
		mapStringForEntities += fmt.Sprintf("%v: %v,", k, this.Entities[k])
	}
	mapStringForEntities += "}"
	s := strings.Join([]string{`&ClaimAuth{`,
		`Entities:` + mapStringForEntities + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClaims(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClaimCommands) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: ClaimCommands: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimCommands: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v Command
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Command(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Commands = append(m.Commands, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaims
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaims
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Commands) == 0 {
					m.Commands = make([]Command, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Command
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Command(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Commands = append(m.Commands, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Commands", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClaims
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
func (m *ClaimEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: ClaimEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
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
				return fmt.Errorf("proto: wrong wireType = %d for field Commands", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commands == nil {
				m.Commands = make(map[string]ClaimCommands)
			}
			var mapkey string
			mapvalue := &ClaimCommands{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthClaims
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthClaims
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthClaims
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthClaims
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ClaimCommands{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipClaims(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthClaims
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Commands[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClaims
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
func (m *ClaimAuth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: ClaimAuth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimAuth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entities == nil {
				m.Entities = make(map[string]ClaimEntity)
			}
			var mapkey string
			mapvalue := &ClaimEntity{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthClaims
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthClaims
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthClaims
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthClaims
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ClaimEntity{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipClaims(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthClaims
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Entities[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClaims
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
func skipClaims(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaims
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
					return 0, ErrIntOverflowClaims
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
					return 0, ErrIntOverflowClaims
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
				return 0, ErrInvalidLengthClaims
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaims
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaims
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaims        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaims          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaims = fmt.Errorf("proto: unexpected end of group")
)