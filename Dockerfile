FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o hello-web-service

EXPOSE 8080

CMD ["./hello-web-service"]
