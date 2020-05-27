FROM golang:1.14
WORKDIR /src
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y postgresql-client
RUN wget https://github.com/cortesi/modd/releases/download/v0.8/modd-0.8-linux64.tgz
RUN tar zxvf modd-0.8-linux64.tgz
RUN cp modd-0.8-linux64/modd /usr/local/bin/modd
RUN rm -r modd-0.8-linux64 && rm modd-0.8-linux64.tgz
