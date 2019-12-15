all: build

.PHONY: build
build:
	go build -o ./bin/users ./users

clean:
	rm -rf ./bin
