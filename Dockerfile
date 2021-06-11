FROM golang:1.16 AS builder

WORKDIR /telebot/

# go mod for deps
COPY go.mod go.sum /telebot/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./telebot

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /telebot/telebot /root/
CMD ["./telebot"]
