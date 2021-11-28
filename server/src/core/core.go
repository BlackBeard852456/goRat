// Paquet qui s'occupe du coeur du fonctionnement du serveur
package core

import (
	"fmt"
	"log"
	"net"
)

// Permet de lancer le serveur
func Listening(ip string, port string) net.Listener {
	listenAddress := fmt.Sprintf("%s:%s", ip, port)
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	return listener
}

func ConnectionWithClient(listener net.Listener) {
	clientConnection, err := listener.Accept()
	defer clientConnection.Close()
	if err != nil {
		fmt.Println("[-] Connexion avec le client non établie !")
	} else {
		fmt.Println("[+] Connexion avec le client établie => ", clientConnection.RemoteAddr().String())
	}
}
