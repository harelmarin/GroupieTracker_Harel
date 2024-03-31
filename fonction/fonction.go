package fonction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var DecodeData DataSearch

// Struct de données pour récupérer les données envoyés par mon API
type DataSearch []struct {
	Translations struct {
		Fra struct {
			Common string `json:"common"`
		} `json:"fra"`
	} `json:"translations"`
	Independent bool    `json:"independent"`
	Area        float64 `json:"area"`
	Flags       struct {
		Png string `json:"png"`
	} `json:"flags"`
	Population float64 `json:"population"`
	Region     string  `json:"region"`
	IsFavorite bool
}
type SearchResults struct {
	Name        string
	Flag        string
	Area        int
	Pop         int
	Reg         string
	Independent bool
	IsFavorite  bool
}

// requête pour récupérer les données d'un pays via une recherche user (recherche en Français)
func SearchCountry(usersearch string) []SearchResults {

	// EndPoint + la recherche de l'utilisateur dans la barre
	URLSearch := "https://restcountries.com/v3.1/translation/" + usersearch

	//Init Client
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	//Créer requête HTTP
	req, errReq := http.NewRequest(http.MethodGet, URLSearch, nil)
	if errReq != nil {
		fmt.Println("Erreur avec la requête : ", errReq.Error())

	}

	//Exécution HTTP
	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de la requête HTTP")
	}

	defer res.Body.Close()
	//Lecture et Récup de la requête
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Problème avec lecture du corps")
	}

	//Décodage des données
	if err := json.Unmarshal(body, &DecodeData); err != nil {
		fmt.Println("Erreur lors du décodage des données JSON (Nom problablement incorrect)")
		fmt.Println("Réponse JSON reçue:", string(body))

	}

	// Bar de recherche qui recherche via un endpoint de mon API (récupération du nom en français)
	var searchResults []SearchResults
	for _, c := range DecodeData {
		if strings.Contains(strings.ToLower(c.Translations.Fra.Common), strings.ToLower(usersearch)) {
			result := SearchResults{
				Name:        c.Translations.Fra.Common,
				Flag:        c.Flags.Png,
				Area:        int(c.Area),
				Pop:         int(c.Population),
				Reg:         c.Region,
				IsFavorite:  bool(IsCountryFavorite(c.Translations.Fra.Common)),
				Independent: bool(c.Independent),
			}
			// J'append les résultats pour les return
			searchResults = append(searchResults, result)
		}
	}

	return searchResults
}

// requête pour récupérer les données des pays du continent choisi
func GetContinent(categorie string) []SearchResults {

	// Endpoint + continent selectionné
	URLSearch := "https://restcountries.com/v3.1/region/" + categorie

	//Init Client
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	//Créer requête HTTP
	req, errReq := http.NewRequest(http.MethodGet, URLSearch, nil)
	if errReq != nil {
		fmt.Println("Erreur avec la requête : ", errReq.Error())

	}

	//Exécution HTTP
	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de la requête HTTP")
	}
	defer res.Body.Close()
	//Lecture et Récup de la requête
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Problème avec lecture du corps")
	}

	//Décodage des données
	if err := json.Unmarshal(body, &DecodeData); err != nil {
		fmt.Println("Erreur lors du décodage des données JSON (Nom problablement incorrect)")
		fmt.Println("Réponse JSON reçue:", string(body))
	}

	var searchResults []SearchResults
	for _, c := range DecodeData {
		result := SearchResults{
			Name:        c.Translations.Fra.Common,
			Flag:        c.Flags.Png,
			Area:        int(c.Area),
			Pop:         int(c.Population),
			Reg:         c.Region,
			IsFavorite:  bool(IsCountryFavorite(c.Translations.Fra.Common)),
			Independent: bool(c.Independent),
		}
		// J'append les résultats pour les return
		searchResults = append(searchResults, result)
	}
	return searchResults

}

// Struct pour les filtres
type FilterOptions struct {
	Independent  bool
	Alphabetical bool
}

// Index
func SearchIndex() ([]SearchResults, error) {

	// Affiche tous les pays
	URLSearch := "https://restcountries.com/v3.1/all"

	//Init Client
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	//Créer requête HTTP
	req, errReq := http.NewRequest(http.MethodGet, URLSearch, nil)
	if errReq != nil {
		return nil, fmt.Errorf("erreur avec la requête : %s", errReq.Error())

	}

	//Exécution HTTP
	res, errRes := httpClient.Do(req)
	if errRes != nil {
		return nil, fmt.Errorf("erreur lors de la requête HTTP : %s", errRes.Error())

	}
	defer res.Body.Close()
	//Lecture et Récup de la requête
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Problème avec lecture du corps")
	}

	//Décodage des données
	if err := json.Unmarshal(body, &DecodeData); err != nil {
		fmt.Println("Erreur lors du décodage des données JSON (Nom problablement incorrect)")
		fmt.Println("Réponse JSON reçue:", string(body))

	}

	var searchResults []SearchResults
	for _, c := range DecodeData {
		result := SearchResults{
			Name:        c.Translations.Fra.Common,
			Flag:        c.Flags.Png,
			Area:        int(c.Area),
			Pop:         int(c.Population),
			Reg:         c.Region,
			IsFavorite:  bool(IsCountryFavorite(c.Translations.Fra.Common)),
			Independent: bool(c.Independent),
		}
		// Append les résultats
		searchResults = append(searchResults, result)
	}
	return searchResults, nil

}

// Fonction pour filtrer par indépendance
func FilterIndependent(results []SearchResults) []SearchResults {
	var filteredResults []SearchResults
	for _, result := range results {
		if result.Independent {
			filteredResults = append(filteredResults, result)
		}
	}
	return filteredResults
}

// Fonction pour filtrer par population
func FilterByPopulation(results []SearchResults, minPopulation, maxPopulation int) []SearchResults {
	var filteredResults []SearchResults
	for _, result := range results {
		// Si le min n'est pas spécifié ou si la population est supérieure ou égale au min spécifié
		minPass := minPopulation == 0 || result.Pop >= minPopulation
		// Si le max n'est pas spécifié ou si la population est inférieure ou égale au max spécifié
		maxPass := maxPopulation == 0 || result.Pop <= maxPopulation

		// Ajoutez le résultat si les conditions min et max sont remplies
		if minPass && maxPass {
			filteredResults = append(filteredResults, result)
		}
	}
	return filteredResults
}

// Fonction pour filtrer par ordre alphabétique
func FilterAlphabetical(results []SearchResults) []SearchResults {
	var filteredResults []SearchResults

	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Name > results[j].Name {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	filteredResults = append(filteredResults, results...)

	return filteredResults
}

// Permet de return les résultats paginé
func GetPageResults(results []SearchResults, page int) []SearchResults {
	startIndex := (page - 1) * 10
	endIndex := startIndex + 10

	if startIndex >= len(results) {
		return []SearchResults{}
	}

	if endIndex > len(results) {
		endIndex = len(results)
	}

	return results[startIndex:endIndex]
}
