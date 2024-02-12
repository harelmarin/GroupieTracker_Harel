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
	Name       string
	Flag       string
	Area       int
	Pop        int
	Reg        string
	NbPage     int
	IsFavorite bool
}

// requête pour récupérer les données d'un pays via une recherche user (recherche en Français)
func SearchCountry(usersearch string) []SearchResults {

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

	var searchResults []SearchResults
	for _, c := range DecodeData {
		if strings.Contains(strings.ToLower(c.Translations.Fra.Common), strings.ToLower(usersearch)) {
			result := SearchResults{
				Name:       c.Translations.Fra.Common,
				Flag:       c.Flags.Png,
				Area:       int(c.Area),
				Pop:        int(c.Population),
				Reg:        c.Region,
				IsFavorite: bool(IsCountryFavorite(c.Translations.Fra.Common)),
			}
			searchResults = append(searchResults, result)
		}
	}
	fmt.Println(searchResults)
	return searchResults

}

// requête pour récupérer les données des pays du continent choisi
func GetContinent(categorie string) []SearchResults {

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
			Name: c.Translations.Fra.Common,
			Flag: c.Flags.Png,
			Area: int(c.Area),
			Pop:  int(c.Population),
		}
		searchResults = append(searchResults, result)
	}
	return searchResults
}

// Met des espaces tous les 3 nombres pour rendre le tout lisible
func addThousandsSeparator(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return addThousandsSeparator(s[:n-3]) + " " + s[n-3:]
}

// Formate l'aire récuperé pour la rendre lisible facilement en enlevant les 0 après le "."
func formatNumber(Number float64) string {

	areaStr := fmt.Sprintf("%.0f", Number)

	parts := strings.Split(areaStr, ".")

	// Récupérer la partie entière
	intPart := parts[0]

	formattedIntPart := addThousandsSeparator(intPart)

	if len(parts) > 1 {
		return formattedIntPart + "." + parts[1]
	}

	return formattedIntPart
}
