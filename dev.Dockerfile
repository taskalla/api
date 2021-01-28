FROM golang:1.15.7-alpine3.13

WORKDIR /usr/src/app

COPY . .

RUN go get ./...

RUN go get -u github.com/cosmtrek/air
ENV air_wd /usr/src/app

EXPOSE 3000

CMD [ "air", "-c", ".air.toml" ]
