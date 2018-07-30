FROM golang:1.10

WORKDIR /go/src/github.com/medtune/beta-platform

# Copy Go package
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./server ./server

# Copy static files
COPY ./static ./static

# Copy config
COPY config.yml .

# Download dependencies
RUN go get -v ./...

# Build binaries
RUN go build -v -o ./medtune-beta ./cmd/main.go

