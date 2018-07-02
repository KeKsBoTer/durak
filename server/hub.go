package server

import (
	"fmt"
	"log"
)

type Hub struct {
	// Registered clients.
	queue PlayerQueue

	// Register requests from the clients.
	join chan *Client

	// Unregister requests from clients.
	leave chan *Client

	// Inbound messages for clients.
	messages chan []byte
}

func newHub() *Hub {
	return &Hub{
		messages: make(chan []byte),
		join:     make(chan *Client),
		leave:    make(chan *Client),
		queue:    PlayerQueue{},
	}
}

func (h *Hub) broadcastUserCount() {
	h.messages <- []byte(fmt.Sprintf("%d players in queue", h.queue.size))
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.join:
			h.queue.Append(client)
			log.Printf("user %v joined queue", client.user)
			// do async so that the thread is not blocked
			go h.broadcastUserCount()
		case client := <-h.leave:
			h.queue.Remove(client)
			log.Printf("user %v left queue", client.user)
			go h.broadcastUserCount()
		case message := <-h.messages:
			h.queue.forEach(func(client *Client) {
				select {
				case client.send <- message:
				default:
					h.queue.Remove(client)
				}
			})
		}
	}
}
