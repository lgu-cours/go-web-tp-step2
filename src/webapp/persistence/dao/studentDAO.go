package dao

import (
	"webapp/entities"
)

// DAO interface for Student
type StudentDAO interface {
	FindAll() []entities.Student
	Find(id int) *entities.Student
	Exists(id int) bool
	Delete(id int) bool
	Create(student entities.Student) bool
	Update(student entities.Student) bool
}
