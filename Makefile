build:
	@echo "building default binaries"
	go build -o medtune-beta ./cmd/main.go

build-alpine:
	@echo "building binaries ~ linux alpine"
	GOOS=linux go build -o medtune-beta

docker-build:
	docker build -t medtune/beta-platform:v0.0.3 .

docker-push:
	docker push medtune/beta-platform:v0.0.3

clean:
	rm medtune-beta



