package services

import (
	"context"
	"pafaul/ritter/models"
	"pafaul/ritter/ritter_db"
)

func CreateUser(ctx context.Context, q *ritter_db.Queries, user *models.CreateUser) error {
	err := q.CreateUser(ctx, ritter_db.CreateUserParams{
		Nickname: user.Nickname,
		Passwd:   user.Passwd,
	})
	return err
}
