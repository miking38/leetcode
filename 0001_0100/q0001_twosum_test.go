package q0001_twosum_test 

import "testing"

func twoSum(nums []int, target int) []int {
    index_map := make(map[int]int)
    for i, _ := range(nums) {
        if j, exist := index_map[target - nums[i]]; exist {
            return []int{j, i}
        } else {
            index_map[nums[i]] = i
        }
    }
    
    return []int{-1,-1}
}

func TestTwoSum(t *testing.T) {
	a1 := twoSum([]int{1,2,4}, 6)
	if a1[0] != 1 || a1[1] != 2 {
		t.Error("expected :  [1,2]")
	}
}
