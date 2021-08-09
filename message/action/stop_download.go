package action

import (
	"errors"
)

func StopDownload(data interface{}) (interface{}, error) {
	if currentDownloader == nil || currentDownloader.VideoUrl == "" {
		return nil, errors.New("not exist current download task")
	}

	if err := currentDownloader.Stop(); err != nil {
		return nil, errors.New("cannot stop current download task")
	}

	currentDownloader = nil

	return nil, nil
}
