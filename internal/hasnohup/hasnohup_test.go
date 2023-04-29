package hasnohup_test

import (
	"localhost/ngtools/internal/hasnohup"
	"os"
	"testing"
)

func TestHasNoHup(t *testing.T) {
	// create a dummy "nohup" file
	file, err := os.Create("nohup")
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// call the HasNoHup function
	hasnohup.HasNoHup()

	// check if the "nohup" file still exists
	if _, err := os.Stat("nohup"); err == nil {
		t.Errorf("HasNoHup did not remove the 'nohup' file")
	}
}
