package entities

import "fmt"

type Student struct {
	Id        int 
	FirstName string
	LastName  string
	Age       int 
	LanguageCode string
}

func NewStudent() Student {
	// new Student with default values ( 'zero values' )
	return Student{}
}

func (this Student) String() string {
    return fmt.Sprintf(
    	"[%d : %s, %s, %d, %s]", 
    	this.Id, 
    	this.FirstName, 
    	this.LastName, 
    	this.Age, 
    	this.LanguageCode) 
}