package main

import (
	"fmt"
	"strings"
)

func main() {
	//直接定义字符串
	str := "jinlInihao"

	//是否包含字符
	contains := strings.Contains("j", str)
	fmt.Println("contains", contains)

	containsAny := strings.ContainsAny("jl", str)
	fmt.Println("containsAny", containsAny)

	//统计字符数量
	count := strings.Count("a", str)
	fmt.Println("count", count)

	//匹配前缀
	prefix := strings.HasPrefix("h", str)
	fmt.Println("prefix", prefix)

	//匹配后缀
	suffix := strings.HasSuffix("a", str)
	fmt.Println("suffix", suffix)

	//寻找下标
	/*strings.Index()
	strings.IndexAny()
	strings.LastIndex()*/

	//字符拼接
	str1 := []string{"a", "b", "c"}
	join := strings.Join(str1, "=") // a=b=c
	fmt.Println("join", join)

	//自己拷贝自己
	repeat := strings.Repeat("i", 7) // iiiiiii
	fmt.Println("repeat", repeat)

	//替换敏感字符
	replace := strings.Replace(str, "j", "*", 2)
	fmt.Println("replace", replace) // *inlInihao

	//用户名大小写转换
	fmt.Println("大小写转换", str)
	upper := strings.ToUpper(str)
	fmt.Println("upper", upper)
	fmt.Println("str", str)
	lower := strings.ToLower(str)
	fmt.Println("lower", lower)
	fmt.Println("str", str)

}
