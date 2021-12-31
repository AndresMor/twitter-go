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

	router.HandleFunc("/list", middlew.ChequeoDB(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")

	router.HandleFunc("/upload/avatar", middlew.ChequeoDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/get/avatar", middlew.ChequeoDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")

	//Tweets
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweet/followers", middlew.ChequeoDB(middlew.ValidateJWT(routers.GetTweetsRelationShip))).Methods("GET")
	router.HandleFunc("/get/tweets", middlew.ChequeoDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/del/tweet", middlew.ChequeoDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	//Relationship
	router.HandleFunc("/create/relationship", middlew.ChequeoDB(middlew.ValidateJWT(routers.CreateRelationship))).Methods("POST")
	router.HandleFunc("/del/relationship", middlew.ChequeoDB(middlew.ValidateJWT(routers.DeleteRelationShip))).Methods("DELETE")
	router.HandleFunc("/read/relationship", middlew.ChequeoDB(middlew.ValidateJWT(routers.ReadRelationship))).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
