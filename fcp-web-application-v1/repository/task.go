package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	err := t.db.Model(&task).Where("id = ?", task.ID).Updates(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	err := t.db.Where("id = ?", id).Delete(&model.Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	list := []model.Task{}
	err := t.db.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	list := []model.TaskCategory{}
	err := t.db.Raw("SELECT t.id AS ID, t.title AS Title, c.name AS category FROM tasks t, categories c WHERE c.id = t.id and c.id = ?", id).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
