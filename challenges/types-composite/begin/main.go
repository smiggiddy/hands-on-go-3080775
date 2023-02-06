// challenges/types-composite/begin/main.go
package main

import "fmt"

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(nb book) {
	l[nb.author.name] = append(l[nb.author.name], nb)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(author string) []book {
	return l[author]
}

func main() {
	// create a new library
	myLibrary := library{}

	// add 2 books to the library by the same author
	book1 := book{
		title:  "Book1",
		author: author{name: "Some guy"},
	}
	book2 := book{
		title:  "Book2",
		author: author{name: "Some guy"},
	}
	myLibrary.addBook(book1)
	myLibrary.addBook(book2)
	// add 1 book to the library by a different author
	book3 := book{
		title:  "different",
		author: author{name: "different guy"},
	}
	myLibrary.addBook(book3)
	// dump the library with spew
	fmt.Printf("%#v\n", myLibrary)

	// lookup books by known author in the library
	fmt.Println(myLibrary.lookupByAuthor("Some guy"))

	// print out the first book's title and its author's name
	fmt.Println(myLibrary.lookupByAuthor("Some guy")[0])

}
