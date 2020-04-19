FROM golang:alpine AS builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o server

FROM alpine:latest as final

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]

