package domain

import "github.com/Banking-System/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	// 业务意图，而不是技术意图
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
