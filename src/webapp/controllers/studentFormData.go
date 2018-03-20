package controllers

import (
	"webapp/entities"
)

type StudentFormData struct {
	CreationMode  bool
    Student       entities.Student 
    Languages     []entities.Language
}

