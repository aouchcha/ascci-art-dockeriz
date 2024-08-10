FROM  golang:1.22-alpine AS builder

WORKDIR /ascii-art-web-docker

RUN apk update

RUN apk add bash

RUN bash

COPY . .

RUN go build -o hello main.go 

EXPOSE 8082

CMD  ["./hello"]