package models

const OK int = 200
const InnerError int = 500
const ParamError int = 400

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
