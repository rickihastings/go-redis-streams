FROM golang:latest AS builder

ENV GOPATH=/go
ENV HOST=redis:6379

ADD . /go/src/github.com/rickihastings/go-redis-streams/example/producer
WORKDIR /go/src/github.com/rickihastings/go-redis-streams/example/producer

RUN set -x && go get all ./...

RUN ls -la
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main main.go

FROM scratch

COPY --from=builder /main .

CMD ["/main"]
