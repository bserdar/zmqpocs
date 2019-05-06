package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	ServerID string
	LastMsg  int
	Sub      *zmq.Socket
}

var clientID string

func (s *Server) NewMsg(msg int) {
	if s.LastMsg+1 != msg {
		fmt.Printf("%s: Out of order msg: %s %d\n", clientID, s.ServerID, msg)
	} else {
		fmt.Printf("%s: Receive %s: %d\n", clientID, s.ServerID, msg)
	}
	s.LastMsg = msg
}

// Run main with three args: clientId s1=endpoint,s2=endpoint,s3=endpoint...
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Run with args: clientId server1=endpoint,server2=endpoint,...")
		return
	}
	ctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}

	poller := zmq.NewPoller()
	servers := make(map[string]*Server)
	for _, x := range strings.Split(os.Args[2], ",") {
		ix := strings.IndexRune(x, '=')
		endpoint := x[ix+1:]
		sub, _ := ctx.NewSocket(zmq.SUB)
		err = sub.Connect(endpoint)

		if err != nil {
			panic(err)
		}
		sub.SetSubscribe("")
		servers[x[:ix]] = &Server{ServerID: x[:ix], Sub: sub}
		fmt.Printf("Listening to %s on %s\n", x[:ix], x[ix+1:])
		poller.Add(sub, zmq.POLLIN)
	}
	clientID = os.Args[1]

	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			for _, soc := range polled {
				msg, _ := soc.Socket.Recv(0)
				ix := strings.IndexRune(msg, ':')
				server := msg[:ix]
				data, _ := strconv.Atoi(msg[ix+1:])
				if s, ok := servers[server]; ok {
					s.NewMsg(data)
				}
			}
		}
	}
}
