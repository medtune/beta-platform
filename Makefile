PROJECT=beta-platform
OS_TYPE=$(shell uname -a)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
MAJOR=0
MINOR=1
PATCH=4
REVISION=alpha
VERSION=v$(MAJOR).$(MINOR).$(PATCH)
GOVERSION=1.11
LONGVERSION=v$(MAJOR).$(MINOR).$(PATCH)-$(REVISION)
CWD=$(shell pwd)
VPATH=github.com/medtune/beta-platform/pkg
PROJECTPATH=$(CWD)
AUTHORS=Hilaly.Mohammed-Amine/El.bouchti.Alaa
OWNERS=$(AUTHORS)
LICENSETYPE=Apache-v2.0
LICENSEURL=https://raw.githubusercontent.com/medtune/beta-platform/master/LICENSE.txt

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


release-alpine:
	GOOS=linux go build \
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

build-alpine: release-alpine
	@echo building linux prod container
	docker build \
		-t medtune/beta-platform:prod-alpine \
		-f build/prod.linux.Dockerfile \
		.

build-base:
	@echo building base image
	docker build \
		-t medtune/beta-platform:base \
		-f build/base.Dockerfile \
		.

	docker tag \
		medtune/beta-platform:base \
		medtune/beta-platform:go-1.10-linux-$(VERSION)-base

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


build-all: build-base build-compile build-prod

push-image:
	docker push medtune/beta-platform:$(VERSION)
	docker push medtune/beta-platform:latest

tests:
	@echo running global test
	docker build \
		-t medtune/beta-platform:test \
		-f test/test.Dockerfile \
		.

test-cov:
	@echo running global code coverage tests
	docker build \
		-t medtune/beta-platform:test \
		-f test/test-codecov.Dockerfile \
		--build-arg CODECOV_TOKEN=$(CODECOV_TOKEN) \
		.


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

chexray-pp-helper:
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
	chexray-dn-121

start-capsules: 
	docker start mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121

stop-capsules:
	docker stop mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121

kill-capsules:
	docker kill mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121

	docker rm mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2 \
		chexray-dn-121

start:
	./medtune-beta start \
		-f dev.config.yml \
		--syncdb \
		--create-users \
		--wait

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
	rm -f static/demos/mura/images/*mn_v2_cam.png
	rm -f static/demos/chexray/images/*mn_v2_cam.png
	rm -rf static/demos/chexray/images/000*

verify:
	GO111MODULE=on go mod verify

vendor:
	GO111MODULE=on go mod vendor

