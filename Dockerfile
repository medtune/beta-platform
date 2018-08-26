#Base image
FROM medtune/capsul:dev-v0.1.0

WORKDIR /go/src/github.com/medtune/beta-platform

ADD . .

RUN bash ./hack/install-deps.sh --hack

RUN go build -o medtune-beta cmd/main.go
