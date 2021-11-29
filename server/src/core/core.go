// Paquet qui s'occupe du coeur du fonctionnement du serveur
package core

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Data struct {
	Name string
	Id   int
	Age  float32
}

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
	data := &Data{Name: "Amolixs", Id: 50, Age: 52.3}
	encoder := gob.NewEncoder(connectionClient)
	err := encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[+] Données envoyées au client")
	}
}

// Boucle principal du programme
func mainLoop() {
	loopControl := true
	for loopControl {
	}
}
