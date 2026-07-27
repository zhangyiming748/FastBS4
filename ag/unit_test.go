package ph

import (
	"strings"
	"testing"
)

// go test -v -run TestGetFromFile
func TestGetFromFile(t *testing.T) {
	name := "女王的性生活"
	name = strings.Join([]string{name,"txt"}, ".")
	GetFromFile(name)
}
  