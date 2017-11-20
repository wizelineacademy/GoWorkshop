FROM golang:1.9-alpine

ENV PKG_PATH /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $PKG_PATH/proto
ADD web $PKG_PATH/web
WORKDIR $PKG_PATH/web
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get github.com/gocraft/web github.com/gorilla/context google.golang.org/grpc golang.org/x/net/context github.com/sirupsen/logrus

RUN go install

ENTRYPOINT /go/bin/web
EXPOSE 8080
