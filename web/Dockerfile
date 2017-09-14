FROM golang:1.8

ADD proto /go/src/github.com/wizelineacademy/GoWorkshop/proto
ADD web /go/src/github.com/wizelineacademy/GoWorkshop/web
WORKDIR /go/src/github.com/wizelineacademy/GoWorkshop/web

RUN curl https://glide.sh/get | sh
RUN glide install
RUN go install github.com/wizelineacademy/GoWorkshop/web

ENTRYPOINT /go/bin/web
EXPOSE 8080
