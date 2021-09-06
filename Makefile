compose:
	docker-compose build && docker-compose up -d
compose-with-migration:
	docker-compose -f docker-compose.yml -f docker-compose-migrate.yml up -d
compose-down:
	docker-compose down --remove-orphans
protoc:
	 rm -rf ./protocol/pb && protoc --proto_path=protocol/proto protocol/proto/*.proto --go_out=./protocol --go-grpc_out=./protocol
grpcui:
	grpcui -plaintext localhost:$(PORT)

.PHONY: compose, compose-with-migration, compose-down, protoc, grpcui