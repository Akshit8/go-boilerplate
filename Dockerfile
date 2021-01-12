FROM golang:1.15.6-alpine as builder
COPY go.mod go.sum /go/src/github.com/Akshit8/go-boilerplate/
WORKDIR /go/src/github.com/Akshit8/go-boilerplate
RUN go mod download
COPY . /go/src/github.com/Akshit8/go-boilerplate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-boilerplate github.com/Akshit8/go-boilerplate/cmd

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/Akshit8/go-boilerplate/build/go-boilerplate /usr/bin/go-boilerplate
EXPOSE 8000
ENTRYPOINT ["/usr/bin/go-boilerplate"]