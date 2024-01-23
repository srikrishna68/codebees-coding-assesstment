package sample

import "github.com/srikrishna68/codebees-coding-assesstment/pb"

// NewCPU returns a new sample CPU
func NewBlog() *pb.Post {

	blog := &pb.Post{
		Title:   "New Book",
		Content: "ABCD",
	}

	return blog
}
