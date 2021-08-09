package action

import (
	"errors"

	"github.com/satom9to5/m3u8-download/m3u8"
)

type downloader struct {
	VideoUrl string `mapstructure:"video_url"`
	AudioUrl string `mapstructure:"audio_url"`
	Filename string `mapstructure:"filename"`
	si       *m3u8.StreamInfo
}

var (
	currentDownloader = &downloader{}
	dummyDownloader   = &downloader{}
)

func (d *downloader) Start() error {
	d.si = &m3u8.StreamInfo{
		VideoUri: d.VideoUrl,
		AudioUri: d.AudioUrl,
	}

	return d.si.Download(d.Filename)
}

func (d *downloader) Stop() error {
	if d.si == nil {
		return errors.New("not downloading")
	}

	return d.si.Cancel()
}
