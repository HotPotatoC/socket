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
	ctx.message.data = &payload
	ctx.event.h = &h
	ctx.event.code = &code
	ctx.sender = a
	return ctx, nil
}

func frameBuilderAndSender(a *Actor, code interface{}, data []byte, status ...ws.StatusCode) (err error) {
	var ok bool
	var val TypeCode
	var frame ws.Frame

	if val, ok = parseTypeCode(code); ok {
		if ok, err = val.Eq(TypeDisconnected); ok {
			if len(status) == 1 {
				frame = ws.NewCloseFrame(ws.NewCloseFrameBody(status[0], string(data)))
			}
		} else {
			frame = ws.NewFrame(ws.OpCode(val), true, data)
		}
		err = ws.WriteFrame(*a.conn, frame)
	} else {
		err = typeNotSupported
	}
	return
}
