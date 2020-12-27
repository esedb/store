FROM golang:latest

WORKDIR /go/src/estore

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install

WORKDIR /go/bin


EXPOSE 7000

# ENTRYPOINT go/bin/estore

CMD ["./estore"]