package events

type MetaEventType string

const (
	LIFECYCLE MetaEventType = "lifecycle"
	HEARTBEAT MetaEventType = "heartbeat"
)

type ICommonMeta interface {
	GetMetaPostType() PostType
	GetMetaEventType() MetaEventType
}

type ILifeCycle interface {
	ICommonMeta
	GetMetaSubType() string
}

type IHeartbeat interface {
	ICommonMeta
	GetHeartbeatStatus() any
	GetHeartbeatInterval() int64
}

type MetaEventStruct struct {
	Time          int64         `json:"time"`
	SelfID        int64         `json:"self_id"`
	PostType      PostType      `json:"post_type"`
	MetaEventType MetaEventType `json:"meta_event_type"`
	SubType       string        `json:"sub_type"`
	Status        any           `json:"status"`   // 状态信息
	Interval      int64         `json:"interval"` // 到下次心跳的间隔，单位毫秒
}

func (e *Event) GetMetaPostType() PostType {
	return e.MetaEventStruct.PostType
}

func (e *Event) GetMetaEventType() MetaEventType {
	return e.MetaEventStruct.MetaEventType
}

func (e *Event) GetHeartbeatStatus() any {
	return e.MetaEventStruct.Status

}

func (e *Event) GetHeartbeatInterval() int64 {
	return e.MetaEventStruct.Interval
}

func (e *Event) GetMetaSubType() string {
	return e.MetaEventStruct.SubType
}
