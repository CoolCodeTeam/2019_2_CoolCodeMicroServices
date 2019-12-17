fuser -n tcp -k 8001 ; cd ./users && go run . &
fuser -n tcp -k 8002 ; cd ./chats && go run . &
fuser -n tcp -k 8003 ; cd ./notifications && go run . &
fuser -n tcp -k 8004 ; cd ./messages && go run .
