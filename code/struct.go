package code

// PersonInfo 用户接受消息结构体
type PersonInfo struct {
	From         string `json:"from"`
	EventID      uint32 `json:"event_id"`
	EventContent []byte `json:"event_content"`
	Description  string `json:"description"`
}
