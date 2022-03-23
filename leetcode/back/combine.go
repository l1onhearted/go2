//77组合

package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	var backtrack func(n, k, start int, track []int)
	backtrack = func(n, k, start int, track []int) {
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, track)
		}
		if len(track)+n-start+1 < k { //剪枝
			return
		}
		for i := start; i < n; i++ {
			track = append(track, i)
			backtrack(n, k, i+1, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(n, k, 1, []int{})
	return res
}

//216.组合总和
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	
}

// func main() {
// 	t := []int{1, 2, 3, 4, 5}
// 	t = t[0:4]
// 	fmt.Println(t)
// }
