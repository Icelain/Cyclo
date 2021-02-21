package checks

import (
	"github.com/fatih/color"
	"os"
)

const (
	VERSION = "v1.0"
	AUTHOR = "Icelain"
 )


func CheckArgs(args []string){

	if len(args)==1{
		color.Red("No search term specified")
		os.Exit(0)
	}

	v := args[1]

	switch true{

	case (v == "-h" || v == "--help"):
		color.Yellow("Cyclo "+ VERSION + " by "+ AUTHOR + "\nHelp: --help or -h\nVersion: --version or v\nAuthor: --author or -a")
		os.Exit(0)

	case (v == "-v" || v == "--version"):
		color.Yellow(VERSION)
		os.Exit(0)

	case (v == "-a" || v == "--author"):
        color.Yellow(AUTHOR)
        os.Exit(0)

	}

}

func CheckWikipediaErrs(s string, rse string) string{

	if (s == "Other reasons this message may be displayed:"){
		return color.RedString("No result found for: ") + color.MagentaString(rse)
	}
	return color.GreenString(s)
}
