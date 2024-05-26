proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

server:
	export STOCKS_API_KEY=4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD
	go run main.go

stocks_cli:
	cd stocks_cli
	go install
	cd ..

evans:
	evans --host 127.0.0.1 --port 9090 -r repl

.PHONY: proto server evans