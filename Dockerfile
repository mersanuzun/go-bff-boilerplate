FROM golang:1.13.4-alpine3.10

WORKDIR /go/src/github.com/mersanuzun/go-bff-boilerplate

ENV GO111MODULE=on

RUN apk update && apk add git curl gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

RUN go get -u github.com/swaggo/swag/cmd/swag

COPY . .

RUN swag init

RUN go build -o main .

CMD ["./main"]
