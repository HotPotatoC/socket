package socket

import (
	"io"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func (s *Socket) contextBuilder(a *Actor) (*Context, error) {
	header, err := ws.ReadHeader(*a.conn)
	if err != nil {

	}
	payload := make([]byte, header.Length)
	_, err = io.ReadFull(*a.conn, payload)
	if err != nil {
		// handle error
	}
	data, _, err := wsutil.ReadClientData(*a.conn)

	ctx := createContext(s.config)
	ctx.event.header = &header
	ctx.message.data = &data
	ctx.sender = *a
	return ctx, nil
}
