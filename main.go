package main

import (
	"./segmentGo"
	"fmt"
	"github.com/violetgo/tools/io"
	"time"
)

func main() {

	// fmt.Println(io.ReadFile("165.txt"))

	dic := new(segmentGo.DicTree)
	dic.LoadDic("./segmentGo/dic/main.dic")
	str := io.ReadFile("165.txt")
	nano := time.Now()

	// fmt.Println(len(temp.children))
	// fmt.Println(temp.getWord("语言"))
	// fmt.Println(temp.getWord("中华人民共和国"))
	// fmt.Println(temp.getWord("中华人民共和国万岁"))
	// getWordByRune(temp.children, []rune("语")[0])
	// temp.sgement("中华人民共和国万岁")
	// temp.sgement("中国人民站起来了")
	// result := segmentGo.Sgement(dic, "掌握一门语言的标准就是 你能不能用它表达你的思想.掌握之后，就剩下实实在在的问题：你有没有思想要表达.")
	// temp.sgement("小便当汤 大便当饭")
	// temp.sgement("是不是谣言")
	segmentGo.Sgement(dic, str)
	// nano2 := time.Now().Nanosecond()
	// fmt.Println(nano2, nano)
	// temp.sgement("工信处女干事每月经过机房时都会检查二十四口交换机")
	nano2 := time.Now()

	fmt.Println(nano2.Sub(nano).Seconds())
	// fmt.Println(result)
	// for result != nil {
	// 	fmt.Println(string(result.Word))
	// 	result = result.Next
	// }

}
