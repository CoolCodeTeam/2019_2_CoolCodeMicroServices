sudo lsof -i :8001  ; go run ./users &
sudo lsof -i :8002  ; go run ./chats &
sudo lsof -i :8003; go run ./notifications &
sudo lsof -i :8004  ; go run ./messages
