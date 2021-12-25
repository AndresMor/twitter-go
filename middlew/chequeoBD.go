package middlew

import (
	"Documents/Go/twitter-go/bd"
	"net/http"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == false {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
