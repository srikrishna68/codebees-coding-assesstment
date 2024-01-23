package database

import pb "github.com/srikrishna68/codebees-coding-assesstment/blog/api"

// LaptopStore is an interface to store laptop
type BlogStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Post) error
}
