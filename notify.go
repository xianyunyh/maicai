package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Notifyer interface {
	Send() error
}

type TmuxNotify struct {
	media string
}

func (t TmuxNotify) Send() error {
	path := os.Getenv("PATH")
	cmdPath := fmt.Sprintf("%s/%s", path, "termux-media-player")
	homePath := os.Getenv("HOME")
	mediaPath := fmt.Sprintf("%s/%s", homePath, t.media)
	cmd := exec.Command(cmdPath, "play", mediaPath)
	return cmd.Run()
}
