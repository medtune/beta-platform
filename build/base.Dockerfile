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
ADD Makefile .

ADD MAINTAINERS.txt .
ADD LICENSE.txt .

ADD config.yml .
ADD CXPBA.xlsx .

ADD .codecov.yml .
ADD .circleci .circleci
ADD .git/ .git

RUN bash ./hack/install-deps.sh --hack
