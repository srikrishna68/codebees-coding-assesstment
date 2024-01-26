package src

import (
	"context"
	`log`
	`time`

	"codebees-coding-assesstment/pb"

	"google.golang.org/grpc"
)

// BlogClient is a client to call blog service RPCs
type BlogClient struct {
	service pb.BlogServiceClient
}

// NewBlogClient returns a new blog client
func NewBlogClient(cc *grpc.ClientConn) *BlogClient {
	service := pb.NewBlogServiceClient(cc)
	return &BlogClient{service}
}

func (blogClient *BlogClient) CreatePost(req *pb.Post) (*pb.PostResponse, error) {

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := blogClient.service.CreatePost(ctx, req)
	if err != nil {
		log.Fatal("cannot create post: ", err)

		return nil, err
	}

	log.Printf("created post with id: %s", res.Post.PostId)
	return res, nil
}

func (blogClient *BlogClient) ReadPost(req *pb.PostID) (*pb.PostResponse, error) {

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := blogClient.service.ReadPost(ctx, req)
	if err != nil || res.Status == pb.PostResponse_ERROR {
		log.Fatal("cannot read post: ", res.Error)

		return nil, err
	}

	log.Printf("read post with id: %s", res.Post.PostId)
	return res, nil
}

func (blogClient *BlogClient) UpdatePost(req *pb.Post) (*pb.PostResponse, error) {

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := blogClient.service.UpdatePost(ctx, req)
	if err != nil {
		log.Fatal("cannot update post: ", err)

		return nil, err
	}

	log.Printf("Updated post with id: %s", res.Post.PostId)
	return res, nil
}

func (blogClient *BlogClient) DeletePost(req *pb.PostID) (*pb.DeleteResponse, error) {

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := blogClient.service.DeletePost(ctx, req)
	if err != nil {
		log.Fatal("cannot delete post: ", err)

		return nil, err
	}

	log.Printf("deleted post with id: %s", req.Id)

	return res, nil
}
