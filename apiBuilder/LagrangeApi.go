package apiBuilder

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type DoApi interface {
	Do(client *websocket.Conn) error
}

type Request struct {
	Action string `json:"action"`
	Params struct {
		GroupID  int64  `json:"group_id,omitempty"`
		UserID   int64  `json:"user_id,omitempty"`
		Message  string `json:"message,omitempty"`
		Duration int    `json:"duration,omitempty"`
	} `json:"params"`
}

func (r *Request) BuildBody() ([]byte, error) {
	body, err := json.Marshal(r)
	return body, err
}

func (r *Request) Do(client *websocket.Conn) error {
	body, err := r.BuildBody()
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	// 发送 JSON 消息
	err = client.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		return err
	}
	return nil
}

func NewApi() IMainFuncApi {
	return &Request{}
}
