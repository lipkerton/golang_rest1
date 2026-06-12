package service

import "github.com/lipkerton/wildcard/internal/domain"

type ThreadRepository interface {
	Create(thread *domain.Thread) error
	GetById(id int) (*domain.Thread, error)
	List() ([]domain.Thread, error)
	Delete(id int) error
	UpdateBumpedAt(id int) error
}

type CommentRepository interface {
	Create(comment *domain.Comment) error
	GetById(id int) (*domain.Comment, error)
	Delete(id int) error
}
