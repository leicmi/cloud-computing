.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./status/status
	
build:
	GOOS=linux GOARCH=amd64 go build -o status/status ./status