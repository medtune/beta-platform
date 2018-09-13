FROM medtune/beta-platform:build

ARG CODECOV_TOKEN

RUN bash ./hack/go.test-race.sh

RUN mkdir /shared -p

ENTRYPOINT [ "mv", "coverage.txt", "/shared" ]

