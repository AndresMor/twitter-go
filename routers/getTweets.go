package routers

import (
	"Documents/Go/twitter-go/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pag")) < 1 {
		http.Error(w, "Debe enviar la página", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(r.URL.Query().Get("pag"))
	if err != nil {
		http.Error(w, "Debe enviar la página con valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(pag)
	response, status := bd.ReadTweets(ID, page)
	if status == false {
		http.Error(w, "Error al leer los tweets "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
