package controllers

import (
	"log"
	"net/http"
	"strconv"
	
	"webapp/persistence/dao"
	"webapp/entities"
	"webapp/webutil"
)

type StudentController struct {
	dao          dao.StudentDAO
	languagesDAO dao.LanguageDAO	
}

// Constructor : creates a new structure with fields initialization
func NewStudentController() StudentController {
	return StudentController {
		dao : dao.GetStudentDAO(),
		languagesDAO : dao.GetLanguageDAO(),
		}
}


func (this *StudentController) ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ListHandler - URL path '" + r.URL.Path )

	if r.Method == "GET" {
	    this.processList(w,r)
	} else {
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func (this *StudentController) FormHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("FormHandler - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    this.processForm(w,r)
	case "POST":
	    this.processPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func (this *StudentController) processList(w http.ResponseWriter, r *http.Request) {
	// get data
	data := this.dao.FindAll()
	// forward to view ( list page )
	webutil.Forward(w, "templates/studentList.gohtml", data)
}

func (this *StudentController) processForm(w http.ResponseWriter, r *http.Request) {
	// init form data
	// student := entities.Student{} // new Student with default values ( 'zero values' )
	student := entities.NewStudent()
	formData := this.NewStudentFormData(true, student)
	
	id := webutil.GetParameter(r, "id") 
	if  id != "" {
		i, _ := strconv.Atoi(id)
		student := this.dao.Find(i)
		if student != nil {
			formData.CreationMode = false
			formData.Student = *student
		}
	} 
	
	// forward to view ( form page )
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
}

func (this *StudentController) processPost(w http.ResponseWriter, r *http.Request) {
	log.Print("processPost " )
	
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    submit := r.Form.Get("submit")

	log.Print("processPost submit = " + submit )
    
    switch submit {
    	case "create":
	    	this.processCreate(w,r)
    	case "delete":
	    	this.processDelete(w,r)
    	case "update":
			this.processUpdate(w,r)
    	default:
	    	webutil.ErrorPage(w, "Unexpected action ")
    }
}

func (this *StudentController)  processCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("processCreate " )
    
    student := this.buildStudent(r)
	this.dao.Create(student) 

	formData := this.NewStudentFormData(false, student)
		
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
}

func (this *StudentController)  processDelete(w http.ResponseWriter, r *http.Request) {
	log.Print("processDelete " )
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    
    id, _ := strconv.Atoi( r.Form.Get("id") )
    
	log.Printf("Delete : id = %d", id )
	
	this.dao.Delete(id) 

	this.processList(w, r)
}

func (this *StudentController)  processUpdate(w http.ResponseWriter, r *http.Request) {
	log.Print("processUpdate " )
    student := this.buildStudent(r)
    
	this.dao.Update(student) 

	formData := this.NewStudentFormData(false, student)
	
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
}

func (this *StudentController)  buildStudent(r *http.Request) entities.Student {
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)

	log.Printf("buildStudent..." )
    
    student := entities.Student { 
    	Id: webutil.FormGetParamAsInt(r, "id", 0),
    	FirstName: r.Form.Get("firstname"), 
    	LastName: r.Form.Get("lastname"), 
    	Age: webutil.FormGetParamAsInt(r, "age", 0),
    	LanguageCode: r.Form.Get("languageCode") }
    
    //log.Printf("Student built : " + student.ToString() )
    log.Printf("Student built : " + student.String() )
	return student
}

func (this *StudentController) NewStudentFormData(creationMode bool, student entities.Student ) StudentFormData {
	// New structure
	var formData StudentFormData
	// Init structure fields
	formData.CreationMode = creationMode
	formData.Student      = student 
	formData.Languages    = this.getLanguages()  // The current list of languages
	// Return structure
	return formData
}

func (this *StudentController) getLanguages() []entities.Language {
	return this.languagesDAO.FindAll()
}