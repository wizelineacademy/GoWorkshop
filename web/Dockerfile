FROM golang:1.9

ENV root_dir /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $root_dir/proto
ADD web $root_dir/web
WORKDIR $root_dir/web

RUN go get github.com/gocraft/web
RUN go get github.com/gorilla/context
RUN go get google.golang.org/grpc
RUN go get golang.org/x/net/context

#RUN curl https://glide.sh/get | sh
#RUN glide install
RUN go install

ENTRYPOINT /go/bin/web
EXPOSE 8080
