package main

//@xyz Model
type Book struct {
	Name string
	Author int
}

//@xyz Repo
type BookRepository interface {
	FindByName(name string) // select * from books where name=$name
	UpdateNameAndFNameBasedOnId(name string, fname string, id string)
}



