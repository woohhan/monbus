.DEFAULT_GOAL := usage

test:
	go run cmd/monbus/main.go -logtostderr=true

run:
	CGO_ENABLED=0 go build -o bin/monbus cmd/monbus/main.go
	docker build -t quay.io/woohhan/monbus:latest .
	docker push quay.io/woohhan/monbus:latest
	docker rm --force monbus || true
	docker run --name monbus -d --restart always -e "TZ=Asia/Seoul" quay.io/woohhan/monbus:latest

log:
	docker logs monbus

usage:
	@echo "usage: make [command]"
	@echo "test"
	@echo "run"
	@echo "log"
