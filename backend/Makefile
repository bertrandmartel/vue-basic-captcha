include .env
ifneq ("$(wildcard .env.secret)","")
include .env.secret
endif
export

clean:
	@rm -rf dist
	@mkdir -p dist

build: clean

run:
	goimports -w .
	gofmt -s -w .
	go run ./main.go

install:
	go get -u ./...

test:
	go test ./... --cover

