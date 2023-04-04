package main

import (
	"fmt"
	"time"
)

func main() {
	//time1() //2022-11-28 16:04:43.8440248 +0800 CST m=+0.001448301
	//go语言的格式化模板是固定的
	now := time.Now()
	unix := time.Now().Unix()
	fmt.Println("时间转时间戳：", unix) // 1669625440  多少秒
	t := time.Unix(unix, 0)            //参数0
	fmt.Println("时间戳转时间：", t)    //时间戳转时间： 2022-11-28 16:50:40 +0800 CST

	//2022-11-28 16:19:43:年-月-日 时：分：秒 Println调用这个方法
	fmt.Println(now.Format("2006-01-02 15:04:05"))    //2006-01-02 15:04:05  go语言的诞生时间
	fmt.Println(now.Format("2006-01-02 03:04:05 PM")) //2022-11-28 04:26:32 PM
	fmt.Println(now.Format("2006/01/02 15:04"))       //2022/11/28 16:27
	//fmt.Println(now.Format("yyyy-01-02 15:04:05")) //写字母没有用
	location, _ := time.LoadLocation("Asia/Shanghai")
	inLocation, _ := time.ParseInLocation("2006/01/02 15:04", "2022/11/28 16:40", location)
	fmt.Println("时区：", inLocation) //时区： 2022-11-28 16:40:00 +0800 CST（CST 代表的是中国标准时间）

}
func time1() {
	//获取当前的时间
	now := time.Now()
	/*fmt.Println(now.Year())
	fmt.Println(now.Date()) //2022 November 28
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())*/
	now.Minute()
	now.Second()
	//2022-11-28 16:11:29
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//2022-11-28 16:11:292022-11-28 16:11:29.6969697 +0800 CST m=+0.001491301
	fmt.Println(now)
}
