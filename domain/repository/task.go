package repository

import "github.com/paypay3/go-echo-gorm-sample/domain/model"

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}
