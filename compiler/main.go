package main

import (
	"fmt"
	"os"
	"strings"
)

func newname(n string) string {
	sp := strings.Split(n, ".")
	count := len(sp)
	if count > 1 {
		sp = sp[:count-1]
	}

	left := ""
	for _, s := range sp {
		left += s
	}

	return left + ".bin"
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage:", os.Args[0], "<file to compile>")
	}
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		var scr Scanner
		words := scr.Scan(dat)
		fmt.Println("Scan pass completed")
		dumpWords(words)
		intermediate := parse(words)
		fmt.Println("\n\nParse pass completed")
		dumpWords(intermediate)
		//code := assemble(intermediate)
		//lib := link(code)
		//os.WriteFile(newname(os.Args[1]), []byte(lib), 0644)
	}
}

func dumpWords(words []WORD) {
	for _, w := range words {
		fmt.Println(w.symbol, ":", w)
	}
}
