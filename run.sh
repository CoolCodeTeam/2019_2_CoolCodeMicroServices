fuser -n tcp -k 8001 & go run ./users &
fuser -n tcp -k 8002 & go run ./chats &
fuser -n tcp -k 8003 & go run ./notifications &
fuser -n tcp -k 8004 & go run ./messages