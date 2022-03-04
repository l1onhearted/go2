//剑指58 左旋转字符串

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, n-1)      //反转前n个字符
	reverse(b, n, len(b)-1) //反转n到最后一个字符
	reverse(b, 0, len(b)-1) //反转整个字符
	return string(b)
}
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}