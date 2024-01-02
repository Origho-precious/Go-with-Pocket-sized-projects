package main

import (
	"fmt"
	// "reflect"
	"testing"
)

type LoadBookwormsTestCase struct {
	bookwormsFile string
	want          []Bookworm
	wantErr       bool
}

type BooksCountTestCase struct {
	bookCountMap map[Book]uint
	want         bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}

	testBookCount1 = map[Book]uint{
		handmaidsTale: 2,
		oryxAndCrake:  1,
		theBellJar:    1,
		janeEyre:      1,
	}

	testBookCount2 = map[Book]uint{
		handmaidsTale: 1,
		oryxAndCrake:  1,
		theBellJar:    1,
		janeEyre:      1,
	}
)

func equalBooks(books, targetBooks []Book) bool {
	for i := range books {
		if books[i].Author != targetBooks[i].Author {
			return false
		}

		if books[i].Title != targetBooks[i].Title {
			return false
		}
	}

	return true
}

func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if len(bookworms[i].Books) != len(target[i].Books) {
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]

		if !ok || count != targetCount {
			return false
		}
	}

	return true
}

func TestLoadBookworms(t *testing.T) {
	tests := map[string]LoadBookwormsTestCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)

			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if !equalBookworms(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}

			// Does the same as what's above
			// if !reflect.DeepEqual(got, testCase.want) {
			// 	t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			// }
		})
	}
}

func TestBooksCount(t *testing.T) {
	tests := map[string]BooksCountTestCase{
		"Equal book count": {
			bookCountMap: testBookCount1,
			want:         true,
		},
		"unequal book count": {
			bookCountMap: testBookCount2,
			want:         false,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			bookworms, err := loadBookworms("./testdata/bookworms.json")

			if err != nil {
				fmt.Println("Couldn't load file")
			}

			actualBooksCount := booksCount(bookworms)

			got := equalBooksCount(actualBooksCount, testCase.bookCountMap)

			if testCase.want != got {
				t.Errorf("expected: %v, got: %v", testCase.want, got)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	testCases := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale}},
			},
			want: []Book{handmaidsTale},
		},
	}

	for name, tc := range testCases {
		if name == "no common book" {
			t.Run(name, func(t *testing.T) {
				got := findCommonBooks(tc.input)

				if !equalBooks(tc.want, got) {
					t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
				}
			})
		} else {
			t.Run(name, func(t *testing.T) {
				got := findCommonBooks(tc.input)

				if !equalBooks(tc.want, got) {
					t.Fatalf("expected: %v, got %v", tc.want, got)
				}
			})
		}
	}
}
