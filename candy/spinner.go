package candy

import (
	"github.com/briandowns/spinner"
	"time"
)

func CreateSpinner() *spinner.Spinner{
	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Color("cyan")
	return s
}

func ShowSpinner(s *spinner.Spinner, spcase chan bool){
	s.Start()
	<- spcase
	s.Stop()
}