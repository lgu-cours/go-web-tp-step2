package memdb

import (
	"log"
	"sync"

	"webapp/entities"
)

// Structure definition
type LanguageDataMap struct {
	dataMap map[string]entities.Language // the map to store Language entities
	lock    sync.RWMutex
}

var languageDataOnce sync.Once
var languageDataMap LanguageDataMap

func GetLanguageDataMap() *LanguageDataMap {
	log.Printf("LanguageDataMap - GetLanguageDataMap() ")
	languageDataOnce.Do(newLanguageDataMap)
	return &languageDataMap
}

func newLanguageDataMap() {
	log.Printf("LanguageDataMap - newLanguageDataMap() ***** ")
	languageDataMap = LanguageDataMap{
		dataMap: make(map[string]entities.Language),
		lock:    sync.RWMutex{},
	}
}

func (this *LanguageDataMap) read(code string) *entities.Language {
	log.Printf("LanguageDataMap - read(%s) ", code)
	this.lock.RLock()
	defer this.lock.RUnlock()
	language, exists := this.dataMap[code]
	if exists {
		return &language
	} else {
		return nil
	}
}
func (this *LanguageDataMap) exists(code string) bool {
	log.Printf("LanguageDataMap - exists(%s) ", code)
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, exists := this.dataMap[code]
	return exists
}

func (this *LanguageDataMap) write(language entities.Language) {
	log.Printf("LanguageDataMap - write(%s) ", language.String())
	this.lock.Lock()
	defer this.lock.Unlock()
	this.dataMap[language.Code] = language
}

func (this *LanguageDataMap) remove(code string) {
	log.Printf("LanguageDataMap - remove(%s) ", code)
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.dataMap, code) // delete in map
}

func (this *LanguageDataMap) values() []entities.Language {
	var a = make([]entities.Language, len(this.dataMap))
	i := 0
	for _, v := range this.dataMap {
		a[i] = v
		i++
	}
	return a
}
