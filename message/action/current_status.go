package action

type DownloadStatus struct {
	Status string `json:"status"`
}

func CurrentStatus(data interface{}) (interface{}, error) {
	di := DownloadStatus{
		Status: "not_downloading",
	}

	if currentDownloader.VideoUrl == "" {
		return di, nil
	}

	di.Status = "downloading"
	return di, nil
}
