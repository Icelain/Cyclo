package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/fatih/color"
	"github.com/icelain/cyclo/candy"
	"github.com/icelain/cyclo/checks"
	"github.com/icelain/cyclo/parser"
	"os"
	"strings"
)

func ErrOut(errProp string) {
	color.Red("No result found for: " + errProp)
	os.Exit(0)
}

var (
	base = "https://en.wikipedia.org/wiki/"
)

func main() {
	checks.CheckArgs(os.Args)

	argument := parser.ParseUnderScore(strings.TrimSpace(os.Args[1]))
	cse := make(chan struct{})
	spinner := candy.CreateSpinner()

	go candy.ShowSpinner(spinner, cse)

	s, err := soup.Get(base + argument)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	root := soup.HTMLParse(s)

	f := parser.ParsePage(&root)

	if f == "" {
		ErrOut(os.Args[1])
	}

	parser.ParseData(&f)

	cse <- struct{}{}
	fmt.Println("\n" + checks.CheckWikipediaErrs(f, os.Args[1]))
}
