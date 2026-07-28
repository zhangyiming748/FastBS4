package ph

import (
	"strings"
	"testing"
)

// go test -v -run TestGetFromFile
func TestGetFromFile(t *testing.T) {
	name := "步非烟"
	name = strings.Join([]string{name,"txt"}, ".")
	GetFromFile(name)
}
  