package segmentGo

import (
	"bufio"
	// "fmt"
	sysio "io"
	"os"
)

// 前缀树节点
type DicTree struct {
	bit      int8       // 该节点对应的字元
	isWord   bool       // 当此节点没有对应的分词时值为nil
	isToekn  bool       // 当此节点没有对应的分词时值为nil
	children []*DicTree // 该字元后继的所有可能字元，当为叶子节点时为空
	parent   *DicTree
}

//添加词 在词添加完成后标识这是一个词
func (dic *DicTree) addToken(word string, isReverse bool) {
	wordRune := []rune(word)
	var endWord *DicTree
	var index int
	var cur rune
	if isReverse {
		index = len(wordRune) - 1
		for ; index >= 0; index-- {
			cur = wordRune[index]
			endWord = dic.updateDic(endWord, stringToIntArr(cur))
		}
		// endWord.isWord = true
	} else {
		index = 0
		for ; index < len(wordRune); index++ {
			cur = wordRune[index]
			endWord = dic.updateDic(endWord, stringToIntArr(cur))
		}
		endWord.isWord = true
	}

}

//更新Node Node维持一个可以直接索引的字典树
//dic维持这个字典树的根
func (this *DicTree) updateNode(cur int8) *DicTree {
	if this.bit == -1 {
		this.bit = cur
		this.children = make([]*DicTree, 10)
		for index, _ := range this.children {
			this.children[index] = &DicTree{-1, false, false, nil, this}
		}
	} else {

		if this.children[cur].bit == -1 {
			return this.children[cur].updateNode(cur)
		} else {
			return this.children[cur]
		}
	}
	return this
}

//更新Dic 在当前Token添加完成后标识这个是一个字
func (dic *DicTree) updateDic(node *DicTree, token []int8) *DicTree {
	if len(dic.children) == 0 {
		dic.children = make([]*DicTree, 10)
		for index, _ := range dic.children {
			dic.children[index] = &DicTree{-1, false, false, nil, nil}
		}
	}
	length := len(token)
	firstToken := token[length-1]
	var result *DicTree
	var index int
	if node == nil {
		if dic.children[firstToken].bit == -1 {
			result = dic.children[firstToken].updateNode(firstToken)
		} else {
			result = dic.children[firstToken]
		}
		index = 1
	} else {
		result = node
		index = 0
	}

	for ; index < length; index++ {
		result = result.updateNode(token[length-1-index])
	}
	result.isToekn = true
	return result
}

func (dic *DicTree) GetWord(word string) bool {
	wordRune := []rune(word)
	lastNode := dic
	for _, cur := range wordRune {
		lastNode = getWordByRune(lastNode.children, cur)
		// fmt.Println(lastNode)
		if lastNode == nil || !lastNode.isToekn {
			return false
		}
	}
	if !lastNode.isWord {
		return false
	}
	return true
}

func getWordByRune(curDic []*DicTree, cur rune) *DicTree {

	var lastNode *DicTree

	curArr := stringToIntArr(cur)
	length := len(curArr) - 1
	for index := 0; index <= length; index++ {
		curInt := curArr[length-index]
		if curDic[curInt].bit == -1 {
			return nil
		}
		lastNode = curDic[curInt]
		curDic = curDic[curInt].children
	}

	if !lastNode.isToekn {
		return nil
	}
	return lastNode
}

func (dic *DicTree) LoadDic(path string) {

	f, err := os.Open(path)
	defer f.Close()
	if nil != err {
		return
	}

	buff := bufio.NewReader(f)
	for {
		line, err := buff.ReadString('\n')

		dic.addToken(process(line), false)

		if err != nil || sysio.EOF == err {
			break
		}
	}

}

func (dic *DicTree) LoadDic4Reverse(path string) {

	f, err := os.Open(path)
	defer f.Close()
	if nil != err {
		return
	}

	buff := bufio.NewReader(f)
	for {
		line, err := buff.ReadString('\n')

		dic.addToken(process(line), true)

		if err != nil || sysio.EOF == err {
			break
		}
	}

}
