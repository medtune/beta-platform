#Base image
FROM medtune/capsul:latest

# FROM medtune/capsul:dev-latest

WORKDIR /go/src/github.com/medtune/beta-platform

ADD build build
ADD cmd cmd
ADD hack hack
ADD internal internal
ADD static static
ADD test test
ADD Makefile .

ADD MAINTAINERS.txt .
ADD LICENSE.txt .

ADD etc/config.yml .
ADD assets/CXPBA.xlsx .

ADD .codecov.yml .
ADD .circleci .circleci
ADD .git/ .git

RUN bash ./hack/install-deps.sh --hack
