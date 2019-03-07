FROM golang:latest

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./client.web

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/bin/client.web .
COPY --from=0 /go/src/imc-web-interface/client.web/public/ ./public/
CMD ["./client.web"]