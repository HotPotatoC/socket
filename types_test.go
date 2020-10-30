package socket_test

import (
	"testing"

	"github.com/Stalync/socket"
	"github.com/gobwas/ws"
)

func TestTypeCode_Eq(t *testing.T) {
	var etype = socket.TypeText
	if x, y := etype.Eq(ws.OpText); y != nil || !x {
		t.Error("Why this is wrong?, this shoud be true")
	}
}
