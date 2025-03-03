package main

import (
	"fmt"
	"log"

	"github.com/deadsy/go-cli"
)

func main() {
	prompt := "user> "
	ln := cli.NewLineNoise()
	ln.SetMultiline(true)
	ln.HistorySetMaxlen(100)

	// reader := bufio.NewReader(os.Stdin)

	for {
		text, err := ln.Read(prompt, "")
		if err != nil {
			log.Fatal(err)
		}
		ln.HistoryAdd(text)
		ln.HistorySave("history.txt")
		out := rep(text)
		fmt.Println(out)
	}
}

func read(s string) string {
	// log.Println("Read")
	return s
}

func eval(s string) string {
	// log.Println("Eval")
	return s
}

func print(s string) string {
	// log.Println("Print")
	return s
}

func rep(s string) string {
	s = read(s)
	s = eval(s)
	s = print(s)
	return s
}
