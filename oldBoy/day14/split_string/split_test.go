package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用力失败了
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
