package main

import (
	"fmt"

	"webapp/persistence/dao"
	"webapp/persistence/bolt"
	"webapp/entities"
)

func main() {
	fmt.Println("DAO testing ")

	fmt.Println("In memory DAO testing ")
	dao.SetDAOImplementation(dao.MEMORY)
	testDao(dao.GetStudentDAO())

	fmt.Println("Bolt DAO testing ")
	bolt.BoltStart("boltdb.data")
	dao.SetDAOImplementation(dao.BOLTDB)
	testDao(dao.GetStudentDAO())
}

func testDao(dao dao.StudentDAO) {

	if !dao.Exists(999) {
		fmt.Println("Student 999 not found => creation")
		student := entities.NewStudent()
		student.Id = 999
		student.FirstName = "Aaa"
		student.LastName = "Bbb"
		
		dao.Create(student)
	}

	student := dao.Find(999)

	if student == nil {
		fmt.Println("Student 999 not found")
	} else {
		fmt.Println("OK, Student found")
		fmt.Println(student.String())
	}
}
