package routers

import (
	"Documents/Go/twitter-go/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetTweetsRelationShip(w http.ResponseWriter, r *http.Request) {
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
	respt, correct := bd.ReadTweetsFollowers(IDUser, page)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respt)

}
