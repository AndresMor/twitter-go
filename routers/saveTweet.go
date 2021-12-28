package routers

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/models"
	"encoding/json"
	"net/http"
	"time"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.RequestTweet
	err := json.NewDecoder(r.Body).Decode(&message)

	record := models.Tweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(record)
	if err != nil {
		http.Error(w, "Ocurrio un error al interntar guardar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el twwet ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
