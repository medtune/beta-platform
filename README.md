# MedTune Beta Platform

## Build status

Circle CI [configuration](./.circleci/config.yaml)

| branch | status |
| --- | --- | 
| Iron | circleci |
| Iron-dev | circleci |

## Table of content

- [1. Build status](#medtune-beta-platform)

- [2. Table of content](#table-of-content)

- [3. Overview](#overview)

- [4. Concepts](#concepts)
    - [1. Capsules]()
    - [2. Demos]()
    - [3. Slides]()

- [5. Getting Started](#getting-started)
   - [1. Pre-requirements](#1-pre-requirements)
   - [2. Configuration](#2-configuration)
   - [3. Run server](#3-run-server)
      - [1. Using Golang compiler](#using-golang-compiler)
      - [2. Using go get](#using-go-get-goguette)
      - [3. Using Docker engine](#using-docker-engine)
      - [4. Using Docker hub](#using-docker-hub)
      - [5. Using Kubernetes](#using-kubernetes)

- [6. Command line specs](#command-line)

- [7. Change log](#change-log)

- [8. Contributing](#contributing)

- [9. Maintainers](#maintainers)

- [10. License](#License)


### Overview

MedTune Beta platform is a proof of concept for building robust and scalable AI solutions served into one user interface.

### Concepts

- Medtune Platform
- Capsules
- Architecture
- Service Discovery


### Getting started

##### 1 - Pre-Requirements

Some of MedTune demos might need to setup a set of local Capsules. You can use [run_capsules.sh]() for default setup.

To run Capsules we recommand container engine like [Docker](https://github.com/moby/moby) or [cri-o](https://github.com/kubernetes/cri-o), please refer to [medtune/capsules](https://github.com/medtune/capsules) setup instructions for more informations.

You can assert that needed Capsules are running, by adding the flag `--check-capsules=true` to medtune-beta command line.

To gain access to all of our demos, please make sure the following Capsules are running and well configured in [config.yaml](config.yaml):
- inception
- mnist

##### 2 - Configuration

All configurations needed are well commented in [config.yaml](config.yaml)

##### 3 - Web application

Once capsules are set (or not), you can build MedTune Beta Platform and use it localy:

###### Using Golang compiler

```shell
# Set your $GOPATH
# Make sure your working directory is $GOPATH/src/github.com/medtune

# Clone project
git clone https://github.com/medtune/beta-platform && cd beta-platform

# Compile for current OS/ARCH 
# Use GOOS=linux GOARCH=amd if you are compiling for other platforms
go build -o medtune-beta ./cmd/medtune-beta/main.go

# Run server
medtune-beta run --port=8080 --sync-db --check-capsules
```

###### Using go get (GOGUETTE)

```shell
# Install package
go get -u github.com/medtune/beta-platform/

# Run server
medtune-beta run --port=8080 --sync-db --check-capsules
```

###### Using Docker engine

```shell
# Build image
docker build -t . medtune-beta:iron

# Run container
docker run --name=beta-platform -p 8080:8007 medtune/beta-platform:latest
```


###### Using Docker hub

```shell
# Pull image
docker pull medtune/beta-platform:latest

# Run container
docker run --name=beta-platform -p 8080:8007 medtune/beta-platform:latest
```


###### Using kubernetes

```shell
# Create deployment
kubectl create -f deploy/k8s-deployment.yaml
# Expose service
kubectl create -f deploy/k8s-servie.yaml 
```

### Command line


Main flags

| flags | Description |
| --- | --- |
| `run` | Run server using local config.yaml  |
| `debug` | Debug server (Front-end development)|
| `sync-db` | Sync database models |
| `check-capsules` | Checks if configurated capsules are running |


Subflags
| sub-flags | Description |
| --- | --- |
| `--config -c file` | Run server using different configuration file (default is `config.yaml`)|
| | 


### Change log

- 0.1.0 - Iron

### Contributing

For contributions please refer to our friendly code of conduct in [how-we-work](https://github.com/medtune/how-we-work) repository :smile:.

### Maintainers

See [MAINTAINERS](MAINTAINERS)

### License

Medtune Beta Platform is released under the Apache 2.0 license. See [LICENSE.txt](License).
