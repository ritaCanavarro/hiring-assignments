DOCKER_IMAGE := documentkeeper

build:
	docker build -t ${DOCKER_IMAGE} .
.PHONY: build 

run: build
	docker run --rm -it -p 4096:4096 ${DOCKER_IMAGE}
.PHONY: run
