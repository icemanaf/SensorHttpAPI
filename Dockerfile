FROM golang:alpine AS builder

RUN apk add --no-cache gcc libc-dev

COPY . "/go/src/github.com/icemanaf/HttpConcepts"

WORKDIR "/go/src/github.com/icemanaf/HttpConcepts"

RUN go mod init


RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -o server

FROM arm32v7/alpine as final

COPY --from=builder /go/src/github.com/icemanaf/HttpConcepts/server .

EXPOSE 8080

CMD ["./server"]

