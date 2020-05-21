package split_string

// 测试用例文件命名：xx_test.go
// 函数必须是Testxx开头命名

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcdbef", "b")
	want := []string{"a", "cd", "ef"}
	if !reflect.DeepEqual(got, want) {
		// 测试用例失败
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("abcbef", "b")
	want := []string{"a", "c", "ef"}
	if !reflect.DeepEqual(got, want) {
		// 测试用例失败
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}
