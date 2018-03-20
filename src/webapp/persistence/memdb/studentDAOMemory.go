package memdb

import (
	"log"
	"sort"

	"webapp/entities"
)

// Structure definition
type StudentDAOMemory struct {
	dataMap map[int]entities.Student // the map to store entities
}

// Structure pseudo-construtor
func NewStudentDAOMemory() StudentDAOMemory {
	log.Printf("NewStudentDAOMemory()")
	//dao := StudentDAOMemory{} // structure creation 
	//dao.init() // structure init
	//return dao
	return StudentDAOMemory{
		dataMap: make(map[int]entities.Student),
	} 
}

//// Structure initialization
//func (this *StudentDAOMemory) init() {
//	log.Printf("init()")
//	this.dataMap = make(map[int]entities.Student) // map creation
//}

func (this *StudentDAOMemory) values(m map[int]entities.Student) []entities.Student {
	var a = make([]entities.Student, len(m))
	i := 0
	for _, v := range m {
		a[i] = v
		i++
	}
	this.sortById(a)
	return a
}

func (this *StudentDAOMemory) sortById(students []entities.Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Id < students[j].Id
	})
}

func (this *StudentDAOMemory) FindAll() []entities.Student {
	log.Print("DAO - FindAll() ")
	return this.values(this.dataMap)
}

func (this *StudentDAOMemory) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id)
	student := this.dataMap[id]
	return &student
}

func (this *StudentDAOMemory) Exists(id int) bool {
	log.Printf("DAO - Exists(%d) ", id)
	_, exists := this.dataMap[id] // search in map
	log.Printf("DAO - Exists(%d) : ", id, exists)
	return exists
}

func (this *StudentDAOMemory) Create(student entities.Student) bool {
	log.Printf("DAO - Create(%d) ", student.Id)

	if this.Exists(student.Id) {
		// already exists => cannot create !
		log.Printf("DAO - Create(%d) : already exists => cannot create", student.Id)
		return false
	} else {
		// not found => create
		log.Printf("DAO - Create(%d) : not found => created", student.Id)
		this.dataMap[student.Id] = student
		return true
	}
}

func (this *StudentDAOMemory) Delete(id int) bool {
	log.Printf("DAO - Delete(%d) ", id)
	if this.Exists(id) {
		delete(this.dataMap, id) // delete in map
		return true              // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *StudentDAOMemory) Update(student entities.Student) bool {
	log.Printf("DAO - Update(%d) ", student.Id)
	if this.Exists(student.Id) {
		this.dataMap[student.Id] = student // update in map
		return true                        // found and updated
	} else {
		return false // not found => not updated
	}
}
