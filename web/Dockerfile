FROM golang:1.9

ENV PKG_PATH /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $PKG_PATH/proto
ADD web $PKG_PATH/web
WORKDIR $PKG_PATH/web

RUN go get github.com/gocraft/web
RUN go get github.com/gorilla/context
RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context

RUN go install

ENTRYPOINT /go/bin/web
EXPOSE 8080
