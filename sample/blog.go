package sample

import (
	`time`

	`codebees-coding-assesstment/pb`
)

// newPost returns a new sample post
func NewPost() *pb.Post {

	post := &pb.Post{
		Title:           "First Title",
		Content:         "First Title",
		Author:          "First Author",
		PublicationDate: time.DateOnly,
		Tags:            []string{"abc"},
	}

	return post
}

func NewInvalidPostNoTitle() *pb.Post {

	post := &pb.Post{
		Content:         "First Title",
		Author:          "First Author",
		PublicationDate: time.DateOnly,
		Tags:            []string{"abc"},
	}

	return post
}

func NewInvalidPostNoContent() *pb.Post {

	post := &pb.Post{
		Title:           "First Title",
		Author:          "First Author",
		PublicationDate: time.DateOnly,
		Tags:            []string{"abc"},
	}

	return post
}

func NewInvalidPostNoAuthor() *pb.Post {

	post := &pb.Post{
		Title:           "First Title",
		Content:         "First Title",
		PublicationDate: time.DateOnly,
		Tags:            []string{"abc"},
	}

	return post
}

func NewInvalidPostNoPublicationDate() *pb.Post {

	post := &pb.Post{
		Title:   "First Title",
		Content: "First Title",
		Author:  "First Author",
		Tags:    []string{"abc"},
	}

	return post
}

func NewInvalidPostNoTags() *pb.Post {

	post := &pb.Post{
		Title:           "First Title",
		Content:         "First Title",
		Author:          "First Author",
		PublicationDate: time.DateOnly,
	}

	return post
}

func NewUpdatePost() *pb.Post {
	post := &pb.Post{
		Title:   "Update Title",
		Content: "First Title",
		Author:  "First Author",
		Tags:    []string{"abc"},
	}

	return post
}
