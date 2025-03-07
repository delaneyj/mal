package mal

import (
	"fmt"
	"log"

	"github.com/deadsy/go-cli"
)

func eval(s string) string {
	// log.Println("Eval")
	return s
}

func print(s string) string {
	// log.Println("Print")
	return s
}

func Run() {
	prompt := "user> "
	ln := cli.NewLineNoise()
	ln.SetMultiline(true)
	ln.HistorySetMaxlen(100)

	// reader := bufio.NewReader(os.Stdin)

	for {
		contents, err := ln.Read(prompt, "")
		if err != nil {
			log.Fatal(err)
		}
		ln.HistoryAdd(contents)
		ln.HistorySave("history.txt")

		tokens := read(contents)
		ast := eval(tokens)
		results := print(ast)
		fmt.Println(results)
	}
}
