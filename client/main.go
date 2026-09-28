package main

import (
	"bufio"
	"fmt"
	"net"
	"log"
)

func main() {
	MyClient()
}

func MyClient() {
	// 1.1
	// creer un listenner, en gros simplement un truc qui va ecouter sur le port 8080 ce qu il se passe
	dial, err := net.Dial("tcp", ":8080")
	// si erreur on stop
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()

	// send a message with the client
	message := "Hello server!"
	_, err = fmt.Fprintf(dial, message)
	if err != nil {
	   log.Fatal(err)
	}

	// Read the response from the server
	response, err := bufio.NewReader(dial).ReadString('\n')
	if err != nil {
	    fmt.Printf("Read error: %v\n", err)
	    return
	}

	fmt.Printf("Server response: %s", response)
}