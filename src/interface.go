package src

import "github.cisco.com/srikrishna68/codebees-coding-assesstment/pb"

// BlogStore is an interface to store posts
type BlogStore interface {
	// Create saves the post to the store
	Create(post *pb.Post) (*pb.Post, error)

	Update(post *pb.Post) (*pb.Post, error)
	Delete(postID string) *pb.DeleteResponse
	Read(postID string) (*pb.Post, error)
}
