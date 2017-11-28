package gmap

import (
	"testing"
	"os"
)

var client *GoogleMap

func TestMain(m *testing.M) {
	client = New("AIzaSyDYGqqQFZrUqWYiR_ez9gIx8t7CbTpjsQE", "")
	exitCode := m.Run()
	os.Exit(exitCode)
}
