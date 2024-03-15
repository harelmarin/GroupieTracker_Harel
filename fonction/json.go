package fonction

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Struct info des Pays Favoris
type FavoriteInfo struct {
	Name       string `json:"name"`
	Continent  string `json:"continent"`
	Area       string `json:"area"`
	Population string `json:"population"`
	Flag       string `json:"flag"`
	IsFavorite bool   `json:"isfavorite"`
}

type ConsultInfos struct {
}

// Constante pour le fichier Json
var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(b) + "\\"
)
var jsonfile = path + "../content\\data.json"

// Lit le fichier Json pour retrouver toutes les infos des favoris
func RetrieveFavorite() ([]FavoriteInfo, error) {
	var Fav []FavoriteInfo

	data, err := os.ReadFile(jsonfile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return Fav, nil
	}

	err = json.Unmarshal(data, &Fav)
	if err != nil {
		return nil, err
	}

	return Fav, nil
}

// ChangeFav écrit la liste de favoris dans le fichier JSON
func ChangeFav(Fav []FavoriteInfo) error {
	data, errJSON := json.Marshal(Fav)
	if errJSON != nil {
		return errJSON
	}

	errWrite := os.WriteFile(jsonfile, data, 0666)
	if errWrite != nil {
		return errWrite
	}

	return nil
}

// Add an adventurer in Json
func AddFav(info FavoriteInfo) {
	// Récupérer la liste des favoris actuelle
	Fav, err := RetrieveFavorite()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}

	// Ajouter le nouveau favori à la liste
	Fav = append(Fav, info)

	// Enregistrer la liste mise à jour dans le fichier JSON
	err = ChangeFav(Fav)
	if err != nil {
		log.Fatal("log: addArticle()\t WriteFile error!\n", err)
	}
}

// Remove le pays en favoris
func RemoveFavorite(countryName string) error {

	favs, err := RetrieveFavorite()
	if err != nil {
		return err
	}

	// Parcourir la liste des favoris pour trouver le pays à supprimer
	for i, fav := range favs {
		if fav.Name == countryName {

			favs[i] = favs[len(favs)-1]
			favs = favs[:len(favs)-1]

			err := ChangeFav(favs)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

// Fonction pour savoir si le pays est dans les favoris ou pas ||| Si oui => return true
func IsCountryFavorite(countryName string) bool {
	favs, err := RetrieveFavorite()
	if err != nil {
		return false
	}

	for _, fav := range favs {
		if fav.Name == countryName && fav.IsFavorite {
			return true
		}
	}
	return false
}
