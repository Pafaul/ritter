package models

type CreateUser struct {
	Nickname string `json:"nickname" validate:"required,min=8,max=100"`
	Passwd   string `json:"passwd" validate:"required"`
}
