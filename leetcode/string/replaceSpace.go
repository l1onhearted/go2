//剑指05 替换空格

package main

func replaceSpace(s string) string {
	str := []byte(s)
	count := 0
	for _, i := range str { //计算空格数量
		if i == ' ' {
			count++
		}
	}
	//扩展原切片
	resize := count * 2
	tmp := make([]byte, resize)
	str = append(str, tmp...) //...表示添加切片
	i := len(s) - 1
	j := len(str) - 1 //双指针，一个指向原始长度，一个指向扩容后的长度
	for i >= 0 {
		if str[i] != ' ' {
			str[j] = str[i]
			i--
			j--
		} else {
			str[j] = '0'
			str[j-1] = '2'
			str[j-2] = '%'
			i--
			j -= 3
		}
	}
	return string(str)
}
