
proto:
	protoc -I=api --go_out=. --go-grpc_out=. movie.proto

bench:
	go test -bench=. ./...
