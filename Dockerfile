FROM golang:1.12

ADD . /go/src/test_service_timofeev

RUN go install github.com/Sane4ek55/grpc_test

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300