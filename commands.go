package njord

import(
	"fmt"
)


/**
 * @brief      Main commands management function.
 *
 * @param      stdOut  The standard out
 *
 * @return
 */
func initializeCommandMap(stdOut chan string) {
	// fmt.Printf("þ Njord Initializing command map\r\n")
	stdOut<-"Njord Initializing command map"
	/**
	 * Make the commands has table
	 */
	commands = make(map[string]Command)

	help := help{name:"help", description:"Print help page", key:""}
	commands[help.name] = &help;

	exit := exit{name:"exit", description:"Shut down njord", key:""}
	commands[exit.name] = &exit;

	hecate := hecate{name:"hecate", description:"Testing TCP", key:""}
	commands[hecate.name] = &hecate;

	list := list{name:"list", description:"List other commands", key:""}
	commands[list.name] = &list;

	stdOut<-"Njord finished defining command map:"
	sumStr := fmt.Sprintf("%d commands initialized.", len(commands))
	stdOut<-sumStr

}

/**
 * @brief      The Command interface
 *
 * @param      done  The done
 *
 * @return
 */
type Command interface{
	njordExecute(done chan bool, stdOut chan string, args []string)
	njordStop(done chan bool, stdOut chan string)
	getName() string
	getDesc() string
}