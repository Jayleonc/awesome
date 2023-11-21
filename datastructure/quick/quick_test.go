package quick

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name       string
		list       []int
		answerList []int
	}{
		{
			name:       "第一个排序",
			list:       []int{50, 30, 20, 60, 40, 10},
			answerList: []int{10, 20, 30, 40, 50, 60},
		},
		{
			name:       "第二个排序",
			list:       []int{300, 500, 200, 100, 400, 600},
			answerList: []int{100, 200, 300, 400, 500, 600},
		},
		{
			name:       "第三个排序",
			list:       []int{320, 520, 220, 120, 420, 620},
			answerList: []int{120, 220, 320, 420, 520, 620},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			ints := Sort(v.list, 0, len(v.list)-1)
			assert.Equal(t, ints, v.answerList)
		})
	}
}
