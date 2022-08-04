[![Go Report Card](https://goreportcard.com/badge/github.com/Luqqk/go-cli-chat)](https://goreportcard.com/report/github.com/Luqqk/go-cli-chat)

## 💬 go-cli-chat

Chat server and client written in Go (simple prototype for learning purposes). The application utilizes goroutines and channels.

![![asciicast](https://asciinema.org/a/512757.svg)](https://asciinema.org/a/512757)

### Usage

```bash
# Build and start container:
docker compose up -d
# Enter go-cli-chat container:
docker exec -it go-cli-chat bash
# Run chat server within go-cli-chat container:
go-cli-chat-server
# Open another go-cli-chat container and start client:
go-cli-chat-client -username Luqqk
```

You can also make changes and rebuild either `client` or `server` by using:

```bash
$ make build-server
```

```bash
$ make build-client
```

### Contributing

I am open to, and grateful for, any contributions made by the community.
