package main

import (
	"net/http"
	
	"webapp/controllers"  // keep full path for "go build"
)

func initControllers() {

	// Specific Paths with specific controllers 

	languageController := controllers.NewLanguageController() 
	http.HandleFunc("/language/list", languageController.ListHandler)
	http.HandleFunc("/language/form", languageController.FormHandler)

	studentController := controllers.NewStudentController() 
	http.HandleFunc("/student/list", studentController.ListHandler )
	http.HandleFunc("/student/form", studentController.FormHandler )

}