# MedTune Platform ~ Beta v0.1.6

/!\ This repository is no longer maintained. MedTune official platform under development

[![GitHub license](https://img.shields.io/github/license/medtune/beta-platform.svg)](https://github.com/medtune/beta-platform/blob/iron-master/LICENSE.txt) [![GitHub pull-requests](https://img.shields.io/github/issues-pr/medtune/beta-platform.svg)](https://GitHub.com/medtune/beta-platform/pull/) [![GitHub issues](https://img.shields.io/github/issues/medtune/beta-platform.svg)](https://github.com/medtune/beta-platform/issues) [![GitHub release](https://img.shields.io/github/release/medtune/beta-platform.svg)](https://GitHub.com/medtune/beta-platform/releases) 

## Build Status & Coverage

| branch | Status | Quality | Coverage |
| --- | --- | --- | --- |
| Master | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/master.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/master) | [![CodeFactor](https://www.codefactor.io/repository/github/medtune/beta-platform/badge)](https://www.codefactor.io/repository/github/medtune/beta-platform) | [![codecov](https://codecov.io/gh/medtune/beta-platform/branch/master/graph/badge.svg)](https://codecov.io/gh/medtune/beta-platform) |
| Dev | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/dev.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/dev) | [![CodeFactor](https://www.codefactor.io/repository/github/medtune/beta-platform/badge/dev)](https://www.codefactor.io/repository/github/medtune/beta-platform/overview/dev) | |

***

## Table of content

- [1. Build status](#medtune-beta-platform)
- [2. Table of content](#table-of-content)
- [3. Overview](#overview)
- [4. Concepts](#concepts)
- [5. Getting started](#getting-started)
   - [1. Prerequisites](#1-prerequisites)
   - [2. Configuration](#2-configuration)
   - [3. Deploy application](#3-deploy)
      - [1. Using Golang compiler](#using-golang-compiler)
      - [2. Using go get](#using-go-get-goguette)
      - [3. Using Docker engine](#using-docker-engine)
      - [4. Using Docker hub](#using-docker-hub)
      - [5. Using Kubernetes](#using-kubernetes)
      - [6. Using Swarm](#using-swarm)
- [6. Command line: _medtune-beta_](#command-line)
   - [1. Root](./cmd/README.md#root-command)
   - [2. Version](./cmd/README.md#version)
   - [3. Automigrate](./cmd/README.md#automigrate)
   - [4. Generate views](./cmd/README.md#generate-views)
   - [5. Debug server](./cmd/README.md#debug-server)
   - [6. Run server](./cmd/README.md#run-server)
- [7. Change log](#changelog)
- [8. Contributing](#contributing)
- [9. Maintainers](#maintainers)
- [10. License](#license)
- [11. Todo](#todo)

### Overview

MedTune Beta platform is a _proof of concept_ prototype, focused on serving scalable AI solutions, into different types of customers.

### Concepts

**TODO ASAP**

### Getting started

##### 1 - Prerequisites

You can compile the binaries your self or use our containers images.

To compile medtune-beta command line you need a Golang Compiler (Tested on 1.11, 1.10 and 1.9)

To run medtune-beta platform side services you will need a container engine/orchestrator (like Docker, Kubernetes).

###### Database

Medtune Beta runs next to postgres database (if you want to go beyond signup).

###### Cache

The application cache users sessions on a Redis cache database.

###### Capsules

Some of MedTune demos might need to setup their correspendent Capsules (server that perform inference computations). [Docker](https://github.com/moby/moby) or [cri-o](https://github.com/kubernetes/cri-o) are enought to make things work, but it requires a lot of handy work. We recommand to use [Docker compose](#using-docker-compose) or [Kubernetes](#kubernetes) for easier setup.


##### 2 - Configuration

All configurations needed are well commented in [config.yaml](config.yaml)

##### 3 - Deploy

Once capsules are set (or not), you can build MedTune Beta Platform and use it localy:

###### Using Golang compiler

```shell
# Set your $GOPATH
# Make sure your working directory is $GOPATH/src/github.com/medtune

# Clone project
git clone https://github.com/medtune/beta-platform && cd beta-platform

# Fast run
make run

# Compile for current OS/ARCH 
# Use GOOS=X GOARCH=Y if you are compiling for other platforms
make release
# make release-cmd
# make release-dev
# make release-debug
# make release-linux

# Run server
medtune-beta start
```

###### Using Docker engine

```shell

# build alpine based image
make build-alpine # medtune/beta-platform:prod-alpine

# build alpine based image for k8s cluster
make build-k8s # medtune/beta-platform:prod-k8s

# build all images
make build-all

# Run container
docker run \
   --name=beta-platform \
   -p 8005:8005 \
   medtune/beta-platform:latest
```

###### Using go command

```shell
# Install package
go get -u github.com/medtune/beta-platform/...

# Run server
export $STATIC_PATH=/dir/files/static
medtune-beta run --port=8005 --static=$STATIC_PATH
```


###### Using Docker hub

```shell
# Pull image
docker pull medtune/beta-platform:TAG # TAG=prod | prod-alpine | prod-k8s | latest

# Run container
docker run --name=beta-platform -p 8005:8005 medtune/beta-platform:TAG
```

###### Using Docker compose

```shell
# Create docker swarm
docker-compose up -d
```

###### Using kubernetes

```shell
# Create deployment
kubectl create -f deployments/kubernetes/

# Generate kubernetes configuration from docker-compose.yaml
make gen-k8s
```

### Development utils

###### CICD

See `.circleci/config.yaml` for more informations about CI/CD setup

###### Testing

```shell
# Unit tests
make tests

# racing & coverage tests
make test-cov
```

###### Generating kubernetes deployment yaml files

```shell
# Generate kubernetes configuration from docker-compose.yaml
make gen-k8s
```

###### Cleaning demo files

```shell
# Clean all demos static images
make clean-demos
```

###### Scene debuging (capsules)

```shell
# Makesure your docker engine is connected to medtune hub
make start-scene

# Stop all capsules containers
make stop-scene

# Kill all capsules containers
make kill-scene
```

### Command line

 You can find medtune-beta [command line docs](cmd/README.md) at `./cmd`

### Changelog

See [CHANGELOG.md](CHANGELOG.md)

### Contributing

For contributions please refer to our friendly code of conduct in [how-we-work](https://github.com/medtune/how-we-work) repository :smile:.

### Maintainers

See [MAINTAINERS](MAINTAINERS.txt)

### License

Medtune Beta Platform is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt).

#### TODO

- End2End Tests
- pprof