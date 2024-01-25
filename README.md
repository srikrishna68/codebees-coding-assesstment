# codebees-coding-assesstment

**Application to be coded:**



**Background:** All API referenced are gRPC APIs, not REST ones.



**Objective:** Develop a gRPC-based API in Golang to manage blog posts for a hypothetical blogging platform. The API should support CRUD operations for blog posts, and each post should have the following attributes:

PostID (unique identifier)
Title
Content
Author
Publication Date
Tags (multiple tags per post)
The API should support the following operations:



**Create Post**

Input: Post details (Title, Content, Author, Publication Date, Tags)

Output: The Post (PostID, Title, Content, Author, Publication Date, Tags). Error message, if creation fails.



**Read Post**

Input: PostID of the post to retrieve

Output: Post details (PostID, Title, Content, Author, Publication Date, Tags) or an error message if the post is not found.



**Update Post**

Input: PostID of the post to update and new details (Title, Content, Author, Tags)

Output: Post details (PostID, Title, Content, Author, Publication Date, Tags) or error message if the update failed



**Delete Post**

Input: PostID of the post to delete

Output: Success/Failure message



**src** - has the server and client methods on createPost, UpdatePost, ReadPost and DeletePost

**Server main** present in grpc/server/main.go
![image](https://github.com/srikrishna68/codebees-coding-assesstment/assets/23114891/feb1ef12-8d3e-4dd5-8ca0-04b248eaf05b)

**Client main** present in grpc/client/main.go
![image](https://github.com/srikrishna68/codebees-coding-assesstment/assets/23114891/2a2d47dd-11e5-4279-8f5f-5401a4e1254f)

