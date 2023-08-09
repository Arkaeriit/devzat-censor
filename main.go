package main

import (
    "fmt"
    "os"

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
            censored := rmBadWords(txt)
            replyChan <- censored
            if txt == censored {
                fmt.Printf("'%s' does not need censoring.\n", txt)
            } else {
                fmt.Printf("'%s' get censored into '%s'.\n", txt, censored)
            }
        }
    }
}

