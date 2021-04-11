.PHONY: proto

proto:
	docker build -t bunting-protogen -f ./build/Dockerfile.protogen .
	docker run --rm -v /etc/passwd:/etc/passwd -u $(shell id -u):$(shell id -g) -v $(shell pwd):/home/me/bunting.io/bunting bunting-protogen /home/me/bunting.io/bunting/build/generate-proto.sh
