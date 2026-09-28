package main

import (
	"io"
	"log"
	"net"
)

func main() {
	MyListenner()
}

func MyListenner() {
	// 1.1
	// creer un listenner, en gros simplement un truc qui va ecouter sur le port 8080 ce qu il se passe
	listener, err := net.Listen("tcp", ":8080")
	// si erreur on stop
	if err != nil {
		log.Fatal(err)
	}
	// 1.2
	// defer sert a nettoyer des ressources, il s execute instanement mais n est resolu quand la fin
	defer listener.Close()
	// on va ecouter les connections en continue sur le port
	for {
		// 1.3
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 1.4
		// handle ce qui se passe pendant la connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	io.Copy(conn, conn)
	conn.Close()
	// faire des trucs !

}
