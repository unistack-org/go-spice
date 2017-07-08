package main

import (
	"net"

	spice "github.com/vtolstov/go-spice"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6789")
	if err != nil {
		panic(err)
	}

	sp, err := spice.Connect(conn)
	if err != nil {
		panic(err)
	}
	defer sp.Close()

}
