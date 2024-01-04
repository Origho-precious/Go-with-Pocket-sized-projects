package main

import (
	"math"
	"sort"
)

type Recommendation struct {
	Book  Book
	Score float64
}

type set map[Book]struct{}

func newSet(books []Book) set {
	s := make(set)
	for _, book := range books {
		s[book] = struct{}{}
	}
	return s
}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

func recommend(allReaders []Bookworm, target Bookworm, n int) []Recommendation {
	read := newSet(target.Books)

	recommendations := map[Book]float64{}

	for _, reader := range allReaders {
		if reader.Name == target.Name {
			continue
		}

		var similarity float64
		for _, book := range reader.Books {
			if read.Contains(book) {
				similarity++
			}
		}
		if similarity == 0 {
			continue
		}

		score := math.Log(similarity) + 1
		for _, book := range reader.Books {
			if !read.Contains(book) {
				recommendations[book] += score
			}
		}
	}

	// TODO: sort by score
	var sortedRecomms []Recommendation

	for book, score := range recommendations {
		sortedRecomms = append(sortedRecomms, Recommendation{book, score})
	}

	sort.Slice(sortedRecomms, func(i, j int) bool {
		return sortedRecomms[i].Score > sortedRecomms[j].Score
	})

	// TODO: only output a certain amount of recommendations (n)
	return sortedRecomms
}
