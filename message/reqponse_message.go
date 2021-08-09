package message

import "fmt"

type ResponseMessage struct {
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func NewResponseMessage() *ResponseMessage {
	return &ResponseMessage{}
}

func (rm ResponseMessage) String() string {
	return fmt.Sprintf("Type: \"%s\", Data:\"%s\", Error: \"%s\"", rm.Type, rm.Data, rm.Error)
}
