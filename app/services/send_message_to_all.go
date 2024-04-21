package services

import (
	"encoding/json"
	"momo-im/pkg/client"
	"momo-im/pkg/websocket"
)

func SendToAll(from, data string) {
	mgr := client.GetManagerInstance()

	message := websocket.NewMessage(from, data)
	resp := websocket.NewResponseFromMsg("1", "message", 200, "success", message)
	jsonResponse, _ := json.Marshal(resp)
	for _, c := range mgr.ListAllClients() {

		c.SendMsg(jsonResponse)
	}

}
