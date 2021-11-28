// Paquet qui s'occupe du coeur du fonctionnement du client
package core

import (
	"fmt"
	"net"
)

func ConnectionWithServer(ipServer string, portServer string) {
	addressServer := fmt.Sprintf("%s:%s", ipServer, portServer)
	connectionServer, err := net.Dial("tcp", addressServer)
	if err == nil {
		fmt.Println("[+] Connexion avec le serveur réussie ! => ", connectionServer.RemoteAddr().String())
		connectionServer.Close()
	} else {
		fmt.Println("[-] Connexion avec le serveur impossible !")
	}
}
