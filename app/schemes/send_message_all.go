package schemes

type SendMessageAllRequestBody struct {
	AppID  uint32 `json:"app_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
	MsgID  string `json:"msg_id" binding:"required"`
	Msg    string `json:"msg" binding:"required"`
}
