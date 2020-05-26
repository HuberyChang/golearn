package split_string

// 测试用例文件命名：xx_test.go
// 函数必须是Testxx开头命名

import (
	"reflect"
	"testing"
)

func Test1Split(t *testing.T) {
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

func Test3Split(t *testing.T) {
	got := Split("abcdef", "cd")
	want := []string{"ab", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}
// 	testGroup := []testCase{
// 		testCase{"abcdbef", "b", []string{"a", "cd", "ef"}},
// 		testCase{"abcdef", "b", []string{"a", "cdef"}},
// 		testCase{"abcdef", "cd", []string{"ab", "ef"}},
// 		testCase{"发哈发空", "发", []string{"哈", "空"}},
// 	}
// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	// testGroup := []testCase{
	// 	testCase{"abcdbef", "b", []string{"a", "cd", "ef"}},
	// 	testCase{"abcdef", "b", []string{"a", "cdef"}},
	// 	testCase{"abcdef", "cd", []string{"ab", "ef"}},
	// 	testCase{"发哈发空", "发", []string{"哈", "空"}},
	// }
	testGroup := map[string]testCase{
		"case1": testCase{"abcdbef", "b", []string{"a", "cd", "ef"}},
		"case2": testCase{"abcdef", "b", []string{"a", "cdef"}},
		"case3": testCase{"abcdef", "cd", []string{"ab", "ef"}},
		// "case4": testCase{"发哈发空", "发", []string{"哈", "空"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want:%v but got:%v\n", tc.want, got)
			}
		})
	}
}
