package main

import "fmt"

type books struct {
	title    string
	author   string
	subject  string
	books_id int
}

func main() {

	b := books{
		title:    "chetan's book",
		author:   "chetan",
		subject:  "comedy",
		books_id: 3,
	}
	fmt.Printf("new book name %+v \n", b)

	var book1 books //declare book1 of type book
	// var book2 book //declare book2 of type book

	book1.title = "go prog"
	book1.author = "harsh"
	book1.subject = "tutorials"
	book1.books_id = 01

	//print book 1 info

	// fmt.Printf("book 1 title: %s \n", book1.title)
	// fmt.Printf("book 1 author: %s \n", book1.author)
	// fmt.Printf("book 1 subject: %s \n", book1.subject)
	// fmt.Printf("book 1 book id: %d \n", book1.books_id)

	printbook(book1) //calling functoin

}
func printbook(book books) {
	fmt.Printf("book 1 title: %s \n", book.title)
	fmt.Printf("book 1 author: %s \n", book.author)
	fmt.Printf("book 1 subject: %s \n", book.subject)
	fmt.Printf("book 1 book id: %d \n", book.books_id)

}
