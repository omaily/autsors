FROM golang:1.23 as autstore

ENV CGO_ENABLED 0
ENV GOOS linux

# RUN apt-get update && apt-get install -y curl
WORKDIR /app

COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o autstore .

CMD [".autstore"]