FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o colaborator .

EXPOSE 8080

# Service a
CMD ["./colaborator"]