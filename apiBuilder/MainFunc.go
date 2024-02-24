package apiBuilder

type IMainFunc interface {
	SendGroupMsg() ISendGroupMsg
	SendPrivateMsg() ISendPrivateMsg
}
