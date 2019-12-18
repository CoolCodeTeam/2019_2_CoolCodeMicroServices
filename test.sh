rm messages/go.mod
rm users/go.mod
rm notifications/go.mod
rm chats/go.mod

go test -coverpkg=./... -coverprofile cover.out.tmp ./...
cat cover.out.tmp | grep -v "_mock.go" | grep -v ".pb"  |  grep -v "_easyjson.go"> cover.out
go tool cover -func cover.out

cp ./messages/temp/go.mod ./messages/go.mod
cp ./users/temp/go.mod ./users/go.mod
cp ./notifications/temp/go.mod ./notifications/go.mod
cp ./chats/temp/go.mod ./chats/go.mod