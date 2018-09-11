#Base image
FROM medtune/capsul:latest

# FROM medtune/capsul:dev-latest

WORKDIR /go/src/github.com/medtune/beta-platform

ADD build build
ADD cmd cmd
ADD hack hack
ADD pkg pkg
ADD server server
ADD static static
ADD test test
ADD config.yml .
ADD Makefile .
ADD CXPBA.xlsx .

ADD MAINTAINERS.txt .
ADD LICENSE.txt .

ADD .codecov.yml .
ADD .circleci/config.yml .
ADD .git/ .

RUN bash ./hack/install-deps.sh --hack
