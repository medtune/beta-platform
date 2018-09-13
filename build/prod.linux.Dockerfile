#Base image
FROM alpine:latest

WORKDIR /medtune

ADD medtune-beta .

ADD static static
ADD config.yml config.yml
ADD CXPBA.xlsx CXPBA.xlsx

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", \
    "start", \
    "--syncdb", \
    "--sync-cxpba", \
    "--wait", \
    "--create-users"]
