package controller

import (
	"Groupie/fonction"
	InitTemplate "Groupie/templates"
	"net/http"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "index", fonction.DecodeData)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

func TreatmentCaté(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

func DataRecherche(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
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
