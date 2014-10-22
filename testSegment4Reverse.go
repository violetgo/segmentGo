package main

import (
	"fmt"
	"./segmentGo"

)

func main() {

	// fmt.Println(io.ReadFile("165.txt"))

	dic := new(segmentGo.DicTree)
	dic.LoadDic4Reverse("../segmentGo/dic/main.dic")
	fmt.Print("小便当汤 大便当饭")
	result := //segmentGo.Sgement4Reverse(dic,"工信处女干事每月经过机房时都会检查二十四口交换机")
	// segmentGo.Sgement4Reverse(dic, "掌握一门语言的标准就是 你能不能用它表达你的思想.掌握之后，就剩下实实在在的问题：你有没有思想要表达.")
	segmentGo.Sgement4Reverse(dic,"小便当汤 大便当饭")
	

	// fmt.Println(result)
	for result != nil {
		fmt.Println(string(result.Word))
		result = result.Next
	}

}
