package ch3

import "testing"

func TestMap(t *testing.T) {
	dict := make(map[string]int)
	t.Log(dict)
	telnum := make(map[int]string)
	telnum[110] = "警察局"
	t.Log(telnum)
	telnum[120] = "急救中心"

	value, exists := telnum[112]
	if exists {
		t.Log(value)
	} else {
		t.Log("it's not exist")

	}
	if value != "" {
		t.Log(value)
	}
	delete(telnum, 120)
	t.Log(telnum)
	for key, val := range telnum {
		t.Log(key, val)
	}

}
