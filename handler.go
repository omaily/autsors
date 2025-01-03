package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func postWallet(user *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Header.Get("valletId")
		operationType := r.Header.Get("operationType")
		amount := r.Header.Get("amount")

		if uuid == "" || operationType == "" || amount == "" {
			w.Write([]byte(`there are no values associated with the header`))
			return
		}

		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			w.Write([]byte(`the value in the header "amount" is incorrect`))
			return
		}

		if operationType == "DEPOSIT" {
			err := user.depositPay(r.Context(), uuid, amountInt)
			if err != nil {
				w.Write([]byte(`Failed to credit funds`))
				return
			}
		} else if operationType == "WITHDRAW" {
			err := user.withdrawPay(r.Context(), uuid, amountInt)
			if err != nil {
				w.Write([]byte(`Failed to write off funds`))
				return
			}
		} else {
			w.Write([]byte(`The value in the header "operationType" is incorrect`))
			return
		}

		w.Write([]byte(`post and handler:`))
	}
}

func getWallet(user *Storage) http.HandlerFunc {
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
