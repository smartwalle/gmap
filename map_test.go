package gmap

import (
	"testing"
	"os"
)

var client *GoogleMap

func TestMain(m *testing.M) {
	client = New("", "")
	exitCode := m.Run()
	os.Exit(exitCode)
}
