package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getWallet(user IUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.PathValue("uuid")
		w.Write([]byte(fmt.Sprintf("Hello %s!", uuid)))

		amount, err := user.getAmount(r.Context(), uuid)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(strconv.Itoa(amount)))
	}
}

func postWallet(user IUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("i "))
			return
		}

		w.Write([]byte(`Hello`))
	}
}
