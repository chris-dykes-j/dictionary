package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Entry struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
}

func main() {
	args := os.Args[1:]
	if args == nil || len(args) == 0 {
		fmt.Println("Need a word to define.")
	} else {
		word := args[0]
		url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: Problem with url.")
			panic(err)
		}

		var entries []Entry
		err = json.NewDecoder(resp.Body).Decode(&entries)
		if err != nil {
			panic(err)
		}

		if len(args) > 1 {
			partOfSpeech := args[1]
			printDefinitions(entries, partOfSpeech)
		} else {
			printDefinitions(entries, "")
		}
	}
}

func printDefinitions(entries []Entry, partOfSpeech string) {
	for _, entry := range entries {
		for _, meaning := range entry.Meanings {

			if partOfSpeech != "" && meaning.PartOfSpeech != partOfSpeech {
				continue
			}

			// Header: Word (part of speech)
			fmt.Printf("%s (%s):\n", capitalize(entry.Word), meaning.PartOfSpeech)

			for i, def := range meaning.Definitions {
				fmt.Printf("%d. %s\n", i+1, capitalize(def.Definition))
			}

			fmt.Println()
		}
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[:1] == "(" {
		return strings.ToUpper(s[:2]) + s[2:]
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
