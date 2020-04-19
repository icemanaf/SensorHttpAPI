FROM golang:alpine AS builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -o server

FROM arm32v7/alpine as final

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]

