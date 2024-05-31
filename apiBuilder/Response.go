package apiBuilder

import (
	"encoding/json"
)

type Status struct {
	Status        string          `json:"status"`
	RetCode       int             `json:"retcode"`
	Data          Data            `json:"data,omitempty"`
	DataInterface json.RawMessage `json:"dataInterface,omitempty"`
	Echo          any             `json:"echo"`
}

type Data struct {
	GroupID      int64  `json:"group_id,omitempty"`
	UserID       int64  `json:"user_id,omitempty"`
	MessageID    int64  `json:"message_id,omitempty"`
	ForwardID    string `json:"forward_id,omitempty"`
	NickName     string `json:"nickname,omitempty"`
	Card         any    `json:"card,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Age          int    `json:"age,omitempty"`
	Area         string `json:"area,omitempty"`
	JoinTime     int64  `json:"join_time,omitempty"`
	LastSentTime int64  `json:"last_sent_time,omitempty"`
	Level        string `json:"level,omitempty"`
	Role         string `json:"role,omitempty"`
	UnFriendly   bool   `json:"unfriendly,omitempty"`
	Title        string `json:"title,omitempty"`
	Time         int64  `json:"time,omitempty"`
	MessageType  string `json:"message_type,omitempty"`
	RealID       int64  `json:"real_id,omitempty"`
	Sender       struct {
		UserID   int64  `json:"user_id,omitempty"`
		NickName string `json:"nickname,omitempty"`
		Card     string `json:"card,omitempty"`
		Sex      string `json:"sex,omitempty"`
		Age      int    `json:"age,omitempty"`
		Area     string `json:"area,omitempty"`
		Level    string `json:"level,omitempty"`
		Role     string `json:"role,omitempty"`
		Title    string `json:"title,omitempty"`
	} `json:"sender,omitempty"`
	Message []struct {
		Type string `json:"type"`
		Data struct {
			Text string `json:"text,omitempty"`
			File string `json:"file,omitempty"`
			Url  string `json:"url,omitempty"`
			Data string `json:"data,omitempty"`
			QQ   string `json:"qq,omitempty"`
			ID   string `json:"id,omitempty"`
		}
	} `json:"message,omitempty"`
}

type Response struct {
	originMsg []byte
	response  *Status
}

func NewResponse(msg []byte) *Response {
	return &Response{
		originMsg: msg,
	}
}

func (r *Response) unmarshal() error {
	if r.response == nil {
		r.response = &Status{}
		err := json.Unmarshal(r.originMsg, r.response)
		return err
	}
	return nil
}

func (r *Response) GetData(data interface{}) error {
	if err := r.unmarshal(); err != nil {
		return err
	}
	return json.Unmarshal(r.response.DataInterface, data)
}

func (r *Response) Ok() bool {
	if err := r.unmarshal(); err != nil {
		return false
	}
	return r.response.Status == "ok"
}

func (r *Response) StatusMsg() string {
	if err := r.unmarshal(); err != nil {
		return ""
	}
	return r.response.Status
}

func (r *Response) Result() (status string, retCode int) {
	if err := r.unmarshal(); err != nil {
		return "", 0
	}
	return r.response.Status, r.response.RetCode
}

func (r *Response) GetOrigin() []byte {
	return r.originMsg
}

func (r *Response) GetGroupMessageResponse() (*Response, error) {
	if err := r.unmarshal(); err != nil {
		return nil, err
	}
	data := &Response{}
	err := json.Unmarshal(r.response.DataInterface, data)
	return data, err
}
