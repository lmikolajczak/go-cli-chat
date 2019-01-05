GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install

build: 
	$(GOBUILD) -o ./cmd/chat-server/chat-server ./cmd/chat-server/
	$(GOBUILD) -o ./cmd/chat-client/chat-client ./cmd/chat-client/

install:
	$(GOINSTALL) ./...

run-server:
	$(GOBUILD) -o ./cmd/chat-server/chat-server ./cmd/chat-server/
	./cmd/chat-server/chat-server

run-client:
	$(GOBUILD) -o ./cmd/chat-client/chat-client ./cmd/chat-client/
	./cmd/chat-client/chat-client
