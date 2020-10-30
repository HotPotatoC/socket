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

	if x, y := etype.Eq(socket.TypeContinuation); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}

	if x, y := etype.Eq(socket.TypeConnected); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}

	if x, y := etype.Eq(socket.TypeClose); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}

	if x, y := etype.Eq(socket.TypePing); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}

	if x, y := etype.Eq(socket.TypePong); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}

	if x, y := etype.Eq(socket.TypeBinary); y != nil || x {
		t.Error("Why this is wrong?, this shoud be false")
	}
}
