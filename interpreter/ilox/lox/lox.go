package lox

import (
	"bufio"
	"fmt"
	"ilox/lox/loxerr"
	"ilox/lox/scan"
	"ilox/lox/token"
	"os"
	"strings"
)

type Lox struct {
	err loxerr.Reporter
}

func New() Lox {
	return Lox{}
}

func usage() string {
	s := "Usage:\n"
	s += "  ilox                    runs REPL\n"
	s += "  ilox run [script]       runs the script\n"
	s += "  ilox format [script]    formats the script\n"
	s += "  ilox help\n"
	return s
}

func (l *Lox) Run(args []string) {
	switch len(args) {
	case 0:
		l.runPrompt()
		return
	case 1:
		switch args[0] {
		case "help":
			fmt.Println(usage())
			return
		}
	case 2:
		switch args[0] {
		case "run":
			l.runFile(args[1])
			return
		case "format":
			l.runFormatter(args[1])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unrecognized arguments: %q\n", strings.Join(args, " "))
	fmt.Fprintf(os.Stderr, usage())
	os.Exit(64)
}

func (l *Lox) runPrompt() {
	l.err.UseStdout = true
	fmt.Printf("> ")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			return
		}

		l.run(line)
		l.err.Reset()
		fmt.Printf("> ")
	}
}

func (l *Lox) runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(65)
	}

	l.run(string(bytes))
	if l.err.HasError() {
		os.Exit(65)
	}
}

func (l *Lox) runFormatter(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(65)
	}

	scanner := scan.NewScanner(string(bytes), &l.err)
	tokens := scanner.ScanTokens()
	if l.err.HasError() {
		os.Exit(65)
	}

	src := token.Format(tokens)
	fmt.Println(src)
}

func (l *Lox) run(src string) {
	scanner := scan.NewScanner(src, &l.err)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
