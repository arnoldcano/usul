FROM ubuntu:14.04

MAINTAINER Arnold Cano <arnoldcano@yahoo.com>

RUN apt-get update

RUN apt-get install -y golang
RUN apt-get install -y ruby
RUN apt-get install -y python
RUN apt-get install -y nodejs

RUN ln -s /usr/bin/nodejs /usr/bin/node

RUN apt-get install -y npm
RUN gem install rubocop
RUN apt-get install -y pylint
RUN npm install -g jshint

RUN useradd -m usul

ENV GOPATH /home/usul/go
WORKDIR /home/usul/go/src/github.com/arnoldcano/usul
COPY . /home/usul/go/src/github.com/arnoldcano/usul
RUN go build

CMD ["./usul"]

EXPOSE 8080
