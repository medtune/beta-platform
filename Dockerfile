FROM medtune/capsul:dev-v0.1.0

WORKDIR /go/src/github.com/medtune/beta-platform

ADD . .

RUN go get -v ./pkg/config 
RUN go get -v ./pkg/session 
RUN go get -v ./pkg/store 
RUN go get -v ./pkg/jsonutil

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/spf13/cobra

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", "start", "--syncdb", "--wait", "--create-users" ]

