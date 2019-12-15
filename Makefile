all: build

.PHONY: build
build:
	go build -o ./bin/users ./users
	go build -o ./bin/chats ./chats
	go build -o ./bin/notifications ./notifications
	go build -o ./bin/messages ./messages

clean:
	rm -rf ./bin
