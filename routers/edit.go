package routers

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/models"
	"encoding/json"
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.EditUser(user, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el regristo"+err.Error(), http.StatusBadRequest)
	}

	if status == false {
		http.Error(w, "No se logro modificar el registro del usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
