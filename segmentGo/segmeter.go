package segmentGo

import (
	"fmt"
)

type Token struct {
	Word []rune
	Next *Token
}

//逆向词典
func Sgement4Reverse(dic *DicTree, word string) *Token {

	curNode := dic
	wordRune := []rune(word)
	conster := new(Token)
	var last *Token

	temp := make([]rune, 0, 2)
	isback := 0
	for index := len(wordRune) - 1; index >= 0; {

		cur := wordRune[index]
		//fmt.Print(string(cur))
		curNode = getWordByRune(curNode.children, cur)
		//fmt.Print(string(cur))
		if curNode == nil || !curNode.isToekn {

			// fmt.Println("--end")
			curNode = dic

			last = conster
			last.Word = reverseStr(temp)
			conster = new(Token)
			conster.Next = last

			temp = make([]rune, 0, 2)

			if isback == 1 {
				index--
				isback = 0
				continue
			}

			isback = 1
			continue
		}
		// fmt.Println("--add")
		isback = 0
		temp = append(temp, cur)
		// step++
		index--
	}

	last = conster
	last.Word = reverseStr(temp)
	conster = new(Token)
	conster.Next = last

	return last
}

func parse(cur rune, temp []rune) {

	if len(temp) == 1 {
		fmt.Println(string(cur))
	}

	temp = make([]rune, 0, 2)

}

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
func FindUndef(dic *DicTree, word string) *Token {

	conster := new(Token)
	var result *Token
	result = conster

	curNode := dic
	wordRune := []rune(word)
	tempBuff := make([]rune, 0, 2)
	for index := 0; index < len(wordRune); {
		isWord := false
		cur := wordRune[index]
		// fmt.Println("get", string(cur))
		curNode = getWordByRune(curNode.children, cur)
		rollback := 0
		if curNode == nil || !curNode.isToekn {
			//this is not a word ,so  we just skip it
			//and go to dic root

			index++
			curNode = dic
		} else { //取到
			tempNode := curNode
			tempBuff = append(tempBuff, cur)
			tempindex := index + 1
			for ; tempindex < len(wordRune); tempindex++ {
				tempWord := wordRune[tempindex]
				tempNode = getWordByRune(tempNode.children, tempWord)
				if tempNode == nil || !tempNode.isToekn {
					break
				} else {
					tempBuff = append(tempBuff, tempWord)
					if !isWord {
						isWord = tempNode.isWord
					}
					if isWord && !tempNode.isWord {
						rollback++
					}
				}
			}

			if len(tempBuff) == 1 {

				tempBuff = guessWord(wordRune[tempindex-1:])
				// index = (index - 1 + len(tempBuff))
				// fmt.Println("parse", string(cur))
				isWord = true
			}

			if !isWord {
				// fmt.Println("not word", string(tempBuff))
				conster.Word = tempBuff[0:1]
				conster.Next = new(Token)
				conster = conster.Next
				index = index + 1
				tempBuff = make([]rune, 0, 2)
				curNode = dic
			} else {
				//try get math
				// fmt.Println("...befo", string(tempBuff[0]), rollback)
				tempBuff = guessMath(tempBuff, wordRune[tempindex:])
				// fmt.Println("...", string(tempBuff[0]), len(tempBuff), index)
				index = (index + len(tempBuff))
				index = index - rollback
				//add to list and reset the curNode to dic root
				conster.Word = tempBuff[0 : len(tempBuff)-rollback]
				conster.Next = new(Token)
				conster = conster.Next
				tempBuff = make([]rune, 0, 2)
				curNode = dic
			}

		}
	}

	return result
}

//人名 ,地名 机构名识别
func guessWord(wordlist []rune) []rune {
	if len(wordlist) == 1 {
		return wordlist
	}
	tempBuff := make([]rune, 0, 2)
	for i := 0; i < len(wordlist); i++ {
		// fmt.Println("get", string(wordlist[i]))
		tempBuff = append(tempBuff, wordlist[i])
		if string(wordlist[i]) == "处" {
			return tempBuff
		}
	}

	return wordlist[0:1]
}

//数值识别 通常是数值+量词
func guessMath(word []rune, wordlist []rune) []rune {
	// fmt.Println("test", string(word[0]))
	for _, cur := range word {
		switch string(cur) {
		case "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "零":
			continue
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			continue
		default:
			return word
		}
	}

	// tempBuff := make([]rune, 0, 2)
	for _, cur := range wordlist {
		switch string(cur) {
		case "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "零":
			word = append(word, cur)
			continue
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			word = append(word, cur)
			continue
		case "口":
			word = append(word, cur)
			return word
		default:
			return word
		}
	}
	return word
}

//正向词典
func Sgement(dic *DicTree, word string) *Token {

	curNode := dic
	wordRune := []rune(word)
	conster := new(Token)
	var result *Token
	result = conster
	temp := make([]rune, 0, 2)
	isback := 0
	for index := 0; index < len(wordRune); {

		cur := wordRune[index]
		curNode = getWordByRune(curNode.children, cur)
		if curNode == nil || !curNode.isToekn {
			curNode = dic

			conster.Word = temp
			conster.Next = new(Token)
			conster = conster.Next
			temp = make([]rune, 0, 2)

			if isback == 1 {
				index++
				isback = 0
				continue
			}

			isback = 1
			continue
		}
		isback = 0
		temp = append(temp, cur)
		// step++
		index++
	}

	conster.Word = temp
	conster.Next = new(Token)
	conster = conster.Next

	return result
}
