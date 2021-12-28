//nolint
package handlers

import (
	"Documents/Go/twitter-go/middlew"
	"Documents/Go/twitter-go/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	//Users
	router.HandleFunc("/register", middlew.ChequeoDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.ChequeoDB(middlew.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/edit", middlew.ChequeoDB(middlew.ValidateJWT(routers.Edit))).Methods("PUT")

	//Tweets
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/get/tweets", middlew.ChequeoDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/del/tweet", middlew.ChequeoDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
