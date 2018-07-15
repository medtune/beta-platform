FROM golang:1.10

WORKDIR /go/src/github.com/medtune/beta-platform

COPY . .

RUN go get -v ./...

RUN go build -v -o ./medtune-beta ./cmd/main.go

RUN ls -la

EXPOSE 8005

ENTRYPOINT [ "./medtune-beta", "start" ]

