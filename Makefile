compose:
	docker-compose build && docker-compose up -d
compose-down:
	docker-compose down
protoc:
	 rm -rf ./protocol/pb && protoc --proto_path=protocol/proto protocol/proto/*.proto --go_out=./protocol --go-grpc_out=./protocol
grpcui:
	grpcui -plaintext localhost:$(PORT)

.PHONY: compose, compose-down, protoc, grpcui