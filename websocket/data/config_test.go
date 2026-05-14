package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebSocketConstantsShouldRemainStable(t *testing.T) {
	require.Equal(t, "/save", WSRoute)
	require.Equal(t, "server", ModeServer)
	require.Equal(t, "client", ModeClient)
	require.Equal(t, int32(1), int32(AckMessage))
	require.Equal(t, int32(2), int32(PayloadMessage))
	require.Equal(t, "use of closed network connection", ClosedConnectionMessage)
}

func TestWebSocketErrorsShouldRemainStable(t *testing.T) {
	require.Equal(t, "nil marshaller", ErrNilMarshaller.Error())
	require.Equal(t, "nil payload processor provided", ErrNilPayloadProcessor.Error())
	require.Equal(t, "empty websocket url provided", ErrEmptyUrl.Error())
	require.Equal(t, "invalid web socket host mode", ErrInvalidWebSocketHostMode.Error())
	require.Equal(t, "acknowledge waiting timeout occurred", ErrAckTimeout.Error())
}
