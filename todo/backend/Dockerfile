FROM golang:1.17.6-alpine3.15

RUN apk update
WORKDIR /go/src/github/guranytou/todo

COPY ./backend/app .
RUN go mod download

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]