package websocket

import "encoding/json"

type MessageType string

func (m *MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*m))
}

func (m *MessageType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*m = MessageType(s)
	return nil
}

const (
	MessageTypeText  MessageType = "text"
	MessageTypeEnter MessageType = "enter"
	MessageTypeExit  MessageType = "exit"
)

type Message struct {
	Target string      `json:"target"`
	Type   MessageType `json:"type"`
	Msg    string      `json:"msg"`
	From   string      `json:"from"`
}

func NewMessage(from, msg string) *Message {
	return &Message{
		Type: MessageTypeText,
		From: from,
		Msg:  msg,
	}
}

type MessageController struct {
}

func (mc *MessageController) GetMsgID() string {}
