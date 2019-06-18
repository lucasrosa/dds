.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/get adapters/lambda/domain/get/*

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
