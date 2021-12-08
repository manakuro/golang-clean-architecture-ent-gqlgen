package ulid

import (
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

// ID implements a ULID
type ID string

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	// Seed the default entropy source.
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

// newULID returns a new ULID for time.Now() using the default entropy source.
func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

// MustNew returns a new ULID for time.Now() given a prefix. This uses the default entropy source.
func MustNew(prefix string) ID {
	return ID(prefix + fmt.Sprint(newULID()))
}

// UnmarshalGQL implements the graphql.Unmarshaller interface.
func (i *ID) UnmarshalGQL(v interface{}) error {
	return i.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface.
func (i ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(i)))
}

// Scan implements the Scanner interface.
func (i *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("ulid: expected a value")
	}

	switch s := src.(type) {
	case string:
		*i = ID(s)
	case []byte:
		str := string(s)
		*i = ID(str)
	default:
		return fmt.Errorf("ulid: expected a string %v", s)
	}

	return nil
}

// Value implements the driver Valuer interface.
func (i ID) Value() (driver.Value, error) {
	return string(i), nil
}
