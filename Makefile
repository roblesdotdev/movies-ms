
proto:
	protoc -I=api --go_out=. movie.proto

bench:
	go test -bench=. ./...
