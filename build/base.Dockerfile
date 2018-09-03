#Base image
FROM medtune/capsul:latest

# FROM medtune/capsul:dev-latest

WORKDIR /go/src/github.com/medtune/beta-platform

ADD . .

RUN bash ./hack/install-deps.sh --hack

