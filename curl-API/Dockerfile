FROM golang:1.17.2-buster

WORKDIR /go/src/github.com/guranytou/curl-API

COPY go.mod main.go ./
RUN go mod download
EXPOSE 8000

CMD ["go", "run", "main.go"]