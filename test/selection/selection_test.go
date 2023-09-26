package selection

import (
	"reflect"
	"testing"
)

// 测试函数 前缀为 Test，用来测试程序的逻辑行为是否正确

// 基准函数 前缀为 Benchmark，用来测试函数的性能

// 示例函数 前缀为 Example，用来为文档提供示例

type tests struct {
	arr  []int
	want []int
}

var test = map[string]tests{
	"1": {
		arr:  []int{97, 12, 54, 2, 100},
		want: []int{2, 12, 54, 97, 100},
	},

	"2": {
		arr:  []int{1, 12, 54, 2, 100},
		want: []int{1, 2, 12, 54, 100},
	},
	"3": {
		arr:  []int{90, 2, 54, 2, 199},
		want: []int{2, 2, 54, 90, 199},
	},
	"4": {
		arr:  []int{2},
		want: []int{2},
	},
}

func TestSoreSelection(t *testing.T) {
	for name, item := range test {
		t.Run(name, func(t *testing.T) {
			sort := SorcSelection(item.arr)
			equal := reflect.DeepEqual(sort, item.want)
			if !equal {
				t.Error("排序失败")
			}
		})
	}
}

func BenchmarkSorcSelection(b *testing.B) {

	arr := []int{97, 12, 54, 2, 100}
	for i := 0; i < b.N; i++ {
		SorcSelection(arr)
	}
}
