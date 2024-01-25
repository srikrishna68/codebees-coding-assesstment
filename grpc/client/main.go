package main

import (
	"log"

	`github.cisco.com/srikrishna68/codebees-coding-assesstment/pb`
	`github.cisco.com/srikrishna68/codebees-coding-assesstment/sample`
	`github.cisco.com/srikrishna68/codebees-coding-assesstment/src`
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
	c := src.NewLaptopClient(conn)

	testServer(c)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Blog Title: ")
	// title, _ := reader.ReadString('\n')
	// title = strings.Trim(title, "\n")
	// fmt.Print("Enter Blog contents: ")
	// body, _ := reader.ReadString('\n')
	// body = strings.Trim(body, "\n")
	//
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// _, err = c.CreatePost(ctx, &pb.Post{
	// 	Title:           title,
	// 	Content:         body,
	// 	Author:          title,
	// 	PublicationDate: time.DateOnly,
	// 	Tags:            []string{"1234"},
	// })
	// if err != nil {
	// 	log.Fatalf("Could not create Blog Post :%v", err)
	// }
	//
	// log.Printf("Post Successfully Created")

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
