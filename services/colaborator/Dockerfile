ARG BASE_IMAGE=golang:1.20.0-alpine3.16

FROM ${BASE_IMAGE} AS builder

WORKDIR /go/src

ADD . /go/src/

RUN go mod download && go mod verify

RUN go build -v -o /colaborator ./cmd/*.go

FROM ${BASE_IMAGE} AS candidate

ARG NAME=colaborator

# Install nice to haves
RUN apk add --no-cache openssl ncurses-libs libstdc++ libgcc curl libressl htop nano

WORKDIR /go/src/

COPY --from=builder /$NAME .

EXPOSE 8080

ENTRYPOINT ["./colaborator"]
