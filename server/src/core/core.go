// Paquet qui s'occupe du coeur du fonctionnement du serveur
package core

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
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
func MainLoop() {
	loopControl := true
	for loopControl {
		userOptionInput, err := retrieveChoiceOption()
		if err != nil {
			continue
		}
		fmt.Println(userOptionInput)
	}
}

// Permet de récupérer le choix de l'utilisateur
func retrieveChoiceOption() (userOptionInput string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("[+] Enter Options")
	userOptionInput, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return userOptionInput, err
}
