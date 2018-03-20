package controllers

import (
	"log"
	"net/http"
	
	"webapp/persistence/dao"
	"webapp/entities"
	"webapp/webutil"
)

type LanguageController struct {
	dao dao.LanguageDAO 
}

// Constructor : creates a new structure with fields initialization
func NewLanguageController() LanguageController {
	return LanguageController {
		dao : dao.GetLanguageDAO(),
		}
}

func (this *LanguageController) ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ListHandler - URL path '" + r.URL.Path )

	if r.Method == "GET" {
	    this.processList(w,r)
	} else {
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func (this *LanguageController) FormHandler(w http.ResponseWriter, r *http.Request) {
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

func (this *LanguageController) processList(w http.ResponseWriter, r *http.Request) {
	// get data
	data := this.dao.FindAll()
	// forward to view
	webutil.Forward(w, "templates/languageList.gohtml", data)
}

func (this *LanguageController) processForm(w http.ResponseWriter, r *http.Request) {
	// init form data
	language := entities.Language{} // new entity with default values ( 'zero values' )
	formData := this.newFormData(true, language)
	
	code := webutil.GetParameter(r, "code") 
	if  code != "" {
		language := this.dao.Find(code)
		if language != nil {
			formData.CreationMode = false
			formData.Language = *language
		}
	} 
	
	// forward to view ( form page )
	webutil.Forward(w, "templates/languageForm.gohtml", formData)
}

func (this *LanguageController) processPost(w http.ResponseWriter, r *http.Request) {
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

func (this *LanguageController)  processCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("processCreate " )
    
    language := this.buildLanguage(r)
	this.dao.Create(language) 
	formData := this.newFormData(false, language)
	webutil.Forward(w, "templates/languageForm.gohtml", formData)
}

func (this *LanguageController)  processDelete(w http.ResponseWriter, r *http.Request) {
	log.Print("processDelete " )
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    
//    id, _ := strconv.Atoi( r.Form.Get("id") )
//    
//	log.Printf("Delete : id = %d", id )
	
	code := r.Form.Get("code")
	this.dao.Delete(code) 
	this.processList(w, r)
}

func (this *LanguageController)  processUpdate(w http.ResponseWriter, r *http.Request) {
	log.Print("processUpdate " )
    language := this.buildLanguage(r)
	this.dao.Update(language) 
	formData := this.newFormData(false, language)
	webutil.Forward(w, "templates/languageForm.gohtml", formData)
}

func (this *LanguageController)  buildLanguage(r *http.Request) entities.Language {
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)

	log.Printf("buildLanguage..." )
    
    language := entities.Language { 
    	Code: r.Form.Get("code"), 
    	Name: r.Form.Get("name"), 
    	}
    
    log.Printf("Language built : " + language.String() )
	return language
}


func (this *LanguageController) newFormData(creationMode bool, language entities.Language ) LanguageFormData {
	// New structure
	var formData LanguageFormData
	// Init structure fields
	formData.CreationMode = creationMode
	formData.Language     = language 
	// Return structure
	return formData
}
