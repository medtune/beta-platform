FROM medtune/beta-platform:build

ARG CODECOV_TOKEN

RUN bash ./hack/go.test-race.sh

RUN curl -s https://codecov.io/bash > codecov.sh

RUN bash codecov.sh -t $CODECOV_TOKEN