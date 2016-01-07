FROM ubuntu:14.04
MAINTAINER Arnold Cano <arnoldcano@yahoo.com>
RUN apt-get update
RUN apt-get install -y golang-go
RUN apt-get install -y ruby
RUN apt-get install -y python 
