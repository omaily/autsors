package main

import "context"

type User struct {
	Uuid   string
	Name   string
	Amount int
}

type IUser interface {
	depositPay(context.Context, int) error
	withdrawPay(context.Context, int) error
	getAmount(context.Context, string) (int, error)
}
