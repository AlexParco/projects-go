package model

import (
	"github.com/alexparco/rest-api-todo/database"
)

type TaskRepo interface {
	Get() ([]*Task, error)
	GetOne(id string) (*Task, error)
	Create(taks *Task) error
	Update(task *Task) error
	Delete(id string) (*Task, error)
}

type taskRepo struct {
	*database.DBClient
}

func NewTaskRepo(db *database.DBClient) TaskRepo {
	return &taskRepo{db}
}

func (t *taskRepo) Get() ([]*Task, error) {
	var tasks []*Task

	rows, err := t.Query("SELECT id, title, IF(status, 'true', 'false') FROM todo")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.Id, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *taskRepo) GetOne(id string) (*Task, error) {
	var task Task
	stmt, err := t.Prepare("SELECT id, title, IF(status, 'true', 'false') FROM todo WHERE id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&task.Id, &task.Title, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepo) Create(task *Task) error {
	stmt, err := t.Prepare("INSERT INTO todo (title, status, updated_at) VALUES(?, ?, current_timestamp())")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(task.Title, task.Status)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	task.Id = int(id)

	return nil
}

func (t *taskRepo) Update(task *Task) error {
	stmt, err := t.Prepare("UPDATE todo.todo SET title=?, status=?, updated_at=current_timestamp() WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Title, task.Status, task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepo) Delete(id string) (*Task, error) {
	task, err := t.GetOne(id)
	if err != nil {
		return nil, err
	}

	stmt, err := t.Prepare("DELETE FROM todo WHERE id=?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
