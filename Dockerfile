FROM golang:1.9

# Docker Compose sets this arg
ARG service_name
ENV root_dir /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $root_dir/proto
ADD shared $root_dir/shared
ADD $service_name $root_dir/$service_name
WORKDIR $root_dir/$service_name

RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context
RUN go get gopkg.in/mgo.v2
RUN go get github.com/golang/protobuf/proto

#RUN curl https://glide.sh/get | sh
#RUN glide install
RUN go install
EXPOSE 8080
