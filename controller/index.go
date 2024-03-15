package controller

import (
	"Groupie/fonction"
	InitTemplate "Groupie/templates"
	"net/http"
	"strconv"
)

// Handler pour afficher l'index du site
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	searchlangue := fonction.SearchIndex()
	independent := r.URL.Query().Get("independent")
	alphabetical := r.URL.Query().Get("alphabetical")
	minPopulation := r.URL.Query().Get("min_population")
	maxPopulation := r.URL.Query().Get("max_population")

	var message string

	// Si independant est coché j'append les résultats correspondant par rapport à la recherche d'avant
	if independent == "true" {
		searchlangue = fonction.FilterIndependent(searchlangue)
	}
	// si A-Z est coché je garde les mêmes resultats mais en changeant l'ordre
	if alphabetical == "true" {
		searchlangue = fonction.FilterAlphabetical(searchlangue)
	}
	// Si un Min ou Max sont spécifiés je prends les résultats correspondant dans les résultats qu'ils me restent des autres filtres
	if minPopulation != "" || maxPopulation != "" {
		minPop, _ := strconv.Atoi(minPopulation)
		maxPop, _ := strconv.Atoi(maxPopulation)
		searchlangue = fonction.FilterByPopulation(searchlangue, minPop, maxPop)
	}

	// Message si les résultats = 0
	if len(searchlangue) == 0 {
		message = "Aucun résultat pour votre recherche"
	} else {
		message = ""
	}

	data := struct {
		Message string
		Results []fonction.SearchResults
	}{

		Message: message,
		Results: searchlangue,
	}
	InitTemplate.Temp.ExecuteTemplate(w, "index", data)
}

func IndexTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

// Handler pour afficher les résultats de la catégorie utilisée
func CateHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("category") {
		return
	}
	// Je récupère les données de la query dans mon handler pour les utiliser
	Category := r.URL.Query().Get("category")

	searchcategory := fonction.GetContinent(Category)

	// Calcule le nombre de page par rapport au nombre de résultat
	pageNumber := 1
	pageNumberStr := r.URL.Query().Get("page")
	if pageNumberStr != "" {
		pageNumber, _ = strconv.Atoi(pageNumberStr)
		if pageNumber < 1 {
			pageNumber = 1
		}
	}

	totalResults := len(searchcategory)
	nbPages := totalResults / 10
	if totalResults%10 != 0 {
		nbPages++
	}
	startIndex := (pageNumber - 1) * 10
	endIndex := startIndex + 10
	if endIndex > totalResults {
		endIndex = totalResults
	}
	resultsPage := searchcategory[startIndex:endIndex]

	// Si il n'y a aucun résultat j'affiche un message
	var message string
	if len(searchcategory) == 0 {
		message = "Aucun résultat pour votre recherche"
	} else {
		message = ""
	}

	Next := pageNumber + 1
	Prev := pageNumber - 1

	// Structure de Data que je vais envoyer
	data := struct {
		Category    string
		Results     []fonction.SearchResults
		Message     string
		CurrentPage int
		TotalPages  int
		Next        int
		Prev        int
	}{
		Category:    Category,
		Results:     resultsPage,
		Message:     message,
		CurrentPage: pageNumber,
		TotalPages:  nbPages,
		Next:        Next,
		Prev:        Prev,
	}

	// Init avec les data
	InitTemplate.Temp.ExecuteTemplate(w, "category", data)
}

// Handler pour afficher les résultats de la recherche utilisateur de la searchBar sur une page.
func RechercheHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("usersearch") {
		return
	}
	// Je récupère les données de la query dans mon handler pour les utiliser
	usersearch := r.URL.Query().Get("usersearch")

	searchResults := fonction.SearchCountry(usersearch)
	var message string

	// Message si les résultats = 0
	if len(searchResults) == 0 {
		message = "Aucun résultat pour votre recherche"
	} else {
		message = ""
	}

	// Struct de data que je vais envoyer
	data := struct {
		Usersearch string
		Message    string
		Results    []fonction.SearchResults
	}{
		Usersearch: usersearch,
		Message:    message,
		Results:    searchResults,
	}

	// Init avec mes datas
	InitTemplate.Temp.ExecuteTemplate(w, "search", data)
}

// Handler pour afficher le détail du pays selectionné
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du nom du pays dans la query pour l'afficher dans les détails après
	countryname := r.URL.Query().Get("country")
	Country := fonction.SearchCountry(countryname)
	// Gestio d'erreur
	if len(Country) == 0 {
		http.Error(w, "Not Found: No results found", http.StatusNotFound)
		return
	}

	// Init avec les datas du pays
	InitTemplate.Temp.ExecuteTemplate(w, "details", Country)
}

// Treatment pour ajouter les favoris dans le fichier Json
func AddFavoriteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	// Je récupère les données dans les query
	countryName := r.URL.Query().Get("country")
	continent := r.URL.Query().Get("reg")
	area := r.URL.Query().Get("area")
	population := r.URL.Query().Get("pop")
	flag := r.URL.Query().Get("flag")
	isfavorite := true

	// Struct de données que je vais envoyer dans ma fonction addFav
	favorite := fonction.FavoriteInfo{
		Name:       countryName,
		Continent:  continent,
		Area:       area,
		Population: population,
		Flag:       flag,
		IsFavorite: isfavorite,
	}

	// J'ajoute mes datas dans mon Json
	fonction.AddFav(favorite)
	// Je redirige vers la page favoris pour montrer à l'utilisateur que son pays est bien ajouté
	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}

// Handler pour afficher les pays mis en favoris
func FavoriteHandler(w http.ResponseWriter, r *http.Request) {

	// Peremt d'afficher tous les pays en favoris
	favs, err := fonction.RetrieveFavorite()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des favoris", http.StatusInternalServerError)
		return
	}
	// Message si il n'y a pas de favoris
	var message string
	if len(favs) == 0 {
		message = "Vous n'avez aucun favoris"
	} else {
		message = ""
	}

	// Struct de data qui sera envoyé
	data := struct {
		Favorite []fonction.FavoriteInfo
		Message  string
	}{
		Favorite: favs,
		Message:  message,
	}
	// Init et envoie des données vers le template
	InitTemplate.Temp.ExecuteTemplate(w, "favorite", data)
}

// Handler pour supprimer un pays des favoris
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Je récupère la query du nom du pays
	countryName := r.FormValue("country")
	// Si le pays est présent de mon fichier Json ça l'enlève
	err := fonction.RemoveFavorite(countryName)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du pays favori", http.StatusInternalServerError)
		return
	}
	// Redirection vers favoris
	http.Redirect(w, r, "/favorite", http.StatusSeeOther)
}

// Template about avec toute la doc
func AboutHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "about", nil)
}

// Template Error404
func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "error", nil)
}
