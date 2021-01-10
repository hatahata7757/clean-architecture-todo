FROM golang:1.15.6-alpine3.12

ENV GO111MODULE=on
WORKDIR /go/src/api/

RUN apk update \
  && apk add --no-cache git bash \
  && go mod init \
  && go get -u github.com/gin-gonic/gin \
  && go get -u github.com/jinzhu/gorm \
  # for remote debug
  && go get -u github.com/go-delve/delve/cmd/dlv \
  # for hot reload
  && go get -u github.com/cosmtrek/air
CMD air
