package iterator

// Aggregate 集合接口
// 产生该接口的迭代器
type Aggregate interface {
	iterator() Iterator
}

// Iterator 迭代器接口
type Iterator interface {
	hasNext() bool
	next() interface{}
}

// BookShelf 具体集合
type BookShelf struct {
	books []Book
	last  int
}

func (b *BookShelf) getBookAt(index int) Book {
	return b.books[index]
}

func (b *BookShelf) getLength() int {
	return b.last
}

func (b *BookShelf) iterator() Iterator {
	return &BookShelfIterator{bookShelf: *b, index: 0}
}

func (b *BookShelf) addBook(book Book) {
	b.books = append(b.books, book)
	b.last++
}

type Book struct {
	name string
}

func (b Book) getName() string {
	return b.name
}

// BookShelfIterator 具体的迭代器
type BookShelfIterator struct {
	index     int
	bookShelf BookShelf
}

func (b *BookShelfIterator) hasNext() bool {
	return b.index < b.bookShelf.getLength()
}

func (b *BookShelfIterator) next() interface{} {
	book := b.bookShelf.getBookAt(b.index)
	b.index++
	return book
}
