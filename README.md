## 💬 go-cli-chat

Chat server and client written in golang (simple prototype for learning purposes). The application heavily utilizes goroutines and channels. Go makes the concurrency quite easy to use and it is very cool to write such programs in the golang.

![chat-client](img/client.png)

### Usage

Clone this repository and just simply run:

Server (in go-cli-chat/src/server/):

```bash
$ go run server.go
```

Client (in go-cli-chat/src/client/):

```bash
$ go run client.go
```

You might want to specify host and port (in main func) on which server is listening and client is connecting to.

### TODO

* [client][server] code refactoring
