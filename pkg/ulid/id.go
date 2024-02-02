package ulid

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/gocql/gocql"
	"github.com/oklog/ulid/v2"
)

const (
	length = 16
)

// ID is an alias of ulid.ULID.
type ID []byte

var zero = make(ID, length)

// NewID returns a new random ID.
func NewID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).Bytes())
}

// NewZeroID returns a new zero ID.
func NewZeroID() ID {
	return make(ID, length)
}

// MarshalCQL override marshalling to fit CQL UUID.
func (id ID) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if len(id) != length {
		id = zero
	}

	raw, err := gocql.UUIDFromBytes(id)

	return raw.Bytes(), err
}

// UnmarshalCQL override unmarshalling to fit CQL UUID.
func (id *ID) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	return id.Unmarshal(data)
}

// NewIDs returns a new array of random IDs.
func NewIDs(n int) []ID {
	t := ulid.Timestamp(time.Now())
	ids := make([]ID, n)

	for i := range ids {
		ids[i] = ID(ulid.MustNew(t, rand.Reader).Bytes())
	}

	return ids
}

// NewTimeID returns a new random ID based on time t.
func NewTimeID(t uint64) ID {
	return ID(ulid.MustNew(t, rand.Reader).Bytes())
}

// MustParse alias ulid.MustParse. Panics if s is invalid.
func MustParse(s string) ID {
	return ID(ulid.MustParse(s).Bytes())
}

// Parse alias ulid.Parse. Panics if s is invalid.
func Parse(s string) (ID, error) {
	id, err := ulid.Parse(s)

	return ID(id.Bytes()), err
}

// Time returns ms time of ID.
func (id ID) Time() uint64 {
	var ul ulid.ULID

	if len(id) != length {
		return zero.Time()
	}

	copy(ul[:], id)

	return ul.Time()
}

// IsZero returns if ID is zero value or not.
func (id ID) IsZero() bool {
	return len(id) != length || id[0] == 0
}

// ULID returns original type.
func (id ID) ULID() ulid.ULID {
	var ul ulid.ULID

	if len(id) != length {
		return zero.ULID()
	}

	copy(ul[:], id)

	return ul
}

// Copy returns a hard copy of bytes.
func (id ID) Copy() ID {
	cp := make(ID, length)

	copy(cp[:], id)

	return cp
}

// String returns a human readable string ID.
func (id ID) String() string {
	var ul ulid.ULID

	if len(id) != length {
		return zero.String()
	}

	copy(ul[:], id)

	return ul.String()
}

// Bytes returns id as byte slice for protobuf marshaling.
func (id ID) Bytes() []byte {
	return id[:]
}

// Marshal never returns any error..
func (id ID) Marshal() ([]byte, error) {
	return id, nil
}

// MarshalTo never returns any error.
func (id ID) MarshalTo(data []byte) (n int, err error) {
	copy(data[0:length], id[:])

	return len(id), nil
}

// Unmarshal never returns any error.
func (id *ID) Unmarshal(data []byte) error {
	if len(data) != length {
		return ulid.ErrBufferSize
	}

	if len(*id) != length {
		*id = make(ID, length)
	}

	for i := 0; i < length; i++ {
		(*id)[i] = data[i]
	}

	return nil
}

// Size always returns length.
func (id *ID) Size() int {
	return length
}

// MarshalJSON returns id in human readable string format.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.StdEncoding.EncodeToString(id))
}

// UnmarshalJSON unmarshals and valid data.
func (id *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	raw, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*id = ID(raw) //nolint: wsl

	return nil
}

// Compare only required if the compare option is set.
func (id ID) Compare(other ID) int {
	var ul, ulOther ulid.ULID

	if len(other) != length || len(id) != length {
		return -1
	}

	copy(ul[:], id)
	copy(ulOther[:], other[0:length])

	return ul.Compare(ulOther)
}

// Equal only required if the equal option is set.
func (id ID) Equal(other ID) bool {
	return id.Compare(other) == 0
}

// NewPopulatedID only required if populate option is set.
func NewPopulatedID(r randyID) *ID {
	id := ID(ulid.MustNew(uint64(r.Uint32()), rand.Reader).Bytes())

	return &id
}

type randyID interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}
