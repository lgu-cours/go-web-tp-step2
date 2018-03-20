package bolt

import (
	"log"
	"encoding/json"
	
	"webapp/entities"
)

// This type/struct stores no state, it’s just a collection of methods
type LanguageDAOBolt struct {
	bucketName string 
}

func NewLanguageDAOBolt() LanguageDAOBolt {
	// Creates a DAO using the "language" bucket
	return LanguageDAOBolt{"language"}
}

// Converts the given JSON string to structure ( unmarshalling )
func (this *LanguageDAOBolt) jsonToStruct(value string) entities.Language {
	language := entities.Language{}
	err := json.Unmarshal([]byte(value), &language)
	if ( err != nil ) {
		panic(err)
	}	
	return language
}

// Converts the given structure to a JSON string ( marshalling )
func (this *LanguageDAOBolt) structToJson(language entities.Language) string {
	json, err := json.Marshal(language)
	if ( err != nil ) {
		panic(err)
	}
	// conversion : []byte to string 
	return string(json[:])
}

func (this *LanguageDAOBolt) FindAll() []entities.Language {
	log.Print("DAO - FindAll() ")
	languages := make([]entities.Language,0)
	values := dbGetAll(this.bucketName)
	for _, v := range values {
		language := this.jsonToStruct(v)
		languages = append(languages, language)
	}
	return languages
}

func (this *LanguageDAOBolt) Find(code string) *entities.Language {
	log.Printf("DAO - Find(%s) ", code)
	value := dbGet(this.bucketName, code)
	if value != "" {
		language := this.jsonToStruct(value)
		return &language
	} else {
		return nil 
	}
}

func (this *LanguageDAOBolt) Exists(code string) bool {
	log.Printf("DAO - Exists(%s) ", code)
	exists := false 
	if ( dbGet(this.bucketName, code) != "" ) {
		exists = true 
	}
	return exists
}

func (this *LanguageDAOBolt) Create(language entities.Language) bool {
	log.Printf("DAO - Create(%s) ", language.Code)
	if this.Exists(language.Code) {
		// already exists => cannot create !
		log.Printf("DAO - Create(%s) : already exists => cannot create", language.Code)
		return false
	} else {
		// not found => create
		log.Printf("DAO - Create(%s) : not found => created", language.Code)
		this.Save(language)
		return true
	}
}

func (this *LanguageDAOBolt) Delete(code string) bool {
	log.Printf("DAO - Delete(%s) ", code)
	if this.Exists(code) {
		dbDelete(this.bucketName, code) // delete in Bolt DB
		return true // found and deleted
	} else {
		return false // not found => not deleted
	}
}

func (this *LanguageDAOBolt) Update(language entities.Language) bool {
	log.Printf("DAO - Update(%s) ", language.Code)
	if this.Exists(language.Code) {
		this.Save(language) // update in Bolt DB
		return true // found and updated
	} else {
		return false // not found => not updated
	}
}

func (this *LanguageDAOBolt) Save(language entities.Language) {
	log.Printf("DAO - Save(%s) ", language.Code)
	key := language.Code
//	// value : language converted to JSON
//	value, err := json.Marshal(language)
//	if ( err != nil ) {
//		panic(err)
//	}
//	dbPut(this.bucketName, key, string(value[:]))
	
	value := this.structToJson(language)  
	dbPut(this.bucketName, key, value)
}
