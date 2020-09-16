package threesum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := [][]int{{-1, 0, 1, 2, -1, -4}, {-2, 0, 1, 1, 2}}
	expected := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}, {{-2, 0, 2}, {-2, 1, 1}}}
	for j, num := range nums {
		results := ThreeSum(num)
		if len(results) != len(expected[j]) {
			t.Errorf("Expected %v, got %v", expected[j], results)
			return
		}
		for i, res := range results {
			//fmt.Printf("Got %v : expected %v\n", res, expected[j][i])
			for k, v := range res {
				if v != expected[j][i][k] {
					t.Errorf("Expected %v, got %v", expected[j][i], res)
				}
			}
		}
	}

}

func BenchmarkThreeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, 0, 1, 1, 2}
	for n := 0; n < b.N; n++ {
		ThreeSum(nums)
	}
}
