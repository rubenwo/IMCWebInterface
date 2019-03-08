FROM golang:latest

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./service.config

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/bin/service.config .
COPY --from=0 /go/src/imc-web-interface/service.config/configs.json .

CMD ["./service.config"]