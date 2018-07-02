package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/KeKsBoTer/durak/entity"
	"github.com/gorilla/websocket"
)

var hub Hub

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func queue(hub *Hub, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "player")
	if err != nil {
		log.Println(err)
		http.Error(w, "somthing went wrong", http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	username := session.Values["name"].(string)
	user := entity.User(username)
	if hub.queue.ContainsUser(user) {
		log.Printf("user %v allready in queue\n", user)
		fmt.Fprintf(w, "your are allready in the queue")
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
		user: user,
	}
	hub.join <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.write()
	go client.read()
}

// PlayerQueue holds all users, who are waiting for a game.
type PlayerQueue struct {
	root *queueNode
	size int
	mux  sync.Mutex
}

type queueNode struct {
	client   *Client
	messages chan string
	next     *queueNode
}

// Size returns the amount of useres in the queue
func (q *PlayerQueue) Size() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	return q.size
}

// Append adds user to end of queue
func (q *PlayerQueue) Append(c *Client) {
	q.mux.Lock()
	defer q.mux.Unlock()
	newNode := &queueNode{
		client: c,
	}
	q.size++
	if q.root == nil {
		q.root = newNode
		return
	}
	var node = q.root
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

// Pop removes first user from queue and returns it
// if queue is empty, nil is returned
func (q *PlayerQueue) Pop() *Client {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.root == nil {
		return nil
	}
	client := q.root.client
	q.root = q.root.next
	q.size--
	return client
}

// Remove removes user form queue
func (q *PlayerQueue) Remove(c *Client) {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.root == nil {
		return
	}
	q.size--
	if q.root.client == c {
		q.root = q.root.next
		return
	}
	last, node := q.root, q.root.next
	for node.client != c {
		last, node = node, node.next
		if node == nil {
			return
		}
	}
	last.next = node.next
}

func (q *PlayerQueue) forEach(fn func(*Client)) {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.root == nil {
		return
	}
	node := q.root
	for ; node != nil; node = node.next {
		fn(node.client)
	}
}

// ContainsUser checks if queue contains user
func (q *PlayerQueue) ContainsUser(u entity.User) bool {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.root == nil {
		return false
	}
	node := q.root
	for ; node != nil; node = node.next {
		if node.client.user == u {
			return true
		}
	}
	return false
}
