segmentGo
=========

Go 语言版的分词器;基于字典树和最大匹配算法;略微的加了点消歧;

V0.1.2
支持数值识别,简单的未登录词识别。
示例

	package main

	import (
		"./segmentGo"
		"fmt"
	)

	func main() {
		dic := new(segmentGo.DicTree)
		dic.LoadDic("./segmentGo/dic/main.dic")
		//女干事每月经过机房时都会检查二十四口交换机
		// result:=
		result := segmentGo.FindUndef(dic, "工信处女干事每月经过下属科室都要亲口交代二十四口交换机等技术性器件的安装工作")
		for result != nil {
			fmt.Println(string(result.Word))
			result = result.Next
		}

	}


V0.1
目前只支持正向最大匹配，最大匹配算法分词的准确性不高 但是速度较快 约为160万字/秒

示例

	package main

	import (
		"fmt"
		"github.com/violetgo/segmentGo"
		"time"
	)

	func main() {


		dic := new(segmentGo.DicTree)
		dic.LoadDic("segmentGo/dic/main.dic")
		nano := time.Now()
		result := segmentGo.Sgement(dic, "掌握一门语言的标准就是 你能不能用它表达你的思想.掌握之后，就剩下实实在在的问题：你有没有思想要表达.")
		nano2 := time.Now()
		fmt.Println(nano2.Sub(nano).Seconds())
		fmt.Println(result)
		for result != nil {
		 	fmt.Println(string(result.Word))
		 	result = result.Next
		}

	}
