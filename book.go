package main

//@xyz
type Book struct {
	Name, Title  string
	Author       int
	PriceWithFee int
}

//@xyz
type BookRepository interface {
	FindByNameAndId(name string) (Book, error)
	UpdateNameAndFamilyNameBasedOnId(name string, fname string, id string) error
	DeleteByName(name string) error
}
