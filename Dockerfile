FROM golang:1.9

# users|list|notifier
ARG SRV_NAME

ENV PKG_PATH /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $PKG_PATH/proto
ADD shared $PKG_PATH/shared
ADD $SRV_NAME $PKG_PATH/$SRV_NAME
WORKDIR $PKG_PATH/$SRV_NAME

RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context
RUN go get gopkg.in/mgo.v2
RUN go get github.com/golang/protobuf/proto

RUN go install
EXPOSE 8080
