package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServeur() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("New incoming client from client to orderbook feed : ", ws.RemoteAddr())
	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 4)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		// Read data from a fram of the websocket (message send by client from a port)
		// If the message is not long enought it will fill it
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("thank you for the message! "))
		s.broadcast(msg)
	}
}

// When a piece of bit is received loop over connected clients and send the message
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServeur()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/order", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(":3000", nil)

}

// from client side check client.js
