package websocket

type Response struct {
	Header
	Body Body `json:"body"`
}

type Header struct {
	Seq string `json:"seq"`
	Cmd string `json:"cmd"`
}

type Body struct {
	Code    uint32 `json:"code"`
	CodeMsg string `json:"codeMsg"`
	Data    any    `json:"data"`
}

func NewResponseFromMsg(seq string, cmd string, code uint32, codeMsg string, message *Message) *Response {
	return &Response{
		Header: Header{
			Seq: seq,
			Cmd: cmd,
		},
		Body: Body{
			Code:    code,
			CodeMsg: codeMsg,
			Data:    message,
		},
	}
}
