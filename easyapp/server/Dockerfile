FROM golang:1.14.6
WORKDIR /go/src/server
ADD ./ /go/src/server
RUN go get -d && go get golang.org/x/tools/cmd/godoc