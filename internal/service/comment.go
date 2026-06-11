package service

import (
	"errors"
	"time"

	"github.com/lipkerton/wildcard/internal/domain"
)

var (
	ErrorEmptyBody := errors.New("empty body")
	ErrorEmptyPassword := errors.New("empty password")
	ErrorNotFound := errors.New("comment not found")
	ErrorWrongPassword := errors.New("wrong password")
)

type CommentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(body, author, password string) (*domain.Comment, error) {
	if body == "" {
		return nil, ErrorEmptyBody
	}
	if password == "" {
		return nil, ErrorEmptyPassword
	}
	comment := &domain.Comment{
		Body: body,
		Author:	author,
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
		return nil, ErrorNotFound
	}
	return comment, nil
}

func (s *CommentService) Delete(id int, password string) error {
	comment, err := s.repo.GetById(id)
	if err != nil {
		return nil, ErrorNotFound
	}
	if comment.Password != password {
		return nil, ErrorWrongPassword
	}
	return s.repo.Delete(id)
}
