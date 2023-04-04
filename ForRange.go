package main

import (
	"fmt"
)

/*func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))//空
	fmt.Println("------------------")
	fmt.Println(os.Args[1:])//[]
}*/

func main() {
	s, sep := "", ""
	a := []string{"1", "2", "3"}
	for re, args := range a[0:] {
		s += sep + args
		sep = "  "
		/*re-- 0
		re-- 1
		re-- 2    re:是索引下标
		*/
		fmt.Println("re--", re)
	}

	fmt.Println("***---", s) //***--- 1  2  3

}
