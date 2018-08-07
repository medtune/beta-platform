FROM medtune/capsul:dev-v0.1.0

WORKDIR /go/src/github.com/medtune/beta-platform

ADD . .

RUN go get -v ./pkg/config 
RUN go get -v ./pkg/session 
RUN go get -v ./pkg/store 
RUN go get -v ./pkg/jsonutil

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/spf13/cobra

RUN go build -tags=cicd -o medtune cmd/main.go

RUN ls -la
