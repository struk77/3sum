package threesum

import (
	"sort"
)

//ThreeSum is a fuction
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var resMap = map[[3]int]int{}
	idx := 0
	for i, num := range nums {
		value := 0 - num
		cuttedNums := nums[i+1:]
		if len(cuttedNums) < 2 {
			break
		}
		res := twoSum(cuttedNums, value)
		if len(res) == 0 {
			continue
		}
		for _, r := range res {
			_, exist := resMap[r]
			if !exist {
				resMap[r] = idx
				idx++
			}
		}
	}
	if len(resMap) < 1 {
		return [][]int{}
	}
	var results = make([][]int, len(resMap))
	for k := range resMap {
		i := resMap[k]
		z := make([]int, 3)
		z[0] = k[0]
		z[1] = k[1]
		z[2] = k[2]
		results[len(resMap)-i-1] = z
	}

	return results
}

func twoSum(nums []int, target int) (result [][3]int) {
	var numMap = map[int]int{}
	for i, num := range nums {
		j := target - num
		idx, exist := numMap[j]
		if exist {
			r := [3]int{0 - target, nums[idx], num}
			result = append(result, r)
		}
		numMap[num] = i
	}
	return
}
