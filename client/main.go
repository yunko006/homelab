package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	MyClient()
}

// un client TCP envoie uniquement des bytes (pas de string/int etc)
// donc on doit convertir les datas que l'on veux
// envoyer sous forme de byte
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
	fmt.Print("Envoyer un message a lire par le server: ")
	var message string
	fmt.Scanln(&message)
	// ici je dois convertir le message en bytes afin de pouvoir l'envoyer au server qui lui va le decoder
	dial.Write([]byte(message))
}

// client puisse send un message
func SendMessageFromClient(message string) {
	// message := "Hello server!"

	fmt.Println("Your text was:", message)

}
