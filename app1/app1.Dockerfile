FROM golang:1.23-alpine

WORKDIR /web

ADD . .

RUN go build -o main && chmod +x main

CMD ./main