FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o goradio

ENTRYPOINT ["./goradio"]