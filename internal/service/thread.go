package service

import (
	"time"

	"github.com/lipkerton/wildcard/internal/domain"
)

type ThreadService struct {
	repo ThreadRepository
}

func NewThreadService(repo ThreadRepository) *ThreadService {
	return &ThreadService{repo: repo}
}

func (s *ThreadService) Create(subject, body, author, password string) (*domain.Thread, error) {
	if body == "" {
		return nil, domain.ErrorEmptyBody
	}
	if password == "" {
		return nil, domain.ErrorEmptyPassword
	}

	thread := &domain.Thread{
		Subject:  subject,
		Body:     body,
		Author:   author,
		Password: password,
	}
	now := time.Now()
	thread.CreatedAt = now
	thread.BumpedAt = now

	if err := s.repo.Create(thread); err != nil {
		return nil, err
	}
	return thread, nil
}

func (s *ThreadService) GetByID(id int) (*domain.Thread, error) {
	thread, err := s.repo.GetById(id)
	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return thread, nil
}

func (s *ThreadService) List() ([]domain.Thread, error) {
	threads, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (s *ThreadService) Delete(id int, password string) error {
	thread, err := s.repo.GetById(id)
	if err != nil {
		return domain.ErrorNotFound
	}
	if thread.Password != password {
		return domain.ErrorWrongPassword
	}
	return s.repo.Delete(id)
}
