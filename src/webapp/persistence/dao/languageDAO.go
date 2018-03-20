package dao

import (
	"webapp/entities"
)

// DAO interface for Language
type LanguageDAO interface {
	FindAll() []entities.Language
	Find(code string) *entities.Language
	Exists(code string) bool
	Delete(code string) bool
	Create(language entities.Language) bool
	Update(language entities.Language) bool
}
