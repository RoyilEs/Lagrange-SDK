package apiBuilder

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type DoApi interface {
	Do() error
}

type Request struct {
	Action string `json:"action"`
	Params struct {
		GroupID int64  `json:"group_id"`
		UserID  int64  `json:"user_id"`
		Message string `json:"message"`
	}
}

func (r *Request) BuildStringBody() ([]byte, error) {
	body, err := json.Marshal(r)
	return body, err
}

func (r *Request) Do() error {
	body, err := r.BuildStringBody()
	if err != nil {
		return err
	}
	var conn *websocket.Conn
	// 发送 JSON 消息
	err = conn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		return err
	}
	return nil
}

func New() IMainFunc {
	return &Request{}
}
