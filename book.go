package main

//go:generate xyz $GOFILE

//@xyz
type Book struct {
	Name, Title  string
	Author       int
	PriceWithFee int
}

//@xyz
type BookRepository interface {
	FindByNameAndId(args ...interface{}) ([]Book, error)
	UpdateNameAndFamilyNameBasedOnId(args ...interface{}) error
	DeleteByName(args ...interface{}) error
}
