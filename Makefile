clean:
	rm pb/*
	rm swagger/*

gen:
	protoc --go_out=:pb --go-grpc_out=:pb proto/*.proto

server:
	go run grpc/server/main.go

client:
	go run grpc/client/main.go