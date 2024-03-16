FROM golang:1.22.1-alpine3.18 AS builder

WORKDIR /app

COPY go.* main.go ./
RUN go mod download

RUN go build -o app .


FROM alpine:3.18
RUN apk --update --no-cache add ca-certificates
WORKDIR /app/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
