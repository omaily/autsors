FROM golang:1.23 as autstore

WORKDIR /build

COPY main.go ./

# Установка curl 
RUN apt-get update && apt-get install -y curl
RUN go mod init autstor && go build 

CMD ["./autstor"]