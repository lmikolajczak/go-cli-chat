FROM golang:1.17

WORKDIR /code
COPY . .

CMD tail -f /dev/null
