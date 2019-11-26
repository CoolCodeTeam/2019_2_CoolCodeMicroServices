go test -coverpkg=./... -coverprofile cover.out.tmp ./...
cat cover.out.tmp | grep -v "_mock.go" | grep -v ".pb"  | grep -v "_grpc_"| grep -v "_easyjson.go"> cover.out
go tool cover -func cover.out