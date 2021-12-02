// Paquet qui s'occupe du coeur du fonctionnement du serveur
package core

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
func MainLoop(connectionClient net.Conn) {
	loopControl := true
	for loopControl {
		userOptionInput, err := retrieveChoiceOption()
		if err != nil {
			continue
		}
		manageChoiceOption(userOptionInput, connectionClient, &loopControl)
		connectionClient.Write([]byte(userOptionInput))
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
	userOptionInput = strings.TrimSuffix(userOptionInput, "\n")
	return userOptionInput, err
}

// Gère la gestion du choix de l'option de l'utilisateur
func manageChoiceOption(userOptionInput string, connectionClient net.Conn, loopControl *bool) {
	switch userOptionInput {
	case "1":
		fmt.Println("[+] Command Execution Program")
	case "99":
		fmt.Println("[-] Exiting programm")
		*loopControl = false
	default:
		fmt.Println("[-] Invalid Option Input !")
	}
}

// Permet d'éxecuter une commande à distance du client
func executeCommandRemotly(connectionClient net.Conn) (err error) {
	return err
}
