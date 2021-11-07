FROM golang:1.17

WORKDIR /code
COPY . .

RUN go get -d ./...
RUN make build-server
RUN make build-client

CMD tail -f /dev/null
