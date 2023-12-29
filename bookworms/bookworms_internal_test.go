package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	bookwormsFile string 
	want []Bookworm 
	wantErr bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func equalBooks(books, targetBooks []Book) bool {
	for i := range books{
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

func TestLoadBookworms(t *testing.T) {
	tests := map[string] TestCase { 
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

			// if !equalBookworms(got, testCase.want) {
			// 	t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			// }

			// Does the same as what's above
			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want) }
			})
	}
}