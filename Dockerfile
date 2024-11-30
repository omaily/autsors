FROM golang:latest as autstore

ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GOOS linux

RUN apt-get update && apt-get install -y curl 
RUN apk bash git gcc musl-dev
WORKDIR /app

COPY ./ ./

RUN go mod download
RUN sh -s -- -b $(go env GOPATH)/bin
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main .

CMD ["./crutch.sh"] 
