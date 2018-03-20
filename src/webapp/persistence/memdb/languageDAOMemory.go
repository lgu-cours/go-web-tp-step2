package memdb

import (
	"log"
	
	"webapp/entities"
)

// Langage DAO with "in memory data" (based on a map)
type LanguageDAOMemory struct {
	dataMap map[string]entities.Language  // the map to store Language entities
}

// Structure pseudo-construtor 
func NewLanguageDAOMemory() LanguageDAOMemory {
	log.Printf("NewLanguageDAOMemory()")
	dao := LanguageDAOMemory{}
	dao.init()
	return dao
}

// Structure initialization 
func (this *LanguageDAOMemory) init() {
	log.Printf("init()")
	this.dataMap = make(map[string]entities.Language)
}

func (this *LanguageDAOMemory) values( m map[string]entities.Language ) []entities.Language  {
	var a = make([]entities.Language, len(m))
	i := 0
    for _, v := range m {
	    a[i] = v
    	i++
    }
	return a
}

func (this *LanguageDAOMemory) FindAll() []entities.Language {
	log.Print("DAO - FindAll() " )
	return this.values(this.dataMap)
}

func (this *LanguageDAOMemory) Find(code string) *entities.Language {
	log.Printf("DAO - Find(%s) ", code )
	language := this.dataMap[code]
	return &language
}

func (this *LanguageDAOMemory) Exists(code string) bool {
	log.Printf("DAO - Exists(%s) ", code)
	_, exists := this.dataMap[code] // search in map
	log.Printf("DAO - Exists(%s) : ", code, exists)
	return exists
}

func (this *LanguageDAOMemory) Create(language entities.Language) bool {
	log.Printf("DAO - Create(%s) ", language.Code)
	
	if this.Exists(language.Code) {
		log.Printf("DAO - Create(%s) : already exists => cannot create", language.Code)
		return false
	} else {
		log.Printf("DAO - Create(%s) : not found => created", language.Code)
		this.dataMap[language.Code] = language
		return true
	}
}

func (this *LanguageDAOMemory) Delete(code string) bool {
	log.Printf("DAO - Delete(%s) ", code)
	if this.Exists(code) {
		delete(this.dataMap, code) // delete in map
		return true // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *LanguageDAOMemory) Update(language entities.Language) bool {
	log.Printf("DAO - Update(%s) ", language.Code)
	if this.Exists(language.Code) {
		this.dataMap[language.Code] = language // update in map
		return true // found and updated
	} else {
		return false // not found => not updated
	}
}
