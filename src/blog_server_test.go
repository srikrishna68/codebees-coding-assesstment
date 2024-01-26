package src

import (
	`context`
	`testing`

	`codebees-coding-assesstment/pb`
	`codebees-coding-assesstment/sample`

	`github.com/stretchr/testify/assert`
	`github.com/stretchr/testify/require`
	`google.golang.org/grpc/codes`
)

func TestBlogCreatePost(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		post  *pb.Post
		store BlogStore
		code  codes.Code
	}{
		{
			name:  "success_with_id",
			post:  sample.NewPost(),
			store: NewInMemoryBlogStore(),
			code:  codes.OK,
		},
		{
			name:  "failure_no_title",
			post:  sample.NewInvalidPostNoTitle(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "failure_no_author",
			post:  sample.NewInvalidPostNoAuthor(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "failure_no_tags",
			post:  sample.NewInvalidPostNoTags(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "failure_no_content",
			post:  sample.NewInvalidPostNoContent(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "failure_no_publication_date",
			post:  sample.NewInvalidPostNoPublicationDate(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := NewBlogServer(tc.store)
			res, err := server.CreatePost(context.Background(), tc.post)
			if tc.code == codes.OK {
				assert.Nil(t, err)
				assert.NotNil(t, res)
				assert.NotEmpty(t, res.Post.PostId)
			} else {
				assert.NotNil(t, err)
				require.Nil(t, res)
			}
		})

	}
}

func TestBlogUpdatePost(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		post  *pb.Post
		store BlogStore
		code  codes.Code
	}{
		{
			name:  "success",
			post:  sample.NewUpdatePost(),
			store: NewInMemoryBlogStore(),
			code:  codes.OK,
		},
		{
			name:  "success_publication_should not be added",
			post:  sample.NewPost(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "success_no_title",
			post:  sample.NewInvalidPostNoTitle(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "success_no_author",
			post:  sample.NewInvalidPostNoAuthor(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "success_no_tags",
			post:  sample.NewInvalidPostNoTags(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
		{
			name:  "success_no_tags",
			post:  sample.NewInvalidPostNoTags(),
			store: NewInMemoryBlogStore(),
			code:  codes.Unavailable,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := NewBlogServer(tc.store)
			createResponse, err := server.CreatePost(context.Background(), sample.NewPost())
			assert.Nil(t, err)
			assert.NotNil(t, createResponse)
			readBeforeUpdate, err := server.ReadPost(context.Background(), &pb.PostID{Id: createResponse.Post.PostId})
			assert.Nil(t, err)
			assert.Equal(t, sample.NewPost().PublicationDate, readBeforeUpdate.Post.PublicationDate)

			tc.post.PostId = createResponse.Post.PostId
			updateResponse, err := server.UpdatePost(context.Background(), tc.post)
			if tc.code == codes.OK {
				assert.Nil(t, err, "error should be empty")
				assert.Nil(t, err)
				assert.NotEqualValues(t, createResponse.Post.Title, updateResponse.Post.Title)

				readAfterUpdate, err := server.ReadPost(context.Background(), &pb.PostID{Id: readBeforeUpdate.Post.PostId})
				assert.Nil(t, err)
				assert.Equal(t, tc.post.Title, readAfterUpdate.Post.Title)

			} else {
				assert.NotNil(t, err)
				require.Nil(t, updateResponse)
			}
		})

	}
}
