proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

server:
	go run main.go

stocks_cli:
	cd stocks_cli
	go install
	cd ..

.PHONY: proto server