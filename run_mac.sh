kill $(lsof -i :8001)  ; cd ./users && go run . &
kill $(lsof -i :8002)  ; cd ./chats && go run . &
kill $(lsof -i :8003); cd ./notifications && go run . &
kill $(lsof -i :8004); cd ./messages && go run . &

