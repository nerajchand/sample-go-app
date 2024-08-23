FROM golang:1.23

RUN apt-get update && \
    apt-get install -y git

WORKDIR /go/src/app

COPY . .

# Initialize Go module and download dependencies
RUN go mod init app || true && \
    go mod tidy

RUN go get github.com/kataras/iris/v12@latest

# https://golang.org/cmd/cgo/
# Build the Go application
RUN CGO_ENABLED=0 go build -o main main.go

CMD ["./main"]
