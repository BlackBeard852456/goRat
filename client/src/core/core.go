// Paquet qui s'occupe du coeur du fonctionnement du client
package core

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strings"
)

type Data struct {
	Name string
	Id   int
	Age  float32
}

type Command struct {
	cmdOutput string
	cmdError  string
}

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
	decoder := gob.NewDecoder(connectionServer)
	data := &Data{}
	err := decoder.Decode(data)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[+] Donnée envoyé par le serveur bien recus !")
		fmt.Println(data.Name)
	}
}

func ExecuteCommand(connectionServer net.Conn) {
	reader := bufio.NewReader(connectionServer)
	commandLoop := true
	for commandLoop {
		command, err := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")
		if err != nil {
			fmt.Println(err)
		}
		if command == "stop" {
			commandLoop = false
		}
	}
}
