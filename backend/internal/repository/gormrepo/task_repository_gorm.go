package gormrepo

import (
	"errors"

	"gorm.io/gorm"
	"task-management-app/internal/model"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *GormTaskRepository) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	if err := r.db.Order("id").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *GormTaskRepository) GetByID(id int64) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *GormTaskRepository) Update(task *model.Task) error {
	if err := r.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormTaskRepository) Delete(id int64) error {
	res := r.db.Delete(&model.Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
