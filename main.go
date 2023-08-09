package main

import (
    "fmt"
    "os"
    "unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

    api "github.com/quackduck/devzat/devzatapi"
)

func main() {
    s, err := api.NewSession("devzat.hackclub.com:5556", os.Getenv("DEVZAT_TOKEN"))
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

func removeDiactrics(in string) string {
    t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
    s, _, _ := transform.String(t, in)
    return s
}


