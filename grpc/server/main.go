package main

import (
	`fmt`
	`log`
	`net`

	`codebees-coding-assesstment/pb`
	`codebees-coding-assesstment/src`

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello inside main")
	// Initialize server
	newBlogStore := src.NewInMemoryBlogStore()
	blogServer := src.NewBlogServer(newBlogStore)
	//
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }
	//  _ = runGRPCServer(blogServer, lis)
	// // Start server
	//

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, blogServer)

	fmt.Println("Server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
