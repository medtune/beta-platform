#Base image
FROM alpine:latest

WORKDIR /medtune

ADD medtune-beta .

ADD static static
ADD config.yml config.yml

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", \
    "start", \
    "--syncdb", \
    "--wait", \
    "--create-users"]
