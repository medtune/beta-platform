# MedTune Platform ~ Beta

[![GitHub release](https://img.shields.io/github/release/medtune/beta-platform.svg)](https://GitHub.com/Naereen/StrapDown.js/releases/) [![Percentage of issues still open](http://isitmaintained.com/badge/open/medtune/beta-platform.svg)](http://isitmaintained.com/project/Naereen/badges) [![GitHub license](https://img.shields.io/github/license/medtune/beta-platform.svg)](https://github.com/medtune/beta-platform/blob/master/LICENSE) [![GitHub pull-requests](https://img.shields.io/github/issues-pr/medtune/beta-platform.svg)](https://GitHub.com/Naereen/StrapDown.js/pull/)


## Build status

| branch | Build status |
| --- | --- | 
| Iron-Master | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/iron-master.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/iron-master) |
| Dev | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/dev-1.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/dev-1) |
| 0.0.3 | [![CircleCI](https://circleci.com/gh/medtune/beta-platform/tree/dev-1.svg?style=svg)](https://circleci.com/gh/medtune/beta-platform/tree/dev-1) |

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
   - [1. Root](#root-command)
   - [2. Version](#version)
   - [3. Automigrate](#automigrate)
   - [4. Generate views](#generate-views)
   - [5. Debug server](#debug-server)
   - [6. Run server](#run-server)
- [7. Change log](#changelog)
- [8. Contributing](#contributing)
- [9. Maintainers](#maintainers)
- [10. License](#License)


### Overview

MedTune Beta platform is a _proof of concept_ prototype, focused on serving scalable AI solutions, into different types of customers.

### Concepts

- Medtune Platform
- Capsules
- Architecture
- Service Discovery


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

#### Root command
```
Usage:
  medtune-beta [command]

Available Commands:
  automigrate Auto migrate database
  capsules    Not implemented
  debug       debug server for UI dev
  gen-views   Generate views html files
  help        Help about any command
  start       Run Medtune beta server
  version     Medtune beta actual version

Flags:
  -h, --help   help for medtune-beta
```

#### Subcommands

###### Version
```
Print Medtune Beta version

Usage:
  medtune-beta version [flags]

Flags:
  -h, --help   help for version
```

###### Auto-migrate
```
Sync database models by updating/creating existing
database tables

Usage:
  medtune-beta automigrate [flags]

Aliases:
  automigrate, syncdb

Flags:
  -f, --file string   Configuration file name (default "config.yml")
  -h, --help          help for automigrate
```

###### Generate views
```
Generate views html files

Usage:
  medtune-beta gen-views [flags]

Aliases:
  gen-views, gen-tmpl, gen

Flags:
  -h, --help            help for gen-views
  -o, --output string   output directory (default generate-views) (default "generated-views")
  -v, --views string    views to generate (comma separated string) (default "...")
``` 

###### Debug server
```
Debug UI server for dev purposes

Usage:
  medtune-beta debug [flags]

Aliases:
  debug, debug-server

Flags:
  -h, --help            help for debug
  -p, --port int        port (default 8005)
  -s, --static string   Static files directory (default "./static")
```

###### Run server
```
Run Medtune beta server

Usage:
  medtune-beta start [flags]

Aliases:
  start, run, run-server

Flags:
  -f, --file string          Configuration file name (default "./config.yml")
  -g, --gin-mode int         Gin server mode [0 OR 1]
  -h, --help                 help for start
  -p, --port int             port (default 8005)
  -s, --static string        Static files directory (default "./static")
  -x, --syncdb               Sync database before start (default true)
  -w, --wait                 Wait all services to go up
  -c, --wait-attempts int    Wait max attempts (default 30)
  -t, --wait-timestamp int   Wait timestamp (default 1)
```



### Changelog

See [CHANGELOG.md](CHANGELOG.md)

### Contributing

For contributions please refer to our friendly code of conduct in [how-we-work](https://github.com/medtune/how-we-work) repository :smile:.

### Maintainers

See [MAINTAINERS](MAINTAINERS.txt)

### License

Medtune Beta Platform is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt).
