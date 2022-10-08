package leetcode

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, exists := m[target-v]; exists {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}
