package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	baseURL = "https://www.dictionaryapi.com/api/v3/references/ithesaurus/json/"
)

type Entry struct {
	Meta struct {
		ID        string     `json:"id"`
		Stems     []string   `json:"stems,omitempty"`
		Synonyms  [][]string `json:"syns,omitempty"`
		Antonyms  [][]string `json:"ants,omitempty"`
		Offensive bool       `json:"offensive,omitempty"`
	} `json:"meta"`
	Fl       string   `json:"fl"`
	Shortdef []string `json:"shortdef"`
}

func fetchThesaurus(word string) ([]Entry, error) {
	apiKey := os.Getenv("MW_API_KEY")
	url := fmt.Sprintf("%s%s?key=%s", baseURL, word, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status: %s, body: %s", resp.Status, body)
	}

	var entries []Entry
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&entries); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return entries, nil
}

func thesaurus(word string, done chan bool) {
	entries, err := fetchThesaurus(word)
	if err != nil {
		log.Fatalf("Error fetching thesaurus data: %v", err)
	}

	PrintHeader("\nMerriam-Webster's Intermediate Thesaurus\n")
	for _, entry := range entries {
		for _, syn := range entry.Meta.Synonyms {
			synonyms := fmt.Sprintf("Synonyms: %v", strings.Join(syn, ", "))
			PrintWord(synonyms, "magenta")
		}

		fmt.Printf("Word: %s\nPart of Speech: %s\nDefinitions:\n", entry.Meta.ID, entry.Fl)
		for _, def := range entry.Shortdef {
			fmt.Printf(" - %s\n", def)
		}
		fmt.Println()
	}
	done <- true
}
