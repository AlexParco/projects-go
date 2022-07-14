package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/alexparco/tcp-chat/client"
)

type Server struct {
	Clients []*client.Client
	Msg     chan string
	Lock    *sync.Mutex
}

func New() *Server {
	s := &Server{
		Clients: []*client.Client{},
		Msg:     make(chan string),
		Lock:    &sync.Mutex{},
	}
	return s
}

func (s *Server) Run(addr int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", addr))
	if err != nil {
		panic(err)
	}
	fmt.Printf("TCP CHAT RUN IN 127.0.0.1:%d\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go s.handlerConnection(conn)
	}
}

func (s *Server) handlerConnection(conn net.Conn) {

	cl := client.New(conn)

	s.Lock.Lock()
	s.Clients = append(s.Clients, cl)
	s.Lock.Unlock()

	fmt.Printf("New connection: %s\n", cl.Addr)
	s.Write(cl)
}

func (s *Server) Write(client *client.Client) {
	for msg := range client.Msg {
		s.broadcast(msg, client.Addr)
	}
}

func (s *Server) broadcast(msg string, addr net.Addr) {
	for _, client := range s.Clients {
		if client.Addr != addr {
			client.Write(msg)
		}
	}
}
