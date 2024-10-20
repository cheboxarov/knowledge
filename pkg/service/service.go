package service

import "github.com/cheboxarov/knowledge/pkg/repository"

type Authorization interface {
}

type Article interface {
}

type ArticlesGroup interface {
}

type Service struct {
	Authorization
	Article
	ArticlesGroup
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
