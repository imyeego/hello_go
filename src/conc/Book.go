package main

import "fmt"

type Behavior interface {
	read(title string)
}

type Book struct {
	title string
	author string
	subject string
	book_id int
}

func (book *Book) read(title string) {
	fmt.Println("reading the " + title)
}

func (book *DoubanBook) read(title string)  {
	fmt.Printf("reading the %s with %.2f score", title, book.score)

}

type DoubanBook struct {
	Book
	score float32
}

func printBook(book *Book)  {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book id: %d\n", book.book_id)

}

func (book *Book) showInfo() {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book id: %d\n", book.book_id)
}

func (book *Book) setTitle(title string) {
	book.title = title
}



func (doubanBook *DoubanBook) getScore() float32  {
	return doubanBook.score
}