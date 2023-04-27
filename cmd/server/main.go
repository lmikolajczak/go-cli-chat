package main

import (
	"github.com/lmikolajczak/go-cli-chat/pkg/chat"
	"github.com/lmikolajczak/go-cli-chat/pkg/transport"
)

func main() {
	transport.Serve(chat.NewSupervisor())
}
