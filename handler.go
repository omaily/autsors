package main

import (
	"fmt"
	"net/http"
)

func getWallet(user *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.PathValue("uuid")
		w.Write([]byte(fmt.Sprintf("Hello %s!", uuid)))

		amount, err := user.getAmount(r.Context(), uuid)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(amount))
	}
}

func postWallet(user *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`post and handler`))
	}
}
