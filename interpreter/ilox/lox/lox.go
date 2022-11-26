package lox

import (
	"bufio"
	"fmt"
	"ilox/lox/loxerr"
	"ilox/lox/scan"
	"ilox/lox/token"
	"log"
	"os"
	"strings"
)

type Lox struct {
	errRep loxerr.Reporter
}

func New() Lox {
	return Lox{}
}

func (l *Lox) Run(args []string) {
	unrecognized := func() {
		fmt.Fprintf(os.Stderr, "Unrecognized arguments: %q\n", strings.Join(args, " "))
		fmt.Fprintf(os.Stderr, usage())
		os.Exit(64)
	}

	switch len(args) {
	case 0:
		// REPL
		l.errRep.UseStdout = true
		if err := l.runPrompt(); err != nil {
			log.Fatal(err)
		}
	case 1:
		// help or bad arg
		if args[0] == "help" {
			fmt.Println(usage())
			os.Exit(0)
		} else {
			unrecognized()
		}
	case 2:
		path := args[1]
		switch args[0] {
		case "run":
			if err := l.runFile(path); err != nil {
				log.Fatal(err)
			}
		case "format":
			if err := l.runFormatter(path); err != nil {
				log.Fatal(err)
			}
		default:
			unrecognized()
		}
	default:
		unrecognized()
	}
}

func usage() string {
	s := "Usage:\n"
	s += "  ilox                        runs REPL\n"
	s += "  ilox run [script]           runs the script\n"
	s += "  ilox format [script]        formats the script\n"
	s += "  ilox help\n"
	return s
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

func (l *Lox) runFormatter(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	scanner := scan.NewScanner(string(bytes), &l.errRep)
	tokens := scanner.ScanTokens()

	if l.errRep.HasError() {
		os.Exit(65) // TODO: Should we use the same value as runFile?
	}

	src := token.Format(tokens)
	fmt.Println(src)

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
