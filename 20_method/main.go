// https://www.youtube.com/watch?v=GPlJQ-JIjSA
package main

import "fmt"

type book struct {
	author   string
	bookname string
	books_id int
}

func (b book) detail() {
	fmt.Println(b.author, b.bookname, b.books_id)
}
func main() {
	b := book{
		author:   "harsh bhai",
		bookname: "my book",
		books_id: 01,
	}

	b.detail()
	fmt.Printf("my books details is %+v \n", b)
}
