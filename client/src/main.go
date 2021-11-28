// Paquet principal
package main

import "amolixs/core"

// Fonction principal
func main() {
	connectionServer := core.ConnectionWithServer("127.0.0.1", "9090")
	defer connectionServer.Close()
	core.ReceiveDataOfServer(connectionServer)
}
