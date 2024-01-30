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
