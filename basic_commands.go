package njord

import(
	"log"
	"fmt"
)

type help struct{
	name string
	description string
	key	string
}
type exit struct{
	name string
	description string
	key	string
}
type list struct{
	name string
	description string
	key	string
}


/**
 * HELP Command
 * @brief      implement Command interface
 */
func (s help) njordExecute(done chan bool, stdOut chan string) {
	log.Printf("þ Help executing!\r\n")
	CallClear()
	PrintHelp(s, done, stdOut)
}
func (s help) njordStop(done chan bool, stdOut chan string){
	log.Printf("þ Help stopping.\n\r")
	done <- true
}
func (s help) getName() string{
	return s.name
}
func (s help) getDesc() string{
	return s.description
}

/**
 * @brief      Print help prompt
 */
func PrintHelp(s help, done chan bool, stdOut chan string) {
	log.Println("Printing help.....\r\n")
	s.njordStop(done, stdOut)

}


/**
 * EXIT Command
 *
 * @brief      implement Command interface
 */
func (s exit) njordExecute(done chan bool,stdOut chan string) {
	log.Printf("þ Exit executing!\r\n")
	done <- true
}
func (s exit) njordStop(done chan bool,stdOut chan string){
	log.Printf("þ Exit stopping.\n\r")
	done <- true
}
func (s exit) getName() string{
	return s.name
}
func (s exit) getDesc() string{
	return s.description
}

/**
 * LIST commands
 */
func (s list) njordExecute(done chan bool, stdOut chan string) {
	log.Printf("þ List executing!\r\n")

	/**
	 * Iterate the commands hashmap and print name and description
	 */
	for _, c := range commands{
		name := c.getName()
		desc := c.getDesc()
		fmt.Println("");
		fmt.Println("*************")
		fmt.Printf("Command name: %s\r\n", name)
		fmt.Printf("Command desc: %s\r\n", desc)
		fmt.Println("*************")
		fmt.Println("");
	}

}
func (s list) njordStop(done chan bool, stdOut chan string){
	log.Printf("þ List stopping.\n\r")
	done <- true
}
func (s list) getName() string{
	return s.name
}
func (s list) getDesc() string{
	return s.description
}