package socket

import (
	"errors"

	"github.com/gobwas/ws"
)

// TypeCode type alias for byte enum
type TypeCode byte

// List of every known code
const (
	// Event code
	TypeContinuation TypeCode = 0x0
	TypeConnected    TypeCode = 0x7
	TypeDisconnected TypeCode = 0x8
	TypePing         TypeCode = 0x9
	TypePong         TypeCode = 0xa

	// Message code
	TypeText   TypeCode = 0x1
	TypeBinary TypeCode = 0x2
)

// Eq supposed to compare code from this pkg
// to ws.OpCode
func (t *TypeCode) Eq(code interface{}) (bool, error) {
	var val TypeCode
	var ok bool
	var op ws.OpCode

	if op, ok = code.(ws.OpCode); ok {
		val = TypeCode(op)
	} else {
		val, ok = code.(TypeCode)
	}
	if ok {
		x := val ^ *t
		return x == 0, nil
	}
	return false, errors.New("Type not supported")
}
