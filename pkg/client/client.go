package client

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Addr            string
	Socket          *websocket.Conn // Socket is the websocket connection
	Send            chan []byte     // Send is a channel to hold the data that is to be sent to the client
	AppID           uint32
	UserID          string
	ConnectedAt     uint64
	LoginAt         uint64
	LastHeartBeatAt uint64
}

func NewClient(addr string, socket *websocket.Conn) *Client {
	now := uint64(time.Now().Unix())
	return &Client{
		Addr:            addr,
		Socket:          socket,
		Send:            make(chan []byte, 100),
		ConnectedAt:     now,
		LoginAt:         now,
		LastHeartBeatAt: now,
	}
}

func (c *Client) Key() string {
	return fmt.Sprintf("%d_%s", c.AppID, c.UserID)
}

func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("client send message stop", string(debug.Stack()), r)
		}
	}()

	c.Send <- msg
}

// Write is a goroutine that waits for the data to be written. Once data is available,
// it writes (push) the data to the client's websocket connection.
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()
	for message := range c.Send {
		_ = c.Socket.WriteMessage(websocket.TextMessage, message)
	}

	fmt.Println("client write channel is closed, closing the connection...", c.Addr)
	_ = c.Socket.Close()
}
