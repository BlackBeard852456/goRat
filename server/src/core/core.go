// Paquet qui s'occupe du coeur du fonctionnement du serveur
package core

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

// Permet d'accepter une connexion du client
func ConnectionWithClient(listener net.Listener) net.Conn {
	clientConnection, err := listener.Accept()
	if err != nil {
		fmt.Println("[-] Connexion avec le client non établie !")
	} else {
		fmt.Println("[+] Connexion avec le client établie => ", clientConnection.RemoteAddr().String())
	}
	return clientConnection
}

// Permet d'écrire des données au client
func WriteDataToClient(connectionClient net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	dataToSend, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	connectionClient.Write([]byte(dataToSend))
}
