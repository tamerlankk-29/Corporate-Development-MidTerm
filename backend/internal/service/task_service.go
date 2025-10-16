package service

import (
	"errors"

	"task-management-app/internal/model"
	"task-management-app/internal/repository"
)

var ErrNotFound = errors.New("task not found")

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *model.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.Status == "" {
		task.Status = "pending"
	}
	return s.repo.Create(task)
}

func (s *TaskService) ListTasks() ([]*model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTask(id int64) (*model.Task, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNotFound
	}
	return t, nil
}

func (s *TaskService) UpdateTask(task *model.Task) error {
	existing, err := s.repo.GetByID(task.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrNotFound
	}
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int64) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrNotFound
	}
	return s.repo.Delete(id)
}
