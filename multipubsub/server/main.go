package main

import (
	"fmt"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// Run main with two args: endpoint and server ID
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Run with two args: endpoint serverID")
		return
	}
	endpoint := os.Args[1]
	serverID := os.Args[2]
	ctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}

	publisher, _ := ctx.NewSocket(zmq.PUB)
	err = publisher.Bind(endpoint)
	if err != nil {
		panic(err)
	}

	n := 0
	fmt.Printf("Starting server %s\n", serverID)
	for {
		msg := fmt.Sprintf("%s:%d", serverID, n)
		publisher.Send(msg, zmq.DONTWAIT)
		n++
		time.Sleep(100 * time.Millisecond)
	}
}
