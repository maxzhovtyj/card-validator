FROM golang:1.21.7-alpine3.19 AS builder

RUN go version
ENV GOPATH=/

COPY . /github.com/maxzhovtyj/card-validator/
WORKDIR /github.com/maxzhovtyj/card-validator/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/card-validator-linux-amd64 ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 github.com/maxzhovtyj/card-validator/.bin/card-validator-linux-amd64 .
COPY --from=0 github.com/maxzhovtyj/card-validator/config/config.yml config/

CMD ["./card-validator-linux-amd64", "-configPath=./config/config.yml"]