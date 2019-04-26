package njord

import(
	"log"
	"github.com/google/gopacket"
	//"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"fmt"
)

type sonar struct{
	name string
	description string
	key	string
}

/**
 * @brief      implement Command interface
 */
func (s sonar) njordExecute(done chan bool, stdOut chan string) {
	log.Printf("þ Sonar executing!\r\n")
	captureWithPcap()
}
func (s sonar) njordStop(done chan bool, stdOut chan string){
	log.Printf("þ Sonar stopping.\n\r")
	done <- true
}
func (s sonar) getName() string{
	return s.name
}
func (s sonar) getDesc() string{
	return s.description
}

/**
 * @brief      Capture live packets off the wire with pcap
 *
 * @param      closeChan2  The close channel 2
 *
 * @return     
 */
func captureWithPcap(){

	
	handle, err := pcap.OpenLive("any", 1024, false, pcap.BlockForever);
	if  err != nil {
	  panic(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    for packet := range packetSource.Packets() {
    	fmt.Println(packet)
    }

    
	log.Println("with syscall end")

}