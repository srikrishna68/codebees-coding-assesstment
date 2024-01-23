clean:
	rm pb/*
	rm swagger/*

gen:
	protoc --go_out=. --go-grpc_out=./proto proto/*.proto

run:
	go run src/main.go