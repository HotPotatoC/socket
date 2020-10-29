package ws

// TypeCode type alias for byte enum
type TypeCode byte

// List of every known code
const (
	// Event code
	TypeContinuation TypeCode = 0x0
	TypeConnected    TypeCode = 0x7
	TypeClose        TypeCode = 0x8
	TypePing         TypeCode = 0x9
	TypePong         TypeCode = 0xa

	// Message code
	TypeText   TypeCode = 0x1
	TypeBinary TypeCode = 0x2
)
