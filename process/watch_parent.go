package process

import (
	"os"
	"time"

	"github.com/mitchellh/go-ps"
)

func ParentWatch() {
	go checkParentProcessLive()
}

func checkParentProcessLive() {
	for {
		if proc, err := ps.FindProcess(os.Getppid()); proc == nil || err != nil {
			os.Exit(0)
		}

		time.Sleep(2 * time.Second)
	}
}
