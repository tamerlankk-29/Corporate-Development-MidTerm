package rawrepo

import (
	"database/sql"
	"errors"
	"time"
	
	"task-management-app/internal/model"
)

type RawTaskRepository struct {
	db *sql.DB
}

func NewRawTaskRepository(db *sql.DB) *RawTaskRepository {
	return &RawTaskRepository{db: db}
}

func (r *RawTaskRepository) Create(task *model.Task) error {
	now := time.Now().UTC()
	var description sql.NullString
	if task.Description != nil {
		description = sql.NullString{String: *task.Description, Valid: true}
	}
	query := `INSERT INTO tasks (title, description, status, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	row := r.db.QueryRow(query, task.Title, description, task.Status, now, now)
	if err := row.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (r *RawTaskRepository) GetAll() ([]*model.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var t model.Task
		var description sql.NullString
		if err := rows.Scan(&t.ID, &t.Title, &description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		if description.Valid {
			t.Description = &description.String
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (r *RawTaskRepository) GetByID(id int64) (*model.Task, error) {
	var t model.Task
	var description sql.NullString
	err := r.db.QueryRow(`SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id=$1`, id).
		Scan(&t.ID, &t.Title, &description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if description.Valid {
		t.Description = &description.String
	}
	return &t, nil
}

func (r *RawTaskRepository) Update(task *model.Task) error {
	task.UpdatedAt = time.Now().UTC()
	var description interface{}
	if task.Description == nil {
		description = nil
	} else {
		description = *task.Description
	}
	res, err := r.db.Exec(`UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5`,
		task.Title, description, task.Status, task.UpdatedAt, task.ID)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *RawTaskRepository) Delete(id int64) error {
	res, err := r.db.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return sql.ErrNoRows
	}
	return nil
}
