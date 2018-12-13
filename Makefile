PROJECT=beta-platform
OS_TYPE=$(shell uname -a)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
MAJOR=0
MINOR=1
PATCH=6
REVISION=alpha
VERSION=v$(MAJOR).$(MINOR).$(PATCH)
GOVERSION=1.11
LONGVERSION=v$(MAJOR).$(MINOR).$(PATCH)-$(REVISION)
CWD=$(shell pwd)
VPATH=github.com/medtune/beta-platform/internal
PROJECTPATH=$(CWD)
AUTHORS=El.bouchti.Alaa/Hilaly.Mohammed-Amine
OWNERS=$(AUTHORS)
LICENSETYPE=Apache-v2.0
LICENSEURL=https://raw.githubusercontent.com/medtune/beta-platform/master/LICENSE.txt


# Compile prod version of binaries
release:
	go build \
		-tags=prod \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
		 	-X $(VPATH).Revision=$(REVISION) \
			-X $(VPATH).Authors=$(AUTHORS) \
			-X $(VPATH).Owners=$(OWNERS) \
			-X $(VPATH).LicenseURL=$(LICENSEURL) \
			-X $(VPATH).LicenseType=$(LICENSETYPE) \
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main_prod.go \

# Compile all possible commands binaries
release-cmd:
	go build \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
			-X $(VPATH).Revision=$(REVISION) \
			-X $(VPATH).Authors=$(AUTHORS) \
			-X $(VPATH).Owners=$(OWNERS) \
		    -X $(VPATH).LicenseURL=$(LICENSEURL) \
			-X $(VPATH).LicenseType=$(LICENSETYPE) \
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main.go


# Compile developpement version
# GOCV under darwin/amd64
release-dev:
	go build \
		-tags="gocv" \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
 			-X $(VPATH).Revision=$(REVISION) \
			-X $(VPATH).Authors=$(AUTHORS) \
			-X $(VPATH).Owners=$(OWNERS) \
			-X $(VPATH).LicenseURL=$(LICENSEURL) \
			-X $(VPATH).LicenseType=$(LICENSETYPE) \
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main.go \


# Compile linux version
# 30Mb IMAGE !!
release-linux:
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	GOOS=linux \
	go build \
		-tags="prod" \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
 			-X $(VPATH).Revision=$(REVISION) \
			-X $(VPATH).Authors=$(AUTHORS) \
			-X $(VPATH).Owners=$(OWNERS) \
			-X $(VPATH).LicenseURL=$(LICENSEURL) \
			-X $(VPATH).LicenseType=$(LICENSETYPE) \
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main_prod.go \


# Compile debug binaries
release-debug:
	go build \
		-tags="debug" \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
 			-X $(VPATH).Revision=$(REVISION) \
			-X $(VPATH).Authors=$(AUTHORS) \
			-X $(VPATH).Owners=$(OWNERS) \
			-X $(VPATH).LicenseURL=$(LICENSEURL) \
			-X $(VPATH).LicenseType=$(LICENSETYPE) \
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main_debug.go \


# Build light image version (30mb)
build-alpine:  build-base build-compile
	docker rm build-stage
	docker run --name=build-stage medtune/beta-platform:build
	docker cp build-stage:/go/src/github.com/medtune/beta-platform/medtune-beta $(PWD)/medtune-beta
	@echo building linux prod container
	docker build \
		-t medtune/beta-platform:prod-alpine \
		-f build/prod.linux.Dockerfile \
		.


# Build base image
# Doesnt compile anything
build-base:
	@echo building base image
	docker build \
		-t medtune/beta-platform:base \
		-f build/base.Dockerfile \
		.

	docker tag \
		medtune/beta-platform:base \
		medtune/beta-platform:go-1.10-linux-$(VERSION)-base


# 1 - Build compile use build-base image and compile
# the binaries using this makefile
# 2 - Tag image specifications
build-compile:
	@echo building build image
	docker build \
		-t medtune/beta-platform:build \
		-f build/build.Dockerfile \
		.

	@echo tag build images
	docker tag \
		medtune/beta-platform:build \
		medtune/beta-platform:go-1.10-linux-$(VERSION)-build

	docker tag \
		medtune/beta-platform:build \
		medtune/beta-platform:build


# Build production image (based on build-compile)
build-prod:
	@echo building prod image
	docker build \
		-t medtune/beta-platform:prod \
		-f build/prod.Dockerfile \
		.

	@echo tag prod images
	docker tag \
		medtune/beta-platform:prod \
		medtune/beta-platform:go-1.10-linux-$(VERSION)-prod

	docker tag \
		medtune/beta-platform:prod \
		medtune/beta-platform:$(VERSION)

	docker tag \
		medtune/beta-platform:prod \
		medtune/beta-platform:latest


# Build k8s
build-k8s: #build-base build-compile
	@echo building kubernetes prod image
	docker build \
		-t medtune/beta-platform:k8s \
		-f build/prod.k8s.Dockerfile \
		.

# Just let the make make the magic
build-all: build-base build-compile build-prod


# Public images in docker hub (hub.docker.com/medtune)
push-image:
	docker push medtune/beta-platform:$(VERSION)
	docker push medtune/beta-platform:latest

gen-k8s:
	rm -rf deploy/kubernetes/*
	kompose convert -f deployments/docker-compose.k8s-gen.yml -o deployments/k8s-generated

# Test package
tests:
	@echo running global test
	docker build \
		-t medtune/beta-platform:test \
		-f test/test.Dockerfile \
		.


# Test coverage and push results to codecov
test-cov:
	@echo running global code coverage tests
	docker build \
		-t medtune/beta-platform:test \
		-f test/test-codecov.Dockerfile \
		--build-arg CODECOV_TOKEN=$(CODECOV_TOKEN) \
		.

	docker run \
		-v "$(PWD)/shared":/shared \
		medtune/beta-platform:test



# setup capsules
mnist:
	docker run -dti \
		--name mnist \
		-p 10000:10000 \
		medtune/capsul:mnist


inception:
	docker run -dti \
		--name inception \
		-p 10010:10010 \
		medtune/capsul:inception


mura-mn-v2:
	docker run -dti \
		--name mura-mn-v2 \
		-p 10020:10020 \
		medtune/capsul:mura-mn-v2


mura-mn-v2-cam:
	docker run -dti \
		--name mura-mn-v2-cam \
		-p 11020:11020 \
		-v $(PROJECTPATH)/static/demos/mura/images:/medtune/data \
		medtune/capsul:mura-mn-v2-cam


mura-irn-v2:
	docker run -dti \
		--name mura-irn-v2 \
		-p 10021:10021 \
		medtune/capsul:mura-irn-v2


chexray-dn-121:
	docker run -dti \
		--name chexray-dn-121 \
		-p 10031:10031 \
		medtune/capsul:chexray-dn-121


chexray-dn-121-cam:
	docker run -dti \
		--name chexray-dn-121-cam \
		-p 10031:10031 \
		medtune/capsul:chexray-dn-121-cam


chexray-pp:
	docker run -dti \
			--name chexray-pp-helper \
			-p 12030:12030 \
			-v $(PROJECTPATH)/static/demos/chexray/images:/medtune/data \
			medtune/capsul:chexray-pp-helper


run-capsules: mnist \
	inception \
	mura-mn-v2 \
	mura-mn-v2-cam \
	mura-irn-v2 \
	chexray-dn-121 \
	chexray-pp


start-capsules: 
	docker start mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121 \
		chexray-pp-helper


stop-capsules:
	docker stop mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121 \
		chexray-pp-helper


kill-capsules:
	docker kill mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121 \
		chexray-pp-helper

	docker rm mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121 \
		chexray-pp-helper


start:
	./medtune-beta start \
		-f etc/dev.config.yml \
		--syncdb \
		--create-users \
		--sync-cxpba \
		--wait \
		--gin-mode 1

run:
	go run -ldflags="\
		-X $(VPATH).Major=$(MAJOR) \
		-X $(VPATH).Minor=$(MINOR) \
		-X $(VPATH).Patch=$(PATCH)" \
		-tags=gocv \
		./cmd/main.go \
		start \
		-f etc/dev.config.yml \
		--syncdb \
		--sync-cxpba \
		--create-users \
		--gin-mode 0

debug:
	./medtune-beta debug

up:
	docker-compose up

down:
	docker-compose down
  
clean:
	rm medtune-beta

clean-gen:
	rm -rf genered-views
	

clean-demos:
	rm -f static/demos/mura/images/*_mn_v2_cam.*
	rm -f static/demos/mura/images/[^image_*]*

	rm -f static/demos/chexray/images/[^debug.png][^image_*]*
	rm -f static/demos/chexray/images/[^image_*]*

verify:
	GO111MODULE=on go mod verify

vendor:
	GO111MODULE=on go mod vendor


loc:
	scc --pbl static/reveal.js-3.7.0,vendor,.dev,.git


start-scene: start-capsules
	docker start postgreshost redis


stop-scene: stop-capsules
	docker stop postgreshost redis


kill-scene: kill-capsules
	docker kill postgreshost redis
	docker rm postgreshost redis


