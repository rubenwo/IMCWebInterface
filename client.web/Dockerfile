FROM golang:latest

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./client.web

RUN chmod +x client.web
CMD ["./client.web"]