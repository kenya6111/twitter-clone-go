
FROM golang:tip-alpine3.22

WORKDIR /go/src

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go mod download

RUN go build -o main ./main.go

CMD ["air", "-c", ".air.toml"]