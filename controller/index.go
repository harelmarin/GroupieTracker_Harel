package controller

import (
	"Groupie/fonction"
	InitTemplate "Groupie/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "index", fonction.DecodeData)
}

func IndexTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

func CateHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("category") {
		return
	}
	Category := r.URL.Query().Get("category")
	searchcategory := fonction.GetContinent(Category)
	if len(searchcategory) == 0 {
		http.Error(w, "Not Found: No results found", http.StatusNotFound)
		return
	}
	InitTemplate.Temp.ExecuteTemplate(w, "category", searchcategory)
}

func RechercheHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("usersearch") {
		return
	}

	usersearch := r.URL.Query().Get("usersearch")
	searchResults := fonction.SearchCountry(usersearch)
	if len(searchResults) == 0 {
		http.Error(w, "Not Found: No results found", http.StatusNotFound)
		return
	}
	InitTemplate.Temp.ExecuteTemplate(w, "search", searchResults)
}
