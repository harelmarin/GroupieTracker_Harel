package fonction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var DecodeData Data

type Data []struct {
	Translations struct {
		Fra struct {
			Common string `json:"common"`
		} `json:"fra"`
	} `json:"translations"`
}

func Gettranslationcountry() {
	URLbasewinner := "https://restcountries.com/v3.1/translation/Germany"

	//Init Client
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	//Créer requête HTTP
	req, errReq := http.NewRequest(http.MethodGet, URLbasewinner, nil)
	if errReq != nil {
		fmt.Println("Erreur avec la requête : ", errReq.Error())
		return
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
		return
	}
	for _, country := range DecodeData {
		fmt.Printf("Nom français : %s\n", country.Translations.Fra.Common)
	}
}
