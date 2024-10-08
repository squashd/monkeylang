package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/squashd/monkey/evaluator"
	"github.com/squashd/monkey/lexer"
	"github.com/squashd/monkey/object"
	"github.com/squashd/monkey/parser"
)

func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {

			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
