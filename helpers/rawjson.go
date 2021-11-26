package helpers

import (
	"database/sql/driver"
	"errors"
)

// Type RawJSON represents a JSON encoded []byte slice that the application does not need to decode.
// Useful to pass JSON data between the client and the dbase without decoding/encoding at any point.
// To avoid breaking JSON encoding with empty values, use a pointer (*utilities.RawJSON) or apply the tag `json:",omitempty"`.
type RawJSON []byte

// MarshalJSON returns m as the JSON encoding of m.
// MarshalJSON implements the json.Marshaler interface.
func (m RawJSON) MarshalJSON() ([]byte, error) {
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
// MarshalJSON implements the json.Unmarshaler interface.
func (m *RawJSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawJSON: UnmarshalJSON on nil pointer")
	}
	*m = make([]byte, 0)
	*m = append((*m)[0:0], data...)
	return nil
}

// New returns a new *RawJSON variable.
func New() *RawJSON {
	return new(RawJSON)
}

// Download from SQL database.
// Scan implements the sql.Scanner interface.
func (m *RawJSON) Scan(value interface{}) error {
	if m == nil {
		return errors.New("RawJSON: Scan on nil pointer")
	}
	if value != nil {
		data := value.([]byte)
		*m = make([]byte, 0)
		*m = append((*m)[0:0], data...)
	}
	return nil
}

// Insert in SQL database.
// Value implements the driver.Valuer interface.
func (m *RawJSON) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return []byte(*m), nil
}
