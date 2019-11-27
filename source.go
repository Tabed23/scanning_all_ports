// scann all ports

package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var (
	ip      = "27.255.29.55" // port to scan
	minport = 1
	maxport = 1024
)

func main() {
	activethreads := 0
	donechannals := make(chan bool)
	for port := minport; port <= maxport; port++ {
		go testTCPpConnection(ip, port, donechannals)
		activethreads++
	}
	// waite all threads to finished
	if activethreads > 0 {
		<-donechannals
		activethreads--
	}
}
func testTCPpConnection(ip string, port int, donechannals chan bool) {

	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		log.Printf("Host %s has open port %d\n", ip, port)
	}
	donechannals <- true
}
