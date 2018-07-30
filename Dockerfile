FROM alpine:3.7

WORKDIR /medtune/beta-platform

# Copy binaries
COPY medtune-beta .

# Copy configuration
COPY config/config.yml .

# Copy static files
COPY static ./static

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", "start", "--syncdb", "--wait", "--create-users" ]

