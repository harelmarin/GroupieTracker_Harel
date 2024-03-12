package routeur

import (
	"Groupie/controller"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Gestionnaire des routes
func Serveur() {
	http.HandleFunc("/index", controller.IndexHandler)
	http.HandleFunc("/treatment/index", controller.IndexTreatmentHandler)
	http.HandleFunc("/search", controller.RechercheHandler)                        //Query => /search?usersearch=""
	http.HandleFunc("/category", controller.CateHandler)                           //Query => /category?category=""
	http.HandleFunc("/details", controller.DetailsHandler)                         //Query => /details?country=""
	http.HandleFunc("/favorite/treatment", controller.AddFavoriteTreatmentHandler) //Query => /favorite/treatment?country=""
	http.HandleFunc("/favorite", controller.FavoriteHandler)
	http.HandleFunc("/delete/treatment", controller.DeleteHandler) //Query => /delete/treatment?country=""
	http.HandleFunc("/about", controller.AboutHandler)
	http.HandleFunc("/error", controller.ErrorHandler)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// Gestionnaire pour les routes non trouvées (404)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/error", http.StatusNotFound)
	})
	// Lance le serveur
	runServer()
}

// Fonction qui permet de lancer le serveur
func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/index"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 3)
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}
