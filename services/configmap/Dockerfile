FROM golang:1.15.1-alpine3.12 AS builder

WORKDIR $GOPATH/src/github.com/alextanhongpin/go-server

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -i -o app .

FROM alpine:3.12
RUN apk --update --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /go/src/github.com/alextanhongpin/go-server/app .

EXPOSE 8080

CMD ["./app"]
