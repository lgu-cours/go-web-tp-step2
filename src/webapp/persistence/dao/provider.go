package dao

import (
	"webapp/persistence/bolt"
	"webapp/persistence/memdb"
)

const MEMORY = 1
const BOLTDB = 2

var implementation = MEMORY

func SetDAOImplementation(implem int) {
	if implem == MEMORY || implem == BOLTDB {
		implementation = implem
	} else {
		panic("Cannot set DAO implementation : invalid implementation !")
	}
}

//--------------------------------------------------------------------------
// StudentDAO
//--------------------------------------------------------------------------
var studentDAOBolt   = bolt.NewStudentDAOBolt()
var studentDAOMemory = memdb.NewStudentDAOMemory()

func GetStudentDAO() StudentDAO {

	var dao StudentDAO = &studentDAOMemory
	if implementation == BOLTDB {
		dao = &studentDAOBolt
	}
	return dao
}

//--------------------------------------------------------------------------
// LanguageDAO
//--------------------------------------------------------------------------
var languageDAOBolt   = bolt.NewLanguageDAOBolt()
var languageDAOMemory = memdb.NewLanguageDAOMemory()

func GetLanguageDAO() LanguageDAO {
	if implementation == BOLTDB {
		return &languageDAOBolt
	}
	return &languageDAOMemory // Default implementation
}
