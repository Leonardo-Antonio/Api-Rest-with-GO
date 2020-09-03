package storage

import (
	"database/sql"
	"errors"
)

var (
	// ErrorRowAffected -> rows affected +1
	ErrorRowAffected = errors.New("Se afectaron mas de dos filas")
)

// StringNull valid if is null
func StringNull(data string) sql.NullString {
	null := sql.NullString{String: data}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// IntNull valid if is null
func IntNull(data int32) sql.NullInt32 {
	null := sql.NullInt32{Int32: data}
	if null.Int32 != 0 {
		null.Valid = true
	}
	return null
}
