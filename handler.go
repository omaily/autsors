package main

import (
	"fmt"
	"net/http"
)

func getWallet(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	w.Write([]byte(fmt.Sprintf("Hello %s!", uuid)))
}

func postWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("i "))
		return
	}

	w.Write([]byte(`Hello`))
}
