#Base image
FROM medtune/beta-platform:base

RUN go build -o medtune-beta cmd/main.go
