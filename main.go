package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	re "rocks_master/error"
)

const (
	AGENT_ADDR = "localhost"
	AGENT_PORT = "7878"
	AGENT_TYPE = "udp"
)

func main() {
		service := AGENT_ADDR + ":" + AGENT_PORT

		// Creating Remote Address Information field + Error Handling
		RemoteAddr, err := net.ResolveUDPAddr(AGENT_TYPE, service)
		re.Encounter(re.ErrorInformation{ErrorType: "FATAL", Error: err})

		// Dialing AGENT via UDP + Error Handling
		conn, err := net.DialUDP(AGENT_TYPE, nil, RemoteAddr)
		re.Encounter(re.ErrorInformation{ErrorType: "FATAL", Error: err})
		defer conn.Close()

		// getting metric from cli argument and transmitting it to the client + Error Handling
		input := strings.TrimSuffix(os.Args[1], "\n")
		message := []byte(input)
		_, err = conn.Write(message)
		re.Encounter(re.ErrorInformation{ErrorType: "FATAL", Error: err})

		// Creating Buffer and writing the response of the ROCKS Agent to it
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)

		//Printing the Metric Result
		fmt.Println(string(buffer[:n]))
}