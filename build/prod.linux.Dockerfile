#Base image
FROM alpine:latest

WORKDIR /mdtn

ADD medtune-beta .
ADD static static
ADD config.yml config.yml
ADD CXPBA.xlsx CXPBA.xlsx

RUN chmod +x medtune-beta

RUN ls -la
RUN ./medtune-beta version

# Port to expose
EXPOSE 8005

ENTRYPOINT [ "./medtune-beta", \
    "start", \
    "--syncdb", \
    "--sync-cxpba", \
    "--wait", \
    "--create-users"]
