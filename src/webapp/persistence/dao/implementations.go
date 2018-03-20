package dao

import (
	"webapp/persistence/bolt"
	"webapp/persistence/memdb"
)

// Check all XxxxDAOMemory implements XxxDAO interface
var _  StudentDAO  = (*memdb.StudentDAOMemory)(nil)
var _  LanguageDAO = (*memdb.LanguageDAOMemory)(nil)


// Check all XxxxDAOBolt implements XxxDAO interface
var _  StudentDAO  = (*bolt.StudentDAOBolt)(nil)
var _  LanguageDAO = (*bolt.LanguageDAOBolt)(nil)

