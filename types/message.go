package types

type PostMessageResp struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
		MessageID int `json:"message_id"`
	} `json:"data"`
	Message string       `json:"message"`
	Wording string       `json:"wording"`
	Echo    *interface{} `json:"echo,omitempty"`
}

type BotPostMessageData struct {
	SelfID      int64  `json:"self_id"`
	UserID      int64  `json:"user_id"`
	Time        int64  `json:"time"`
	MessageID   int64  `json:"message_id"`
	MessageSeq  int64  `json:"message_seq"`
	RealID      int64  `json:"real_id"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
	} `json:"sender"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	SubType    string `json:"sub_type"`
	Message    []struct {
		Data struct {
			QQ   string `json:"qq,omitempty"`
			Text string `json:"text,omitempty"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
	GroupID       int64  `json:"group_id"`
}
