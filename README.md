# MedTune Platform ~ Beta v0.1.4

[![GitHub license](https://img.shields.io/github/license/medtune/beta-platform.svg)](https://github.com/medtune/beta-platform/blob/iron-master/LICENSE.txt) [![GitHub pull-requests](https://img.shields.io/github/issues-pr/medtune/beta-platform.svg)](https://GitHub.com/medtune/beta-platform/pull/) [![GitHub issues](https://img.shields.io/github/issues/medtune/beta-platform.svg)](https://github.com/medtune/beta-platform/issues) [![GitHub release](https://img.shields.io/github/release/medtune/beta-platform.svg)](https://GitHub.com/medtune/beta-platform/releases) 

## Build Status & Coverage

| branch | Status | C.C | C.Q |
| --- | --- | --- | --- |
| Master | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/master.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/master) | [![CodeFactor](https://www.codefactor.io/repository/github/medtune/beta-platform/badge)](https://www.codefactor.io/repository/github/medtune/beta-platform) | [![codecov](https://codecov.io/gh/medtune/beta-platform/branch/master/graph/badge.svg)](https://codecov.io/gh/medtune/beta-platform) |
| Dev | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/dev.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/dev) | [![CodeFactor](https://www.codefactor.io/repository/github/medtune/beta-platform/badge/dev)](https://www.codefactor.io/repository/github/medtune/beta-platform/overview/dev) | |

***

## Table of content

- [1. Build status](#medtune-beta-platform)
- [2. Table of content](#table-of-content)
- [3. Overview](#overview)
- [4. Concepts](#concepts)
    - [1. Platform](#platform)
    - [2. Capsules](#capsules)
    - [3. Architecture](#architecture)
    - [4. Service Discovery](#service-discovery)
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
- [10. License](#License)

### Overview

MedTune Beta platform is a _proof of concept_ prototype, focused on serving scalable AI solutions, into different types of customers.

### Concepts

**TODO ASAP**

### Getting started

##### 1 - Prerequisites

To use medtune-beta command line you need a Golang Compiler (Tested on 1.10 and 1.9) 

To run medtune-beta platform services localy you will need a container engine/orchestrator (like Docker swarm, Kubernetes).

###### Database
Medtune Beta runs next to postgres database (if you want to go beyond signup).

###### Capsules

Some of MedTune demos might need to setup some Capsules. [Docker](https://github.com/moby/moby) or [cri-o](https://github.com/kubernetes/cri-o) are enought to make things work, but it requires a lot of handy work. We recommand to use [Swarm](#using-swarm) or [Kubernetes](#kubernetes) for easier setup.


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

# Compile for current OS/ARCH 
# Use GOOS=linux GOARCH=amd if you are compiling for other platforms
go build -o medtune-beta ./cmd/medtune-beta/main.go

# Run server
medtune-beta start
```

###### Using go get (GOGUETTE)

```shell
# Install package
go get -u github.com/medtune/beta-platform/

# Run server
medtune-beta run --port=8005 --static=$GOPATH/src/github.com/medtune/beta-platform/static
```

###### Using Docker engine

```shell
# Build image
docker build -t . medtune-beta:iron

# Run container
docker run --name=beta-platform -p 8005:8005 medtune/beta-platform:latest
```


###### Using Docker hub

```shell
# Pull image
docker pull medtune/beta-platform:latest

# Run container
docker run --name=beta-platform -p 8005:8005 medtune/beta-platform:latest
```


###### Using kubernetes

```shell
# Create deployment
kubectl create -f deploy/kubernetes.yaml
```

###### Using swarm

```shell
# Create docker swarm
docker-compose up -d
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
