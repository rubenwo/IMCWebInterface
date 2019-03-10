FROM golang:latest AS builder

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./service.config

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/bin/service.config .
COPY --from=builder /go/src/imc-web-interface/service.config/configs.json .

CMD ["./service.config"]