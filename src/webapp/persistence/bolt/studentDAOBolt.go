package bolt

import (
	"log"
	"strconv"
	"encoding/json"
	
	"webapp/entities"
)

// This type/struct stores no state, it’s just a collection of methods
type StudentDAOBolt struct {
	bucketName string 
}

func NewStudentDAOBolt() StudentDAOBolt {
	// Creates a DAO using the "students" bucket
	return StudentDAOBolt{"student"}
}

func (this *StudentDAOBolt) toStudent(value string) entities.Student {
	student := entities.Student{}
	err := json.Unmarshal([]byte(value), &student)
	if ( err != nil ) {
		panic(err)
	}	
	return student
}

func (this *StudentDAOBolt) FindAll() []entities.Student {
	log.Print("DAO - FindAll() ")
	students := make([]entities.Student,0)
	values := dbGetAll(this.bucketName)
	for _, v := range values {
		student := this.toStudent(v)
		students = append(students, student)
	}
	return students
}

func (this *StudentDAOBolt) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id)
	key := strconv.Itoa(id)
	value := dbGet(this.bucketName, key)
	if value != "" {
		student := this.toStudent(value)
		return &student
	} else {
		return nil 
	}
}

func (this *StudentDAOBolt) Exists(id int) bool {
	log.Printf("DAO - Exists(%d) ", id)
	exists := false 
	key := strconv.Itoa(id)
	if ( dbGet(this.bucketName, key) != "" ) {
		exists = true 
	}
	return exists
}

func (this *StudentDAOBolt) Create(student entities.Student) bool {
	log.Printf("DAO - Create(%d) ", student.Id)
	if this.Exists(student.Id) {
		// already exists => cannot create !
		log.Printf("DAO - Create(%d) : already exists => cannot create", student.Id)
		return false
	} else {
		// not found => create
		log.Printf("DAO - Create(%d) : not found => created", student.Id)
		this.Save(student)
		return true
	}
}

func (this *StudentDAOBolt) Delete(id int) bool {
	log.Printf("DAO - Delete(%d) ", id)
	if this.Exists(id) {
		key := strconv.Itoa(id)
		dbDelete(this.bucketName, key) // delete in Bolt DB
		return true // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *StudentDAOBolt) Update(student entities.Student) bool {
	log.Printf("DAO - Update(%d) ", student.Id)
	if this.Exists(student.Id) {
		this.Save(student) // update in Bolt DB
		return true // found and updated
	} else {
		return false // not found => not updated
	}
}

func (this *StudentDAOBolt) Save(student entities.Student) {
	log.Printf("DAO - Save(%d) ", student.Id)
	key := strconv.Itoa(student.Id)
	// value : student converted to JSON
	value, err := json.Marshal(student)
	if ( err != nil ) {
		panic(err)
	}
	dbPut(this.bucketName, key, string(value[:]))
}
