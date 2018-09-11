FROM medtune/beta-platform:build

# Port to expose
EXPOSE 8005

ADD config.k8s.yml .

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta",\
    "start",\
    "-f", "config.k8s.yml",\
    "--syncdb",\
    "--sync-cxpba",\
    "--wait",\
    "--create-users"]

