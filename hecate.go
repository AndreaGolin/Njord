package njord

import(
	"log"
)

type hecate struct{
	name string
	description string
	key	string
}

/**
 * @brief      implement Command interface
 */
func (s hecate) njordExecute(done chan bool,stdOut chan string) {
	log.Printf("þ Hecate executing!\r\n")
	startHecate(done)
	
}
func (s hecate) njordStop(done chan bool,stdOut chan string){
	log.Printf("þ Hecate stopping.\n\r")
}
func (s hecate) getName() string{
	return s.name
}
func (s hecate) getDesc() string{
	return s.description
}

func startHecate(done chan bool){
	log.Printf("þ Hecate bootstrapping....\r\n")
}