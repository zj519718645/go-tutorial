//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
package practise

import "testing"

func convert(s string, numRows int) string {
	var srcLen = len(s)
	var target []byte
	if srcLen <= numRows {
		return s

	}
	if numRows == 1 {
		return s
	}
	for ele := 0; row < srcLen; row++ {
		for column := 0; column*numRows < srcLen; column++ {
			target = append(target, s[numRows*column+column])
		}
	}
	return string(target)
}

func TestConvert(t *testing.T) {

}
