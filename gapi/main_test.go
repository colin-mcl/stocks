package gapi

import (
	"os"
	"testing"
)

// Runs all *_test.go files found in the package
func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
