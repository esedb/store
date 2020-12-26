FROM golang

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go get install -v ./...

CMD ["app"]