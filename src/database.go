package src

import (
	`errors`
	`fmt`
	`log`
	`sync`

	`codebees-coding-assesstment/src/utils`

	"codebees-coding-assesstment/pb"

	"github.com/jinzhu/copier"
)

// ErrAlreadyExists is returned when a record with the same
var ErrAlreadyExists = errors.New("record already exists")

// InMemoryBlogStore stores posts in memory
type InMemoryBlogStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Post
}

// NewInMemoryBlogStore returns a new InMemoryBlogStore
func NewInMemoryBlogStore() *InMemoryBlogStore {
	return &InMemoryBlogStore{
		data: make(map[string]*pb.Post),
	}
}

// Create saves the post to the store
func (store *InMemoryBlogStore) Create(post *pb.Post) (*pb.Post, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	uuid, err := utils.GenerateUuid()
	if err != nil {
		return nil, err
	}
	post.PostId = uuid

	other, err := deepCopy(post)
	if err != nil {
		return nil, err
	}

	store.data[other.PostId] = other
	return post, nil
}

func (store *InMemoryBlogStore) Update(post *pb.Post) (*pb.Post, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	_, ok := store.data[post.PostId]
	if !ok {
		log.Printf("post id not present in database %s", post.PostId)
		return nil, fmt.Errorf("post id not present in database %s", post.PostId)
	}

	other, err := deepCopy(post)
	if err != nil {
		return nil, err
	}

	store.data[other.PostId] = other
	return other, nil
}

func (store *InMemoryBlogStore) Read(postID string) (*pb.Post, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	post, ok := store.data[postID]
	if !ok {
		log.Printf("post id not present in database %s", postID)
		return nil, fmt.Errorf("post ID not present in database %s", postID)
	}

	return post, nil
}

func (store *InMemoryBlogStore) Delete(postID string) *pb.DeleteResponse {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	_, ok := store.data[postID]
	if !ok {
		log.Printf("post id not present in database %s", postID)
		response := &pb.DeleteResponse{
			Status: pb.DeleteResponse_ERROR,
			Error:  "post ID not present in database",
		}
		return response
	}

	delete(store.data, postID)
	response := &pb.DeleteResponse{
		Status: pb.DeleteResponse_SUCCESS,
	}

	return response
}

func deepCopy(post *pb.Post) (*pb.Post, error) {
	other := &pb.Post{}

	err := copier.Copy(other, post)
	if err != nil {
		return nil, fmt.Errorf("cannot copy post data: %w", err)
	}

	return other, nil
}
