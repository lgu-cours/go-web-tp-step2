package memdb

import (
	"log"

	"webapp/entities"
)

//-----------------------------------------------------
// Langage DAO with "in memory data" (based on a map)
//-----------------------------------------------------

// Structure definition
type LanguageDAOMemory struct {
	// dataMap * map[string]entities.Language // the map to store Language entities
	dataMap * LanguageDataMap
}

// Structure pseudo-construtor
func NewLanguageDAOMemory() LanguageDAOMemory {
	log.Printf("NewLanguageDAOMemory()")
	// dao := LanguageDAOMemory{}
	// dao.init()
	// return dao
	return LanguageDAOMemory{
		//dataMap: make(map[string]entities.Language),
		dataMap: GetLanguageDataMap(),
	}
}

// Structure initialization
//func (this *LanguageDAOMemory) init() {
//	log.Printf("init()")
//	this.dataMap = make(map[string]entities.Language)
//}

//func (this *LanguageDAOMemory) values(m map[string]entities.Language) []entities.Language {
//	var a = make([]entities.Language, len(m))
//	i := 0
//	for _, v := range m {
//		a[i] = v
//		i++
//	}
//	return a
//}

func (this *LanguageDAOMemory) FindAll() []entities.Language {
	log.Print("LanguageDAOMemory - FindAll() ")
	//return this.values(this.dataMap)
	return this.dataMap.values()
}

func (this *LanguageDAOMemory) Find(code string) *entities.Language {
	log.Printf("LanguageDAOMemory - Find(%s) ", code)
//	language := this.dataMap[code]
//	return &language
	return this.dataMap.read(code)
}

func (this *LanguageDAOMemory) Exists(code string) bool {
	log.Printf("LanguageDAOMemory - Exists(%s) ", code)
	// _, exists := this.dataMap[code] // search in map
	exists := this.dataMap.exists(code)
	log.Printf("LanguageDAOMemory - Exists(%s, %t) : ", code, exists)
	return exists
}

func (this *LanguageDAOMemory) Create(language entities.Language) bool {
	log.Printf("LanguageDAOMemory - Create(%s) ", language.Code)

	if this.Exists(language.Code) {
		log.Printf("LanguageDAOMemory - Create(%s) : already exists => cannot create", language.Code)
		return false
	} else {
		log.Printf("LanguageDAOMemory - Create(%s) : not found => created", language.Code)
		//this.dataMap[language.Code] = language
		this.dataMap.write(language)
		return true
	}
}

func (this *LanguageDAOMemory) Delete(code string) bool {
	log.Printf("LanguageDAOMemory - Delete(%s) ", code)
	if this.Exists(code) {
		//delete(this.dataMap, code) // delete in map
		this.dataMap.remove(code)
		return true                // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *LanguageDAOMemory) Update(language entities.Language) bool {
	log.Printf("LanguageDAOMemory - Update(%s) ", language.Code)
	if this.Exists(language.Code) {
		//this.dataMap[language.Code] = language // update in map
		this.dataMap.write(language)
		return true                            // found and updated
	} else {
		return false // not found => not updated
	}
}
