package parser

import (
	"regexp"
	"strings"
)

func ParseUnderScore(str string) string{
	return strings.Join(strings.Split(str, " "),"_")	
}

func RefParse(main *string){

	c := regexp.MustCompile("It can refer to:")
	*main = c.ReplaceAllLiteralString(*main, "")
}

func ParseData(data *string){
	*data = strings.TrimSpace(*data)

	re := regexp.MustCompile(`listen`)
	*data = re.ReplaceAllLiteralString(*data, "")

	re2 := regexp.MustCompile(`\(\)`)
	*data = re2.ReplaceAllLiteralString(*data, "")

	re3 := regexp.MustCompile(`\[([^\[\]]*)\]`)
	*data = re3.ReplaceAllLiteralString(*data, "")
	
}