.DEFAULT_GOAL := usage

bb:
	CGO_ENABLED=0 go build -o bin/monbus cmd/monbus/main.go

rb:
	make bb
	./bin/monbus --logtostderr -v 2

bi:
	make bb
	docker build -t quay.io/woohhan/monbus:canary .

ri:
	make bi
	docker run quay.io/woohhan/monbus:canary

pi:
	make bi
	docker tag quay.io/woohhan/monbus:canary quay.io/woohhan/monbus:latest
	docker push quay.io/woohhan/monbus:latest

c:
	docker rmi --force quay.io/woohhan/monbus:canary
	rm bin/monbus

usage:
	@echo "usage: make [command]"
	@echo "bb     build binary"
	@echo "rb     run binary"
	@echo "bi     build image"
	@echo "ri     run image"
	@echo "pi     push image"
	@echo "c      clean"