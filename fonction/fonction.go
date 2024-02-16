package fonction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
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
				Name:        c.Translations.Fra.Common,
				Flag:        c.Flags.Png,
				Area:        int(c.Area),
				Pop:         int(c.Population),
				Reg:         c.Region,
				IsFavorite:  bool(IsCountryFavorite(c.Translations.Fra.Common)),
				Independent: bool(c.Independent),
			}
			searchResults = append(searchResults, result)
		}
	}
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
			Name:        c.Translations.Fra.Common,
			Flag:        c.Flags.Png,
			Area:        int(c.Area),
			Pop:         int(c.Population),
			Reg:         c.Region,
			IsFavorite:  bool(IsCountryFavorite(c.Translations.Fra.Common)),
			Independent: bool(c.Independent),
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

type FilterOptions struct {
	Independent  bool
	Alphabetical bool
}

// Fonction qui permet de filter mes pays en fonction des options cliqués
func FilterCountries(countries []SearchResults, options FilterOptions) []SearchResults {

	// Filtre par indépendance
	if options.Independent {
		var filteredCountries []SearchResults
		for _, country := range countries {
			if country.Independent {
				filteredCountries = append(filteredCountries, country)
			}
		}
		countries = filteredCountries
	}

	// Filtrer par ordre alphabétique
	if options.Alphabetical {
		sort.Slice(countries, func(i, j int) bool {
			return countries[i].Name < countries[j].Name
		})
	}

	return countries
}

func SearchIndex() []SearchResults {

	URLSearch := "https://restcountries.com/v3.1/lang/French"

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
		searchResults = append(searchResults, result)
	}
	return searchResults

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
		if result.Pop >= minPopulation && result.Pop <= maxPopulation {
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
