package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello my friend")

	stor := http.NewServeMux()
	stor.HandleFunc("POST /api/v1/wallet", postWallet)
	stor.HandleFunc("GET /api/v1/wallet/{uuid}", getWallet)

	if err := http.ListenAndServe(":8000", stor); err != nil {
		panic(err)
	}
}
