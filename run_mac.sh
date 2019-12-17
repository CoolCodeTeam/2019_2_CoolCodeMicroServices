kill $(lsof -i :8001)  ; go run ./users &
kill $(lsof -i :8002)  ; go run ./chats &
kill $(lsof -i :8003); go run ./notifications &
kill $(lsof -i :8004); go run ./messages

