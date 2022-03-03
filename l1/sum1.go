package l1

func TwoSum(nums []int, target int) []int {
	sum := make([]int, 2)
	for i, n := range nums {
		t := target - n
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == t {
				sum[0] = i
				sum[1] = j
				return sum
			}
		}
	}
	return nil
}

func TwoSum1(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i, n := range nums {
		n1, f := m1[target-n]
		if f {
			return []int{n1, i}
		} else {
			m1[n] = i
		}
	}
	return nil
}
