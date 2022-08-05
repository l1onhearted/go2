package main

import "fmt"

func firstUniqChar(s string) int {
	d:=make(map[byte]int)
	for i:=range s{
		d[s[i]]++
	}
	var t byte
	t='1'
	fmt.Println(d)
	for k,v:=range d{
		if v!=0{
			t=k
			break
		}
	}
	if t=='1'{
		return -1
	}
	for i,v:= range []byte(s){
		if d[v]==1{
			return i
		}
	}
	return -1
}  

func main(){
	s:= "leetcode"
	fmt.Println(firstUniqChar(s))
}