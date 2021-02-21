package main

import (
	"github.com/anaskhan96/soup"
	"fmt"
	"strings"
	"os"
	"github.com/icelain/cyclo/checks"
	"github.com/icelain/cyclo/candy"
	"github.com/icelain/cyclo/parser"
)

var (
	base = "https://en.wikipedia.org/wiki/"
)

func main(){
	checks.CheckArgs(os.Args)

	argument := parser.ParseUnderScore(strings.TrimSpace(os.Args[1]))
	cse := make(chan bool)
	spinner := candy.CreateSpinner()

	go candy.ShowSpinner(spinner, cse)

	s, err := soup.Get(base+argument)
	if err !=nil{
		fmt.Errorf(err.Error())
	}

	root := soup.HTMLParse(s)

	var TextList []string
	for _, v := range root.FindAll("p"){
		if !(strings.TrimSpace(v.FullText())==""){
			TextList = append(TextList,v.FullText())
		}
	}
	f := TextList[0]



	parser.ParseData(&f)
	parser.RefParse(&f)

	cse <- true
	fmt.Println("\n"+checks.CheckWikipediaErrs(f,os.Args[1]))
}
