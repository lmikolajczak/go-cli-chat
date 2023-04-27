package transport

import (
	"errors"
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/lmikolajczak/go-cli-chat/pkg/chat"
)

func Serve(supervisor *chat.Supervisor) {
	// Use websocket.Server because we want to accept non-browser clients,
	// which do not send an Origin header. websocket.Handler does check
	// the Origin header by default.
	http.Handle("/", websocket.Server{
		Handler: supervisor.ServeWS(),
		// Set a Server.Handshake to nil - does not check the origin.
		// We can always provide a custom handshake method to access
		// the handshake http request and implement origin check or
		// other custom logic before the connection is established.
		Handshake: nil,
	})

	err := http.ListenAndServe(":3000", nil)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
