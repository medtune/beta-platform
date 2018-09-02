PROJECT=beta-platform
MAJOR=0
MINOR=1
PATCH=3
GITCOMMIT=$(shell git rev-parse HEAD)
OS_TYPE=$(shell uname -a)
GOVERSION=1.10
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
VERSION=v$(MAJOR).$(MINOR).$(PATCH)
VPATH=github.com/medtune/beta-platform/pkg

release:
	go build \
		-tags=prod \
		-o medtune-beta \
		-ldflags="\
			-X $(VPATH).GitCommit=$(GITCOMMIT) \
			-X $(VPATH).Major=$(MAJOR) \
			-X $(VPATH).Minor=$(MINOR) \
			-X $(VPATH).Patch=$(PATCH) \
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
			-X $(VPATH).BuildDate=$(BUILDDATE)" \
		cmd/main.go \

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
		-v $(HOME)/go/src/github.com/medtune/beta-platform/static/demos/mura/images:/medtune/data \
		medtune/capsul:mura-mn-v2-cam

mura-irn-v2:
	docker run -dti \
		--name mura-irn-v2 \
		-p 10021:10021 \
		medtune/capsul:mura-irn-v2


run-capsules: mnist \
	inception \
	mura-mn-v2 \
	mura-mn-v2-cam \
	mura-irn-v2 \

kill-capsules:
	docker kill mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2

	docker rm mnist \
		inception \
		mura-mn-v2 \
		mura-mn-v2-cam \
		mura-irn-v2

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
	rm -rf static/demos/mura/images/*cam_mn_v2.png