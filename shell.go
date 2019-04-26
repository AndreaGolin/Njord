package njord

import (
	"os"
	"os/signal"
	"bufio"
	"strings"
	"fmt"
	"os/exec"
	"runtime"
	"flag"
)

var clear map[string]func()
var commands map[string]Command

func Init(args []string){
	
	done := make(chan bool)
	stdOut := make(chan string)

	go listenForStdout(stdOut)
	
	initializeCommandMap(stdOut)

	commandFlgPtr 		:= flag.String("c", "", "Command to execute")
	interactiveFlgPtr 	:= flag.Bool("i", false, "Interactive Mode")
	flag.Parse()

	if *commandFlgPtr != ""{
		execCommandWrapper(*commandFlgPtr, done, stdOut, nil)
	}


	if *interactiveFlgPtr == true{

		stdOut<-"Njord shell starting"

		clearInit()

		go listenForSigs(done)

		go listenForCommands(done, stdOut)

		for{
			select{
			case <-done:
				fmt.Printf("þ Stopping Njord, bye!\r\n")
				return
			}
		}
	}

}

func listenForStdout(stdOut chan string){
	for{
		select{
		case msg := <-stdOut:
			fmt.Printf("þ %s\n", msg)
		}
	}
}

/**
 * @brief      Start a loop listening for stdin commands
 *
 * @return
 */
func listenForCommands(done chan bool, stdOut chan string) {
	reader := bufio.NewReader(os.Stdin)
	for{
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		slices := strings.Fields(text)
		if len(slices) > 0{
			execCommandWrapper(slices[0], done, stdOut, slices[1:])
		}
	}
}

/**
 * @brief      Wrapper mini-method for executing commands
 *
 * @param      commandName  The command name
 * @param      done         The done
 *
 * @return     
 */
func execCommandWrapper(commandName string, done chan bool, stdOut chan string, args []string){

	if command, ok := commands[commandName]; ok{
		command.njordExecute(done, stdOut, args)
	}else{
		stdOut<-fmt.Sprintf("Command %s not found.\r\n", commandName)
	}
}

/**
 * @brief      { function_description }
 *
 * @param      chan bool the kill channel
 */
func listenForSigs(done chan bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case s := <-c:
			fmt.Printf("\r\nþ Received os signal: %s\r\n", s.String())
			done <- true
			return
		}
	}
}


/**
 * @brief      console clear functions
 */
func clearInit(){
	clear = make(map[string]func())
	clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}
func CallClear(){
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Platfom unknown, Njord can not clear the console.")
    }
}
