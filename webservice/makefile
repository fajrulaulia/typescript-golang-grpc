gen:
	protoc --proto_path=api api/*.proto --go_out=. --go-grpc_out=.

image-build:
	docker build -t webservice-grpc .

image-run:
	docker run -p 8080:8080 webservice-grpc:latest