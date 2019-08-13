package repl

import (
	"bufio"
	"fmt"
	"github.com/hajjboy95/interpreter-go/lexer"
	"github.com/hajjboy95/interpreter-go/token"
	"io"
)

const Prompt = "$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}