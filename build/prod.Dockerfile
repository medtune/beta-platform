FROM medtune/beta-platform:build

# Port to expose
EXPOSE 8005

# please precise -v $CONFIGPATH:/medtune/beta-platform/deploy
ENTRYPOINT [ "./medtune-beta", "start", "--syncdb", "--wait", "--create-users" ]

