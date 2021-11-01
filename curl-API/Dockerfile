FROM golang:1.17.2-buster

WORKDIR /go/src/github.com/guranytou/curl-API

COPY . .
RUN go mod download

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]