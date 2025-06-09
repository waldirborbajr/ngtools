package hasnohup

import (
	"os"
)

func RemoveNoHup() error {
	return os.Remove("nohup.out")
}

func CreateNoHup() (*os.File, error) {
	return os.OpenFile("nohup.out", os.O_CREATE|os.O_WRONLY, 0600)
}
