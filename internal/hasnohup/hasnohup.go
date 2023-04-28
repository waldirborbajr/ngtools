package hasnohup

import "os"

func HasNoHup() {
	_, err := os.Stat("nohup")

	if !os.IsNotExist(err) {
		os.Remove("nohup")
	}
}
