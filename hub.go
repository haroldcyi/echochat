// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)



// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}


/////////////////////////// by harold

// Pool of active hubs.
var hubs = make(map[string]*Hub)


// Fetch a room record from the DB and initialize it in memory.
func initializeHub(id string) (*Hub, error) {
	// data := struct {
	// 	Password []byte `redis:"password"`
	// }{}

	// db := dbPool.Get()
	// defer db.Close()

	// exists, err := db.GetMap(config.CachePrefixRoom+id, &data)
	// if err != nil {
	// 	return nil, errors.New("Error loading room")
	// } else if !exists {
	// 	return nil, nil
	// }

	// hubs[id] = &Hub{
	// 	Id:             id,
	// 	password:       data.Password,
	// 	timestamp:      time.Now(),
	// 	broadcastQueue: make(chan []byte),
	// 	register:       make(chan *Peer),
	// 	unregister:     make(chan *Peer),
	// 	stop:           make(chan int),
	// 	peers:          make(map[*Peer]bool),
	// 	counter:        1,
	// }

	hubs[id] = newHub()

	log.Println("hub id...")
	log.Println(id)
	log.Println("end.")

	// Hub is loaded, start it.
	go hubs[id].run()

	return hubs[id], nil
}




// Get an initialized hub by id.
func getHub(id string) *Hub {
	hub, exists := hubs[id]

	if exists {
		return hub
	} else {
		return nil
	}
}























