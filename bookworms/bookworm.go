package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type byAuthor []Book

func (b byAuthor) Len() int { return len(b) }

func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}

	return b[i].Title < b[j].Title
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(f).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
