FROM golang:1.7-alpine

ADD . /home 
WORKDIR /home

RUN go get -u ./...
RUN go build main.go

CMD ["./main"]