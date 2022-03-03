//242.有效的字母异位词 https://leetcode-cn.com/problems/valid-anagram/
package main

func isAnagram(s string, t string) bool {
	var record = [26]int{}
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		record[t[i]-'a']--
	}
	for i := 0; i < len(record); i++ {
		if record[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool { //使用哈希表
	table := make(map[byte]int)
	for _, i := range s {
		if v, ok := table[s[i]]; v > 0 && ok {
			table[s[i]]++
		} else {
			table[s[i]] = 0
		}
	}
	for _, i := range t {
		if v, ok := table[t[i]]; v > 0 && ok {
			table[t[i]]--
		} else {
			return false
		}
	}
	return true
}
func main() {
	s := "zhangyuankun"
	t := "yuankunzhangz"
	// fmt.Println(isAnagram(s, t))
	isAnagram1(s, t)

}
