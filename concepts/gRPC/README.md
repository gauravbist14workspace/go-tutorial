1. Make sure protobuf(proto compiler) is installed on your machine. On MAC machine, run the below command:
```js
	brew install protobuf
```

2. Install the dependencies before running protoc command. These libraries helps with the auto code generation.
```js
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. make directory for placing the generated stub
	we gotta make sure that the .proto file has a property `option go_package = "grpc_demo/coffee_shop";`
```js
    mkdir coffee_shop
```

3. Now run the below command to created the actual stubs
```js
	protoc \
	--go_out=./coffee_shop \
	--go-grpc_out=./coffee_shop \
	./proto/coffee_shop.proto
```

<br>

[Note] In case we want to work more easily with the relative paths, in your .proto file, set `option go_package = "./proto"`. And then use the below command instead. It will have everything in a need and easy to use proto directory which can be imported in server, client code using `import pb "grpc_demo/proto"`
```js
	protoc \
	--go_out=. \
	--go-grpc_out=. \
	proto/coffee_shop.proto
```

Now to start using the code from generated files, like marshaling or the protocol buffer related code, you need to get these 2 libraries as well:
```go
	go get google.golang.org/protobuf/proto		// only if required
	go get google.golang.org/grpc
```
