FROM golang:1.12.7-alpine3.10
ENV CGO_ENABLED 0

RUN apk update && apk add bash inotify-tools git

RUN mkdir /build/
WORKDIR /build/

# installing Delve debugger
RUN go get github.com/derekparker/delve/cmd/dlv

# build micro-service binary
COPY main.go /build/main.go
COPY startScript.sh /build/startScript.sh

ENTRYPOINT sh startScript.sh