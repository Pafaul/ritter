package controllers

import "pafaul/ritter/ritter_db"

type (
	Handler struct {
		q *ritter_db.Queries
	}
)

func NewHandler(q *ritter_db.Queries) *Handler {
	handler := new(Handler)
	handler.q = q
	return handler
}
