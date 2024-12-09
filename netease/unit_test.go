package netease

import "testing"

func TestGetTop200Music(t *testing.T) {
	url := "https://music.163.com/discover/toplist?id=3778678"
	ids := GetTop200Music(url)
	t.Logf("ids: %v\n", ids)
}
