FROM ubuntu:14.04

MAINTAINER Arnold Cano <arnoldcano@yahoo.com>

RUN apt-get update
RUN apt-get install -y golang
RUN apt-get install -y ruby
RUN apt-get install -y python
RUN apt-get install -y nodejs

ENV GOPATH /go
WORKDIR /go/src/github.com/arnoldcano/usul
ADD . /go/src/github.com/arnoldcano/usul/
RUN go build

ENTRYPOINT ["go", "run", "main.go"]

EXPOSE 8080
