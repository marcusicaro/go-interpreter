package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := []rune(scanner.Text())
		l := lexer.New(line)
		for tok := l.NextToken(); string(tok.Type) != string(token.EOF); tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
