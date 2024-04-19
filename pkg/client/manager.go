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

func (mgr *Manager) ClientIsExists() bool {
	return false
}

func (mgr *Manager) ListAllClients() []*Client {
	res := make([]*Client, 0)
	mgr.ClientsLock.RLock()
	defer mgr.ClientsLock.RUnlock()

	for client, value := range mgr.Clients {
		if value {
			res = append(res, client)
		}
	}
	return res
}

func (mgr *Manager) AddClient(client *Client) {
	mgr.ClientsLock.Lock()
	defer mgr.ClientsLock.Unlock()
	mgr.Clients[client] = true
}

// DelClient deletes a client from the manager
func (mgr *Manager) DelClient(client *Client) bool {
	mgr.ClientsLock.Lock()
	defer mgr.ClientsLock.Unlock()
	if _, ok := mgr.Clients[client]; ok {
		delete(mgr.Clients, client)
		return true
	}
	return false
}

func (mgr *Manager) AddUser(key string, client *Client) {
	mgr.UserLock.Lock()
	defer mgr.UserLock.Unlock()
	mgr.Users[key] = client
}

func (mgr *Manager) DelUser(key string) bool {
	mgr.UserLock.Lock()
	defer mgr.UserLock.Unlock()
	if _, ok := mgr.Users[key]; ok {
		delete(mgr.Users, key)
		return true
	}
	return false
}
