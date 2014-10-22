package main

import (
	"./segmentGo"
	"fmt"
)

//最大匹配+规则处理分词
//
//识别未登录词
//未登录词包括
//人名		--一般是 姓+名
//地名		--颍上 罗玲 等 如果有后缀县区还好识别 否则很难识别
//机构名	--管委会 工信处
//数字		--二十四 三十三 60等 数字后面还经常跟量词 如 口 米 等
//性别		--可以不处理 如 女干事 女同胞等

//未登录词的识别方法
//基于正向最大匹配
//遍历要处理的文本
//	获取当前要匹配的字
//	从字典树中取该字
//	如果取的到
//		缓存当前字
//		for(){
//			获取下个需要匹配的字
//			是上一个字的后缀? (可以组成词)
//				不是则break掉
//			加入缓冲区
//			继续往后匹配
//		}
//		如果缓存区为1
//			加入未登录词识别程序
//			程序返回处理后的词及下标
//			将词加入缓冲区
//		写缓冲区
//		按索引继续往后匹配

//	如果取不到--一般是符号或者是其他语言
//	现在直接pass

func main() {
	dic := new(segmentGo.DicTree)
	dic.LoadDic("./segmentGo/dic/main.dic")
	//女干事每月经过机房时都会检查二十四口交换机
	// result:=
	 result := segmentGo.FindUndef(dic, "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作")
	// result := segmentGo.FindUndef(dic, "最近晨晨怎么不聊天了")
	// result := segmentGo.FindUndef(dic, "美国会打伊拉克么")
	// result := segmentGo.FindUndef(dic, "小")
	// result := segmentGo.FindUndef(dic, "小明")
	// result := segmentGo.FindUndef(dic, "小明明")
	// result := segmentGo.FindUndef(dic, "小明明显")
	// result := segmentGo.FindUndef(dic, "小明明显得")
	// result := segmentGo.FindUndef(dic, "小明明显得了")
	// result := segmentGo.FindUndef(dic, "小明明显得了肥胖症")
	//result := segmentGo.FindUndef(dic, "【椒盐排骨】1.排骨切段，青红辣椒，洋葱切末。用葱姜，盐，糖，酱油，料酒把排骨腌制一小时;2.鸡蛋和淀粉搅成糊状。排骨裹上蛋糊放入五成热的油中炸酥，捞起滤干油;3.把辣椒，洋葱末加适量的椒盐略炒，放入排骨翻炒均匀。http://ww2.sinaimg.cn/bmiddle/88ffde5fjw1ecuyf0wt3bj20c909zgmh.jpg")
	for result != nil {
		fmt.Print(string(result.Word)+"  ")
		result = result.Next
	}

}
