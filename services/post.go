package services

import (
	"context"
	"pafaul/ritter/models"
	"pafaul/ritter/ritter_db"
)

func CreatePost(ctx context.Context, q *ritter_db.Queries, post *models.Post) error {
	err := q.CreatePost(ctx, ritter_db.CreatePostParams{
		Content: post.Content,
		UserID:  post.UserId,
	})
	return err
}

func GetPosts(ctx context.Context, q *ritter_db.Queries, getPosts *models.GetPosts) ([]ritter_db.GetPostsRow, error) {
	posts, err := q.GetPosts(ctx, ritter_db.GetPostsParams{
		Limit:  getPosts.Limit,
		Offset: getPosts.Offset,
	})
	return posts, err
}
