package src

import (
	`context`
	`fmt`
	`log`

	`github.cisco.com/srikrishna68/codebees-coding-assesstment/pb`
	`google.golang.org/grpc/codes`
	`google.golang.org/grpc/status`
)

// BlogServer is the server that provides blog services
type BlogServer struct {
	pb.UnimplementedBlogServiceServer
	blogStore BlogStore
}

// NewBlogServer returns a new BlogServer
func NewBlogServer(blogStore BlogStore) *BlogServer {
	return &BlogServer{
		blogStore: blogStore,
	}
}

func (blogServer *BlogServer) CreatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {
	log.Printf("receive a create post request: %s", req.Title)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	err := validatePost(req, false)
	if err != nil {
		log.Printf("Error while validating creating post, Error %s", err)
		return nil, err
	}
	// Storing post
	post, err := blogServer.blogStore.Create(req)
	if err != nil {
		log.Printf("Error in creating post, Error %s", err)
		return nil, err
	}
	log.Printf("saved post with title: %s", req.Title)

	return &pb.PostResponse{Post: post}, nil
}

func (blogServer *BlogServer) ReadPost(ctx context.Context, req *pb.PostID) (*pb.PostResponse, error) {
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	post, err := blogServer.blogStore.Read(req.Id)
	if err != nil {
		return &pb.PostResponse{
			Status: pb.PostResponse_ERROR,
			Error:  err.Error(),
		}, nil
	}

	return &pb.PostResponse{
		Status: pb.PostResponse_SUCCESS,
		Post:   post,
	}, nil
}

func (blogServer *BlogServer) UpdatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	_, err := blogServer.ReadPost(ctx, &pb.PostID{Id: req.PostId})
	if err != nil {
		log.Printf("post id not present in database %s", req.PostId)
		return nil, fmt.Errorf("post id not present in database %s", req.PostId)
	}

	err = validatePost(req, true)
	if err != nil {
		log.Printf("Error while validating update post, Error %s", err)
		return nil, err
	}

	post, err := blogServer.blogStore.Update(req)
	if err != nil {
		log.Printf("Error in updating post, Error %s", err)
		return nil, err
	}
	log.Printf("updated post with title: %s", req.Title)

	return &pb.PostResponse{Post: post}, nil
}

func (blogServer *BlogServer) DeletePost(ctx context.Context, req *pb.PostID) (*pb.DeleteResponse, error) {
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	response := blogServer.blogStore.Delete(req.Id)

	return response, nil
}

func validatePost(post *pb.Post, update bool) error {
	if post.Title == "" || len(post.Title) == 0 {
		return fmt.Errorf("empty title, title should be provided (%s)", post.Content)
	}

	if post.Content == "" || len(post.Content) == 0 {
		return fmt.Errorf("empty Content, Content should be provided (%s)", post.Title)
	}
	if post.Author == "" || len(post.Author) == 0 {
		return fmt.Errorf("empty Author, author should be provided (%s)", post.Author)
	}
	if update {
		if len(post.PublicationDate) != 0 {
			return fmt.Errorf(" PublicationDate not be changed for update scenarios(%s)", post.Title)
		}
	} else {
		if post.PublicationDate == "" || len(post.PublicationDate) == 0 {
			return fmt.Errorf("empty PublicationDate, PublicationDate should be provided (%s)", post.PublicationDate)
		}
	}

	if len(post.Tags) == 0 {
		return fmt.Errorf("empty tags, tags should be provided(%s)", post.Tags)
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
