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

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	countryname := r.URL.Query().Get("country")
	Country := fonction.SearchCountry(countryname)
	if len(Country) == 0 {
		http.Error(w, "Not Found: No results found", http.StatusNotFound)
		return
	}

	InitTemplate.Temp.ExecuteTemplate(w, "details", Country)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "createuser", nil)
}

func CreateUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {

}

func AddFavoriteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")
	continent := r.URL.Query().Get("reg")
	area := r.URL.Query().Get("area")
	population := r.URL.Query().Get("pop")
	flag := r.URL.Query().Get("flag")
	isfavorite := true

	favorite := fonction.FavoriteInfo{
		Name:       countryName,
		Continent:  continent,
		Area:       area,
		Population: population,
		Flag:       flag,
		IsFavorite: isfavorite,
	}
	/*favs, err := fonction.RetrieveFavorite()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des favoris", http.StatusInternalServerError)
		return
	}
	for _, fav := range favs {
		if fav.Name == countryName && fav.IsFavorite {
			favorite.Message = "{{.Name}} est déjà dans vos favoris"
		}
	} */

	fonction.AddFav(favorite)
	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}

func FavoriteHandler(w http.ResponseWriter, r *http.Request) {
	favs, err := fonction.RetrieveFavorite()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des favoris", http.StatusInternalServerError)
		return
	}

	data := struct {
		Favorite []fonction.FavoriteInfo
	}{
		Favorite: favs,
	}
	InitTemplate.Temp.ExecuteTemplate(w, "favorite", data)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	countryName := r.FormValue("country")
	err := fonction.RemoveFavorite(countryName)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du pays favori", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}
