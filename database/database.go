package database

import (
	`errors`
	`fmt`
	`sync`

	"github.com/jinzhu/copier"
	pb "github.com/srikrishna68/codebees-coding-assesstment/blog/api"
)

// ErrAlreadyExists is returned when a record with the same
var ErrAlreadyExists = errors.New("record already exists")

// InMemoryBlogStore stores laptop in memory
type InMemoryBlogStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Post
}

// Save saves the laptop to the store
func (store *InMemoryBlogStore) Save(blog *pb.Post) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[blog.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(blog)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(laptop *pb.Post) (*pb.Post, error) {
	other := &pb.Post{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}