package socket

import (
	"io"

	"github.com/gobwas/ws"
)

func (s *Socket) contextBuilder(a *Actor) (*Context, error) {

	h, err := ws.ReadHeader(*a.conn)
	code := TypeCode(h.OpCode)

	if err != nil {
		return nil, err
	}

	// payload := make([]byte, 1024)
	// reader := io.LimitReader(*a.conn, h.Length)
	// _, err = reader.Read(payload)

	payload := make([]byte, h.Length)
	_, err = io.ReadFull(*a.conn, payload)

	if err != nil {
		return nil, err
	}

	if h.Masked {
		ws.Cipher(payload, h.Mask, 0)
	}

	ctx := createContext(s.config)
	ctx.message.data = payload
	ctx.event.code = code
	ctx.sender = a
	return ctx, nil
}

func frameCompiler(a *Actor, code TypeCode, data []byte) error {
	return ws.WriteFrame(*a.conn, ws.NewFrame(ws.OpCode(code), true, data))
}
