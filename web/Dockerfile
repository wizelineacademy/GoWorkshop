FROM golang:1.8

ADD proto /go/src/github.com/wizelineacademy/GoWorkshop/proto
ADD web /go/src/github.com/wizelineacademy/GoWorkshop/web
WORKDIR /go/src/github.com/wizelineacademy/GoWorkshop/web

RUN go get github.com/gocraft/web
RUN go get github.com/gorilla/context
RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context
RUN go install github.com/wizelineacademy/GoWorkshop/web

ENTRYPOINT /go/bin/web
EXPOSE 8080
