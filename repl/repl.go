package repl

import (
	"bufio"
	"fmt"
	lexer "github.com/radekBednarik/waiigo/lexer"
	token "github.com/radekBednarik/waiigo/token"
	"io"
	"os"
)

const PROMPT = ">> "
const EXIT = "exit()"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == EXIT {
			os.Exit(0)
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}

}
