
FROM golang:tip-alpine3.22

WORKDIR /go/src

COPY . .

RUN go mod download

RUN go build -o main ./main.go

CMD ["./main"]