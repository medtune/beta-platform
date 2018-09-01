OS_TYPE='uname -a'
PROJECT=beta-platform
VERSION=v0.1.2
GO_VERSION=1.10

build-base:
	@echo building base image
	docker build \
		-t medtune/beta-platform:base \
		-f build/base.Dockerfile \
		.

	docker tag \
		medtune/beta-platform:base \
		medtune/beta-platform:go-1.10-linux-v0.1.2-base

build-compile:
	@echo building build image
	docker build \
		-t medtune/beta-platform:build \
		-f build/build.Dockerfile \
		.

	@echo tag build images
	docker tag \
		medtune/beta-platform:build \
		medtune/beta-platform:go-1.10-linux-v0.1.2-build

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
		medtune/beta-platform:go-1.10-linux-v0.1.2-prod

	docker tag \
		medtune/beta-platform:prod \
		medtune/beta-platform:v0.1.2

	docker tag \
		medtune/beta-platform:prod \
		medtune/beta-platform:latest

build-all: build-base build-compile build-prod

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
	go run -tags=gocv \
		cmd/main.go start \
		-f dev.config.yml \
		--syncdb --create-users --wait

debug:
	go run cmd/main.go debug

up:
	docker-compose up

down:
	docker-compose down

clean:
	rm medtune-beta

clean-gen:
	rm -rf genered-views
