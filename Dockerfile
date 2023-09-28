FROM golang:1.20

COPY . /go/src/app

WORKDIR /go/src/app/cmd

RUN go build -o gymshark-challenge main.go

EXPOSE 8080

CMD ["./gymshark-challenge"]