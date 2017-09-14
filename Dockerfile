# This is a template Dockerfile for all services
FROM znly/protoc

# Docker Compose sets this arg
ARG service_name
ENV proto_path /go/src/github.com/wizelineacademy/GoWorkshop/proto

# Install dependencies
RUN apk add --no-cache libc-dev curl git go
RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context
RUN go get gopkg.in/mgo.v2

ADD proto $proto_path
ADD shared /go/src/github.com/wizelineacademy/GoWorkshop/shared
ADD $service_name /go/src/github.com/wizelineacademy/GoWorkshop/$service_name
WORKDIR /go/src/github.com/wizelineacademy/GoWorkshop/$service_name

# Build source code for the gRPC client and server interfaces
RUN protoc -I $proto_path --go_out=plugins=grpc:$proto_path $proto_path/users/service.proto
RUN protoc -I $proto_path --go_out=plugins=grpc:$proto_path $proto_path/list/service.proto

#RUN mkdir -p /go/bin
#ENV PATH "/go/bin:${PATH}"
#RUN curl https://glide.sh/get | sh
#RUN glide install
RUN go install
EXPOSE 8080
