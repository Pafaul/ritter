package models

type Post struct {
	UserId  int64  `json:"userId" validate:"required"`
	Content string `json:"content" validate:"required,min=1,max=140"`
}

type GetPosts struct {
	UserId int64 `json:"userId" validate:"required"`
	Limit  int64 `json:"limit" validate:"min=1,max=100"`
	Offset int64 `json:"offset" validate:"min=0"`
}
