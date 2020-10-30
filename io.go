package socket

import (
	"io"

	"github.com/gobwas/ws"
)

func (s *Socket) contextBuilder(a *Actor) (*Context, error) {
	h, err := ws.ReadHeader(*a.conn)
	if err != nil {
		return nil, err
	}

	// payload := make([]byte, 1024)
	payload := make([]byte, h.Length)
	// reader := io.LimitReader(*a.conn, h.Length)
	// _, err = reader.Read(payload)
	_, err = io.ReadFull(*a.conn, payload)
	if err != nil {
		return nil, err
	}
	if h.Masked {
		ws.Cipher(payload, h.Mask, 0)
	}
	ctx := createContext(s.config)
	ctx.event = Event{
		header: &h,
	}
	ctx.message = Message{
		data: &payload,
	}
	ctx.sender = *a
	return ctx, nil
}
