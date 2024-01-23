package sample

import pb "github.com/srikrishna68/codebees-coding-assesstment/proto"

// NewCPU returns a new sample CPU
func NewBlog() *pb.Post {

	blog := &pb.Post{
		Id:      "generate new UUID",
		Title:   "New Book",
		Content: "ABCD",
	}

	return blog
}
