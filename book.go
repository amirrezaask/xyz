package main
//@xyz
type Book struct {
	Name, Title string
	Author int
}

//@xyz
type BookRepository interface {
	FindByNameAndId(name string) // select * from books where name=$name
	UpdateNameAndFamilyNameBasedOnId(name string, fname string, id string)
	DeleteByName(name string)
}



