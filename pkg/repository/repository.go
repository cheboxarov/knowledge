package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Article interface {
}

type ArticlesGroup interface {
}

type Repository struct {
	Authorization
	Article
	ArticlesGroup
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
