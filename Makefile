bin:
	@echo "building default binaries"
	go build -o medtune-beta ./cmd/main.go

build:
	@echo "building default image"
	docker build -t medtune/platform:beta.go1.10-v0.0.3 -f Dockerfile.go-1.10 .

push:
	@echo "pushing default image to docker hub"
	docker push medtune/platform:beta.go1.10-v0.0.3 

bin-alpine:
	@echo "building alpine binaries"
	GOOS=linux go build -o medtune-beta ./cmd/main.go

build-alpine:
	@echo "building alpine docker image"
	docker build -t medtune/platform:beta.alpine-v0.0.3 -f Dockerfile.alpine .

push-alpine:
	@echo "pushing alpine image to docker hub"
	docker push medtune/platform:beta.alpine-v0.0.3 

swarm-up:
	docker-compose up

swarm-down:
	docker-compose down

clean:
	rm medtune-beta

clean-gen:
	rm -rf genered-views
