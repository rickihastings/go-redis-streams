FROM golang:latest AS builder

ENV GOPATH=/go

ADD . /go/src/github.com/rickihastings/go-redis-streams
WORKDIR /go/src/github.com/rickihastings/go-redis-streams/example/consumer

RUN set -x && go get all ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main main.go

FROM scratch

COPY --from=builder /main .
