package parser

import (
	"regexp"
	"strings"
	"github.com/anaskhan96/soup"
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

	RefParse(data)

}

func ParsePage(root *soup.Root) string{

	var TextList []string

	for _, v := range root.FindAll("p"){
		txt := strings.TrimSpace(v.FullText())
		if !(txt=="" || len(strings.Split(txt," ")) < 20){
			TextList = append(TextList,v.FullText())
		}
	}
	if len(TextList)==0{
		return ""
	}
	return TextList[0]

}
