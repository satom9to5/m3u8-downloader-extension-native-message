package action

import (
	"github.com/mitchellh/mapstructure"
	"github.com/satom9to5/m3u8-download/m3u8"
)

type config struct {
	DownloadDir string `mapstructure:"download_dir"`
	WorkDir     string `mapstructure:"work_dir"`
	FFMpegPath  string `mapstructure:"ffmpeg_path"`
	LogTarget   string `mapstructure:"log_target"`
}

func SetConfig(data interface{}) (interface{}, error) {
	c := &config{}

	err := mapstructure.Decode(data, c)

	if err != nil {
		return nil, err
	}

	if c.DownloadDir != "" {
		m3u8.DownloadDir = c.DownloadDir
	}

	if c.WorkDir != "" {
		m3u8.WorkDir = c.WorkDir
	}

	if c.FFMpegPath != "" {
		m3u8.FFMpegPath = c.FFMpegPath
	}

	if c.LogTarget != "" {
		m3u8.SetLogTarget(c.LogTarget)
	}

	return nil, nil
}
