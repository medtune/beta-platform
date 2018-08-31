#Base image
FROM medtune/capsul:v0.0.2

# FROM medtune/capsul:dev-latest

WORKDIR /go/src/github.com/medtune/beta-platform

ADD . .

RUN bash ./hack/install-deps.sh --hack

