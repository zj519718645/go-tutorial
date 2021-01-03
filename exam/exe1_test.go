// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
package practise

import (
	"fmt"
	"strings"
	"testing"
)

func longestPalindrome(src string) string {
	lenS := len(src)
	if lenS == 0 {
		return ""
	} else if lenS == 1 {
		return src
	} else if lenS == 2 {
		if src[0] == src[1] {
			return src
		} else {
			return string(src[0])
		}
	}
	//preprocess
	s := make([]rune, 1)
	s[0] = '#'
	for _, ss := range src {
		s = append(s, ss)
		s = append(s, '#')
	}
	lenS = len(s)
	middleIndex1, middleIndex2 := (lenS-1)/2, lenS/2
	maxLen := 1
	centreIndex := middleIndex1
	stepsLeft, stepsRight := 1, 1
	for leftMaxLen := 1; middleIndex1 > 0 && stepsLeft <= middleIndex1; stepsLeft++ {
		if s[middleIndex1-stepsLeft] == s[middleIndex1+stepsLeft] {
			leftMaxLen = leftMaxLen + 2
		} else {
			leftMaxLen = 1
			middleIndex1 = middleIndex1 - 1
			stepsLeft = 0
		}
		if leftMaxLen > maxLen {
			centreIndex = middleIndex1
			maxLen = leftMaxLen
		}
		if middleIndex1 < maxLen/2 {
			break
		}
	}
	for rightMaxLen := 1; middleIndex2 < lenS-1 && stepsRight <= lenS-middleIndex2-1; stepsRight++ {
		if s[middleIndex2-stepsRight] == s[middleIndex2+stepsRight] {
			rightMaxLen = rightMaxLen + 2
		} else {
			rightMaxLen = 1
			middleIndex2 = middleIndex2 + 1
			stepsRight = 0
		}
		if rightMaxLen > maxLen {
			maxLen = rightMaxLen
			centreIndex = middleIndex2
		}
		if lenS-middleIndex2 < maxLen/2 {
			break
		}
	}
	fmt.Println(maxLen)
	fmt.Println(centreIndex)
	return strings.Replace(string(s[centreIndex-(maxLen/2):centreIndex+((maxLen+1)/2)]), "#", "", -1)
}

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("abbd"))
}
