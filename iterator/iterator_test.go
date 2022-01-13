package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	book1 := Book{name: "1"}
	book2 := Book{name: "2"}
	book3 := Book{name: "3"}
	book4 := Book{name: "4"}

	bookShelf := BookShelf{}
	bookShelf.addBook(book1)
	bookShelf.addBook(book2)
	bookShelf.addBook(book3)
	bookShelf.addBook(book4)

	bookShelfIterator := bookShelf.iterator()

	for bookShelfIterator.hasNext() {
		book := bookShelfIterator.next()
		fmt.Println(book.(Book).getName())
	}
	fmt.Println("over")

}
