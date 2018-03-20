package bolt

import (
	"log"
	
	"github.com/boltdb/bolt"
)

//-----------------------------------------------------------------------------------
// Exported / Public
//-----------------------------------------------------------------------------------

func BoltStart(fileName string) {
	log.Printf("BoltStart - fileName = %s ", fileName)

	if db != nil {
		panic("Bolt database is already started !")
	} else {
		dbOpen(fileName)
	}
	log.Printf("BoltStart : OK, Bolt is ready.")
}

func BoltStop() {
	dbClose()
}

//-----------------------------------------------------------------------------------
// Unexported / Private
//-----------------------------------------------------------------------------------

// The current Bolt Database
var db *bolt.DB

func dbOpen(fileName string) {
	// Open the Bolt data file in the current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open(fileName, 0600, nil)
	if err != nil {
		panic("Cannot open database. File '" + fileName + "' ")
	}	
}

func dbCheckOpen() {
	if db == nil {
		panic("Bolt database is not open!")
	}
}

// Checks if all the preconditions are OK to proceed a Bolt operation
func dbCheckPrecondition(bucketName string) {
	if db == nil {
		panic("Bolt database is not open!")
	}
	if bucketName != "" {
		dbCreateBucketIfNotExists(bucketName)
	}
}

func dbClose() {
	if db != nil {
		db.Close()
	}
}

func dbPath() string {
	dbCheckPrecondition("")
	return db.Path()
}

func dbCreateBucketIfNotExists(bucketName string) {
	err := db.Update(func(tx *bolt.Tx) error {
	    //_, err := tx.CreateBucket( []byte(bucketName) ) 
	    _, err := tx.CreateBucketIfNotExists( []byte(bucketName) )       
	    if err != nil {
	        return err
	    }
	    return nil
	});	
	
	if err != nil {
		panic("Cannot create bucket '" + bucketName + "' ")
	}
}

func dbPut(bucketName string, key string, value string ) {
	dbCheckPrecondition(bucketName)
	
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    bucket.Put([]byte(key), []byte(value))                        
	    return nil
	});	
	
	if err != nil {
		panic("Cannot put key '" + key + "' in bucket '" + bucketName + "'")
	}
}

// Returns a string containing the value found for the given KEY in the given BUCKET
// If not found returns a void string 
func dbGet(bucketName string, key string ) string {
	dbCheckPrecondition(bucketName)
	
	var value string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    value = string(bucket.Get([]byte(key))) // void string if value == nil (not found)      
	    return nil
	});	
	
	
	if err != nil {
		panic("Cannot get key '" + key + "' from bucket '" + bucketName + "'")
	}
	return value
}

// Get all the values stored in the given bucket name
func dbGetAll(bucketName string) []string {
	dbCheckPrecondition(bucketName)
	
	values := make([]string,0)
	
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
		bucket.ForEach(func(k, v []byte) error {
	        //fmt.Println(string(k), string(v))
		    values = append(values, string(v))
	        return nil
		})
		return nil
	});
	
	if err != nil {
		panic("Cannot get all items from bucket '" + bucketName + "'")
	}
	return values
}

func dbDelete(bucketName string, key string ) {
	dbCheckPrecondition(bucketName)
	
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if ( bucket == nil ) {
			panic("Bucket '" + bucketName + "' not found")
		}
	    bucket.Delete([]byte(key))                                                     
	    return nil
	});	
	
	if err != nil {
		panic("Cannot delete key '" + key + "' in bucket '" + bucketName + "'")
	}
}

