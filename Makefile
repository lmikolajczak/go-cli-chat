GOCMD=go
GOBUILD=$(GOCMD) build

build-server:
	cd cmd/server && $(GOBUILD) -o /go/bin/go-cli-chat-server

build-client:
	cd cmd/client && $(GOBUILD) -o /go/bin/go-cli-chat-client
