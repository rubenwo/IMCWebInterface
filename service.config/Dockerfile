FROM golang:latest

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./service.config

RUN chmod +x service.config
CMD ["./service.config"]