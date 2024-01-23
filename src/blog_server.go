package src

import (
	`context`
	`fmt`
	`log`

	`src/database`

	pb `github.com/srikrishna68/codebees-coding-assesstment/proto/api`
	`google.golang.org/grpc/codes`
	`google.golang.org/grpc/status`
)

// BlogServer is the server that provides blog services
type BlogServer struct {
	blogStore database.BlogStore
}

// NewBlogServer returns a new BlogServer
func NewBlogServer(blogStore database.BlogStore) *BlogServer {
	return &BlogServer{
		blogStore: blogStore,
	}
}

func (blogServer *BlogServer) CreatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {
	// Implement create logic
	// ...

	log.Printf("receive a create-laptop request with id: %s", req.Id)

	if err := contextError(ctx); err != nil {
		return nil, err
	}
	// Storing post
	blogServer.blogStore.Save(req)

	log.Printf("saved blog with id: %s", req.Id)

	return &pb.PostResponse{Post: req}, nil
}

func (blogServer *BlogServer) ReadPost(ctx context.Context, req *pb.PostID) (*pb.PostResponse, error) {
	// Implement read logic
	// ...

	// Example: Retrieving post from in-memory
	// post, exists := s.posts[req.Id]
	// if !exists {
	// 	return &pb.PostResponse{Error: "Post not found"}, nil
	// }

	// return &pb.PostResponse{Post: post}, nil
	return nil, nil
}

func (blogServer *BlogServer) UpdatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {
	// Implement update logic
	// ...

	// Example: Updating post in-memory
	// blogServer.posts[req.Id] = req

	return &pb.PostResponse{Post: req}, nil
}

func (blogServer *BlogServer) DeletePost(ctx context.Context, req *pb.PostID) (*pb.DeleteResponse, error) {
	// Implement delete logic
	// ...

	// Example: Deleting post from in-memory
	// delete(re.posts, req.Id)

	return &pb.DeleteResponse{Success: true}, nil
}

func validatePost(post *pb.Post) error {
	if post.Title == "" || len(post.Title) == 0 {
		return fmt.Errorf("empty title, title should be provided post content (%s)", post.Content)
	}

	if post.Content == "" || len(post.Content) == 0 {
		return fmt.Errorf("empty title, title should be provided post content (%s)", post.Content)
	}
	if post == "" || len(post.Author) == 0 {
		return fmt.Errorf("empty title, title should be provided post content (%s)", post.Content)
	}
	if post.Title == "" || len(post.Title) == 0 {
		return fmt.Errorf("empty title, title should be provided post content (%s)", post.Content)
	}
	if post.Title == "" || len(post.Title) == 0 {
		return fmt.Errorf("empty title, title should be provided post content (%s)", post.Content)
	}

	return nil
}
func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
