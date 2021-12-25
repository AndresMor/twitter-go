package routers

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/models"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres ", http.StatusBadRequest)
		return
	}

	_, finded, _ := bd.CheckUserExists(t.Email)
	if finded == true {
		http.Error(w, "Ya existe un usuario con ese Email ", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar el usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el regristro del usuario ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
