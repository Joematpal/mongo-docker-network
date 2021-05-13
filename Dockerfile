FROM golang:buster

ADD . /go/src/github.com/joematpal/mongo-docker-network
WORKDIR /go/src/github.com/joematpal/mongo-docker-network

RUN go install

ENTRYPOINT [ "mongo-docker-network" ]