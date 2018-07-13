# beta-platform

| branch | status | 
| --- | --- | --- |
| Iron | circleci |
| Iron-dev | circleci |
| Master | circleci|

#### Run local

- Using Compile project

You need to set your $GOPATH.

```
# Clone project
git clone https://github.com/medtune/beta-platform && cd beta-platform

# Change GOOS & GOARCH to compile for other platforms
GOOS=linux GOARCH=amd go build -o medtune-beta ./cmd/medtune-beta/main.go

# Run server
medtune-beta run --port=8080 --sync-db --check-capsules
```

- With go get

```
go get -u github.com/medtune/beta-platform/
medtune-beta run --port=8080 --sync-db --check-capsules
```

- With Docker

```
# Pull image
docker pull medtune/beta-platform:latest

# Run container
docker run --name=beta-platform -p 8080:8007 medtune/beta-platform:latest
```

- Binaries

```


```

