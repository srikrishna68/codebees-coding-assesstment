package main

import (
	"log"

	`codebees-coding-assesstment/pb`
	`codebees-coding-assesstment/sample`
	`codebees-coding-assesstment/src`

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := src.NewBlogClient(conn)

	testServer(c)
}

func testServer(blogClient *src.BlogClient) {
	post := sample.NewPost()
	response, err := blogClient.CreatePost(post)
	if err != nil {
		log.Println("Error in creating post", post.PostId)
	}

	log.Println("Successfully created first post", response)

	// Read Post
	postID := &pb.PostID{Id: response.Post.PostId}
	readResponse, err := blogClient.ReadPost(postID)
	if err != nil {
		log.Println("Error in reading post", postID)
	}
	log.Println("response from read", readResponse)

	// Delete Post
	deleteResponse, err := blogClient.DeletePost(postID)
	if err != nil {
		log.Println("Error in dele post", postID)
	}
	log.Println("response from delete", deleteResponse)

	// Read Post
	readResponse, err = blogClient.ReadPost(postID)
	if err != nil {
		log.Println("calling read after delete should be error", postID)
	}
	log.Println("response from read", readResponse)

}
