package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",             // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
	"de": "Hallo Welt", // German
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}

func main() {

	var lang string

	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")

	flag.Parse()

	greeting := greet(language(lang))

	fmt.Println(greeting)
}
