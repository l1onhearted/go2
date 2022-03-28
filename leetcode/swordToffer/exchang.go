func exchange(nums []int) []int {
	// for i, j := 0, len(nums)-1; i < j; {

	// }
	res := []int{}
	for _, i := range nums {
		if i%2 == 0 {
			res = append(res, i)
		} else {
			t := []int{i}
			res = append(t, res...)
		}
	}
	return res
}