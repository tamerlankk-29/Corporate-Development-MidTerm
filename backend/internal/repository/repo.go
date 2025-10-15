package repository

import "task-management-app/internal/model"

type TaskRepository interface {
	Create(task *model.Task) error
	GetAll() ([]*model.Task, error)
	GetByID(id int64) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int64) error
}
