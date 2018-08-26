OS_TYPE='uname -a'
PROJECT=beta-platform
VERSION=v0.1.2
GO_VERSION=1.10

build-base:
	@echo building base image
	docker build -t medtune/beta-platform:base -f build/base.Dockerfile .
	docker tag medtune/beta-platform:base medtune/beta-platform:go-1.10-linux-v0.1.2-base

build-compile:
	@echo building build image
	docker build -t medtune/beta-platform:build -f build/build.Dockerfile .

	@echo tag build images
	docker tag medtune/beta-platform:build medtune/beta-platform:go-1.10-linux-v0.1.2-build
	docker tag medtune/beta-platform:build medtune/beta-platform:build

build-prod:
	@echo building prod image
	docker build -t medtune/beta-platform:prod -f build/prod.Dockerfile .

	@echo tag prod images
	docker tag medtune/beta-platform:prod medtune/beta-platform:go-1.10-linux-v0.1.2-prod
	docker tag medtune/beta-platform:prod medtune/beta-platform:prod

tests:
	@echo running global test
	docker build -t medtune/beta-platform:test -f test/test.Dockerfile .

test-cov:
	@echo running global code coverage tests
	docker build -t medtune/beta-platform:test -f test/test-codecov.Dockerfile .

up:
	docker-compose up

down:
	docker-compose down

clean:
	rm medtune-beta

clean-gen:
	rm -rf genered-views
