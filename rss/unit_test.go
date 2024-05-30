package rss

import "testing"

// go test -v -run TestGetFromFile
func TestGetFromFile(t *testing.T) {
	GetFromFile()
}

// go test -v -run TestGetFromFile2
func TestGetFromFile2(t *testing.T) {
	GetFromFile2()
}

// go test -v -run TestDownload
func TestDownload(t *testing.T) {
	download()
}
