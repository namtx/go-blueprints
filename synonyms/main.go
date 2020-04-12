package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/namtx/go-blueprints/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_API_KEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("Failed when looking for synonyms for '"+word+"'", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find synonyms for '" + word + "'")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
