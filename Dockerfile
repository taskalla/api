FROM golang:1.15.7-alpine3.13

WORKDIR /usr/src/app

COPY . .

RUN go build -o bin/taskalla ./cmd/taskalla

FROM alpine:3.13

COPY --from=0 /usr/src/app/bin/taskalla /usr/bin/taskalla

CMD ["taskalla"]
