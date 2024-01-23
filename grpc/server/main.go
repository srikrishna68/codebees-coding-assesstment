package main

import (
	"fmt"
	`log`
	`net`

	"google.golang.org/grpc"
)



func main() {
	// Initialize server
	s := grpc.NewServer()
	// pb.RegisterBlogServiceServer(s, &server{posts: make(map[string]*pb.Post)})

	// Start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
