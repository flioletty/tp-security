FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o sec-app ./cmd/app/main.go

CMD ["./sec-app"]