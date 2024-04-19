package client

import "sync"

type Manager struct {
	// Clients is a set of all the client's connections
	Clients     map[*Client]bool
	ClientsLock *sync.RWMutex
	// Users is a map of all the clients that are logged in
	Users      map[string]*Client
	UserLock   *sync.RWMutex
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}
