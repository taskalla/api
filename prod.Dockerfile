FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build -o bin/taskalla ./cmd/taskalla

CMD ["./bin/taskalla"]
