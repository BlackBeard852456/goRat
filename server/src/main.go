// Paquet principal
package main

import "amolixs/core"

// Fonction principal
func main() {
	listener := core.Listening("127.0.0.1", "9090")
	clientConnection := core.ConnectionWithClient(listener)
	defer clientConnection.Close()
	core.WriteDataToClient(clientConnection)
	core.MainLoop(clientConnection)
}
