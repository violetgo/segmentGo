package main

import (
	"./segmentGo"
	"fmt"
)

func main() {

	// fmt.Println(io.ReadFile("165.txt"))

	dic := new(segmentGo.DicTree)
	dic.LoadDic("./segmentGo/dic/main.dic")
	// str := io.ReadFile("165.txt")
	// nano := time.Now()

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
	// segmentGo.Sgement(dic, str)
	// nano2 := time.Now().Nanosecond()
	// fmt.Println(nano2, nano)
	result := segmentGo.Sgement(dic, "【椒盐排骨】1.排骨切段，青红辣椒，洋葱切末。用葱姜，盐，糖，酱油，料酒把排骨腌制一小时;2.鸡蛋和淀粉搅成糊状。排骨裹上蛋糊放入五成热的油中炸酥，捞起滤干油;3.把辣椒，洋葱末加适量的椒盐略炒，放入排骨翻炒均匀。http://ww2.sinaimg.cn/bmiddle/88ffde5fjw1ecuyf0wt3bj20c909zgmh.jpg")
	// nano2 := time.Now()

	// fmt.Println(nano2.Sub(nano).Seconds())
	// fmt.Println(result)
	for result != nil {
		fmt.Print(string(result.Word)+"   ")
		result = result.Next
	}

}
