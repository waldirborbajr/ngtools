package hasnohup

import (
	"os"
	"testing"
)

func TestRemoveNoHup(t *testing.T) {
	// Create a dummy nohup.out file
	f, err := os.Create("nohup.out")
	if err != nil {
		t.Fatalf("Failed to create nohup.out: %v", err)
	}
	f.Close()

	// Remove the file using RemoveNoHup
	err = RemoveNoHup()
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("RemoveNoHup returned error: %v", err)
	}

	// Check that the file no longer exists
	if _, err := os.Stat("nohup.out"); !os.IsNotExist(err) {
		t.Errorf("nohup.out should not exist after RemoveNoHup")
	}
}

func TestCreateNoHup(t *testing.T) {
	// Ensure the file does not exist before test
	_ = os.Remove("nohup.out")

	f, err := CreateNoHup()
	if err != nil {
		t.Fatalf("CreateNoHup returned error: %v", err)
	}
	defer f.Close()

	// Check file exists
	info, err := os.Stat("nohup.out")
	if err != nil {
		t.Fatalf("nohup.out does not exist after CreateNoHup: %v", err)
	}

	// Check permissions are 0600
	if info.Mode().Perm() != 0600 {
		t.Errorf("nohup.out permissions = %v; want 0600", info.Mode().Perm())
	}

	// Clean up
	_ = os.Remove("nohup.out")
}
