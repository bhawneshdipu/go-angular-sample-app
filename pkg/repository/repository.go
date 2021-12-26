package repository

import (
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/model"
)

type Repository interface {
	Init()
	FindAll()
	FindById(id int)
	Create(user *model.User)
}
