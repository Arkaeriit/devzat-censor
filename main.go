package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"os"
	"unicode"

	api "github.com/quackduck/devzat/devzatapi"
)

func runPlugin() error {
	host := os.Getenv("PLUGIN_HOST")
	if host == "" {
		host = "devzat.hackclub.com:5556"
	}

	s, err := api.NewSession(host, os.Getenv("PLUGIN_TOKEN"))
	if err != nil {
		panic(err)
	}

	messageChan, replyChan, err := s.RegisterListener(true, false, "")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err = <-s.ErrorChan:
			panic(err)
		case msg := <-messageChan:
			txt := msg.Data
			clean := removeDiactrics(txt)
			censored := rmBadWords(clean)
			if clean == censored {
				fmt.Printf("'%s' does not need censoring.\n", txt)
				replyChan <- txt
			} else {
				fmt.Printf("'%s' get censored into '%s'.\n", txt, censored)
				replyChan <- censored
			}
		}
	}
}

func main() {
	for {
		err := runPlugin()
		fmt.Printf("!!!! %v !!!!\n", err)
	}
}

func removeDiactrics(in string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, in)
	return s
}

