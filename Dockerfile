FROM golang:1.23 as autstore

RUN apt-get update && apt-get install -y curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
# Установка curl 
RUN go mod init autstor && go build 

CMD ["./autstor"]