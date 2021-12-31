package routers

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/models"
	"net/http"
)

func CreateRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
