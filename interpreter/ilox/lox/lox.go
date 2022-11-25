package lox

import (
	"bufio"
	"fmt"
	"ilox/lox/loxerr"
	"ilox/lox/scan"
	"log"
	"os"
)

type Lox struct {
	errRep loxerr.Reporter
}

func New() Lox {
	return Lox{}
}

func (l *Lox) Run(args []string) {
	if len(args) > 1 {
		fmt.Println("Usage: ilox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		if err := l.runFile(args[0]); err != nil {
			log.Fatal(err)
		}
	} else {
		l.errRep.UseStdout = true
		if err := l.runPrompt(); err != nil {
			log.Fatal(err)
		}
	}
}

func (l *Lox) runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	l.run(string(bytes))
	if l.errRep.HasError() {
		os.Exit(65) // TODO: Should we really exit from here?
	}
	return nil
}

func (l *Lox) runPrompt() error {
	fmt.Printf("> ")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			return nil
		}

		l.run(line)
		l.errRep.Reset()
		fmt.Printf("> ")
	}

	return nil
}

func (l *Lox) run(src string) {
	scanner := scan.NewScanner(src, &l.errRep)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
