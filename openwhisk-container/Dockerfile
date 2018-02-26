FROM golang:1.10.0-alpine
WORKDIR /go/src/github.com/kelseyhightower/function
COPY . .
RUN go build -o function .

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
EXPOSE 8080
COPY --from=0 /go/src/github.com/kelseyhightower/function/function .
ENTRYPOINT ["/function"]
