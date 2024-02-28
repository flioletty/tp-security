FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o sec-proxy ./cmd/proxy/main.go

CMD ["./sec-proxy"]