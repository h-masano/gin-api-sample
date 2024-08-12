FROM golang:1.22-alpine
RUN mkdir /go/src/app/
WORKDIR /go/src/app/
COPY ./ /go/src/app/
