package websocket

import (
	"encoding/json"
	"testing"
)

func TestNewResponseFromMsg(t *testing.T) {
	message := NewMessage("test", "test")
	response := NewResponseFromMsg("1", "test", 200, "success", message)

	j, _ := json.Marshal(response)
	t.Log(string(j))
}
