package action

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

func Download(data interface{}) (interface{}, error) {
	if currentDownloader != nil && currentDownloader.VideoUrl != "" {
		return nil, errors.New("exist current download task")
	}

	d := &downloader{}

	err := mapstructure.Decode(data, d)

	if err != nil {
		return nil, err
	}
	if d.VideoUrl == "" || d.Filename == "" {
		return nil, errors.New("Url/Filename is empty!")
	}

	currentDownloader = d

	if err := d.Start(); err != nil {
		return nil, errors.New("cannot start download task")
	}

	currentDownloader = nil

	return d.Filename, nil
}
