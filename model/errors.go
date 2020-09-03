package model

import (
	"errors"
)

var (
	// ErrorRowAffected .
	ErrorRowAffected = errors.New("Se afecto mas de una fila")
)
