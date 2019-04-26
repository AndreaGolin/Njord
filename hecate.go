package njord

import(
	"fmt"
)

type hecate struct{
	name string
	description string
	key	string
}

/**
 * @brief      implement Command interface
 */
func (s hecate) njordExecute(done chan bool,stdOut chan string, args []string) {
	stdOut<-"Executing Hecate"
	startHecate(done, stdOut, args)
	
}
func (s hecate) njordStop(done chan bool,stdOut chan string){
	stdOut<-"Stopping Hecate"
}
func (s hecate) getName() string{
	return s.name
}
func (s hecate) getDesc() string{
	return s.description
}

func startHecate(done chan bool, stdOut chan string, args []string){
	for k,a := range args{
		fmt.Printf("%d -> %s\n", k,a)
	}
}