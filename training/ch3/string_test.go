package ch3

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str1 := "abcd"
	str2 := "efgh"
	t.Log(str1 + str2)

	str3 := fmt.Sprintf("%s%s", str1, str2)
	t.Log(str3)

	var str4 = "我是中国人"
	for index, val := range str4 {
		t.Logf("%d,%c\n", index, val)
	}

	t.Log(strings.Contains(str3, "abcd"))
	t.Log(strings.Contains(str3, "abcdg"))
	t.Log(strings.Index(str3, "abcd"))
}
