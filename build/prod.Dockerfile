FROM medtune/beta-platform:build

RUN rm -rf vendor cmd pkg server

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", \
    "start", \
    "--syncdb", \
    "--sync-cxpba", \
    "--wait", \
    "--create-users" ]

