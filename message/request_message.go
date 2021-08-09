package message

import (
	"fmt"
	// this package
	"m3u8-downloader-extension-native-message/message/action"
)

const (
	SetConfig     = "setConfig"
	Download      = "download"
	StopDownload  = "stopDownload"
	CurrentStatus = "currentStatus"
)

type RequestMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewRequestMessage() *RequestMessage {
	return &RequestMessage{}
}

func (rm RequestMessage) String() string {
	return fmt.Sprintf("Type: %s\tData: [%s]", rm.Type, rm.Data)
}

func (rm *RequestMessage) Run() (interface{}, error) {
	switch rm.Type {
	case SetConfig:
		return action.SetConfig(rm.Data)
	case Download:
		return action.Download(rm.Data)
	case StopDownload:
		return action.StopDownload(rm.Data)
	case CurrentStatus:
		return action.CurrentStatus(rm.Data)
	default:
		return nil, nil
	}
}
