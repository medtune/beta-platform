compile:
	@echo "building alpine binaries"
	GOOS=linux go build -o medtune-beta ./cmd/main.go

build: compile
	@echo "building alpine docker image"
	docker build -t medtune/platform:beta.alpine-v0.0.4 -f Dockerfile .

push:
	@echo "pushing alpine image to docker hub"
	docker push medtune/platform:beta.alpine-v0.0.4

pull:
	@echo "pulling alpine image from docker hub"
	docker pull medtune/platform:beta.alpine-v0.0.4 

up:
	docker-compose up

gen:
	go run ./cmd/main.go gen

down:
	docker-compose down

clean:
	rm medtune-beta

clean-gen:
	rm -rf genered-views
