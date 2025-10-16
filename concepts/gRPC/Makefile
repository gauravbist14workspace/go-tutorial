build_proto_NOT_WORKING:
	protoc \
	--go_out=coffee_shop --go_out=paths=source_relative \
	--go-grpc_out=coffee_shop --go-grpc_out=paths=source_relative \
	./proto/coffee_shop.proto

build_proto2_WORKING:
	mkdir coffee_shop
	protoc \
	--go_out=./coffee_shop \
	--go-grpc_out=./coffee_shop \
	./proto/coffee_shop.proto