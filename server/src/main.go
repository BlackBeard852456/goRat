// Paquet principal
package main

import "amolixs/core"

// Fonction principal
func main() {
	listener := core.Listening("127.0.0.1", "9090")
	core.ConnectionWithClient(listener)
}
