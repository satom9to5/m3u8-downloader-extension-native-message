package main

import (
	"bufio"
	"os"
	// external packages
	"github.com/satom9to5/webext/nativemessaging"
	// this package
	"m3u8-downloader-extension-native-message/message"
	"m3u8-downloader-extension-native-message/process"
)

func main() {
	process.ParentWatch()

	for {
		req := message.NewRequestMessage()
		res := message.NewResponseMessage()
		var err error

		r := bufio.NewReader(os.Stdin)

		if err = nativemessaging.Receive(req, r); err == nil {
			res.Type = req.Type
			if data, err := req.Run(); err == nil {
				res.Data = data
			} else {
				res.Error = err.Error()
			}
		} else {
			res.Error = err.Error()
		}

		nativemessaging.Send(res, os.Stdout)
	}
}
