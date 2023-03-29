package stringutil

import "testing"

func TestToInt(t *testing.T) {
	i := ToInt("123")
	if i == -1 {
		t.Errorf("转换失败！赋默认值-1")
	}
}
