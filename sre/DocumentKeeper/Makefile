DOCKER_IMAGE := documentkeeper

build:
	docker build -t ${DOCKER_IMAGE} . -f Dockerfile
.PHONY: build 

run: build
	docker run --rm -it -p 4096:4096 ${DOCKER_IMAGE}
.PHONY: run
