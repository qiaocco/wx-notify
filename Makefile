DOCKER_REPO ?= qiaocc
IMG ?= wx-notify:0.1

docker-release: docker-build docker-push

docker-build:
	docker build -t ${IMG} .

docker-run:
	docker run -p 8080:8080 ${IMG}

docker-push:
	docker tag ${IMG} ${DOCKER_REPO}/${IMG}
	docker push ${DOCKER_REPO}/${IMG}

build:
	go build -o ./tmp/wx-notify .

run: fmt vet
	go run ./main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

