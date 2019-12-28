package main

//@xyz
type Book struct {
	Name, Title string
	Author int
	PriceWithFee int
}

//@xyz
type BookRepository interface {
	FindByNameAndId(name string)
	UpdateNameAndFamilyNameBasedOnId(name string, fname string, id string)
	DeleteByName(name string)
}

