.PHONY: protoc-setup
protoc-setup:
	docker build -t english-diary-protoc-client -f ../build/Dockerfile_client .
	docker run --rm -v `pwd`:/usr/local/develop/src/client/ english-diary-protoc-client

.PHONY: protoc
protoc: protoc-grpc

.PHONY: protoc-grpc
protoc-grpc:
	protoc -I=proto/ proto/*.proto --js_out=import_style=commonjs:proto/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:proto/

.PHONY: cp-proto
cp-proto:
	cp ../proto/*.proto proto/
