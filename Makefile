.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ddg-api-free ./main.go

.PHONY: docker
docker:
	docker build . -t samge/ddg-api-free -f docker/Dockerfile
