VERSION = latest 
IMAGE = image-name:${VERSION}

run:
	go run src/main.go -p=8001

build:
	go build  -v -mod=vendor -o app src/main.go

docker:
	docker build -f Dockerfile -t ${IMAGE} .

vendor:
	go mod vendor
