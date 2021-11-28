// Paquet qui s'occupe du coeur du fonctionnement du client
package core

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func ConnectionWithServer(ipServer string, portServer string) net.Conn {
	addressServer := fmt.Sprintf("%s:%s", ipServer, portServer)
	connectionServer, err := net.Dial("tcp", addressServer)
	if err == nil {
		fmt.Println("[+] Connexion avec le serveur réussie ! => ", connectionServer.RemoteAddr().String())
	} else {
		fmt.Println("[-] Connexion avec le serveur impossible !")
	}

	return connectionServer
}

// Permet de recevoir les données du serveur
func ReceiveDataOfServer(connectionServer net.Conn) {
	reader := bufio.NewReader(connectionServer)
	dataReceived, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("[-] Impossible de lire les données envoyées par le serveur !")
	}
	dataReceived = strings.TrimSuffix(dataReceived, "\n")
	fmt.Println(dataReceived)
}
