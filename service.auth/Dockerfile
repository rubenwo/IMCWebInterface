FROM golang:latest

WORKDIR /go/src/imc-web-interface
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./service.auth

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/bin/service.auth .
COPY --from=0 /go/src/imc-web-interface/service.auth/accounts.json .

CMD ["./service.auth"]