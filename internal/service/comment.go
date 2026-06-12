package service

import (
	"time"

	"github.com/lipkerton/wildcard/internal/domain"
)

type CommentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(body, author, password string) (*domain.Comment, error) {
	if body == "" {
		return nil, domain.ErrorEmptyBody
	}
	if password == "" {
		return nil, domain.ErrorEmptyPassword
	}
	comment := &domain.Comment{
		Body:     body,
		Author:   author,
		Password: password,
	}
	comment.CreatedAt = time.Now()

	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) GetById(id int) (*domain.Comment, error) {
	comment, err := s.repo.GetById(id)
	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return comment, nil
}

func (s *CommentService) Delete(id int, password string) error {
	comment, err := s.repo.GetById(id)
	if err != nil {
		return domain.ErrorNotFound
	}
	if comment.Password != password {
		return domain.ErrorWrongPassword
	}
	return s.repo.Delete(id)
}
