package repository

import (
	"github.com/bhawneshdipu/go-angular-sample-app/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	DB *gorm.DB
}

func (userRepo *UserRepository) Init() {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to open DB Connection")
	}
	userRepo.DB = db
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.UserTask{})
}
func (userRepo *UserRepository) FindAll() []model.User {
	users := []model.User{}
	userRepo.DB.Find(&users)
	return users
}
func (userRepo *UserRepository) FindById(id int) model.User {
	user := model.User{}
	userRepo.DB.Find(&user, "id = ?", id)
	return user
}
func (userRepo *UserRepository) Create(user *model.User) *model.User {
	userRepo.DB.Create(user)
	return user
}
