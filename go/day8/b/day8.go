package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetContent(filename string) []string {
	fileIn, fileInErr := os.Open(filename)
	if fileInErr != nil {
		fmt.Println("error!")
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)

	str, readerError := finReader.ReadString('\n')
	if readerError == io.EOF {
		return []string{}
	}
	str = strings.TrimSpace(str)
	arr := strings.Split(str, " ")
	return arr

}

type Node struct {
	nodeNum int
	metaNum int
	childs  []*Node
	metas   []int
	value   int
}

func NewNode() *Node {
	c := new(Node)
	c.nodeNum = 0
	c.metaNum = 0
	c.metas = []int{}
	c.childs = []*Node{}
	c.value = 0
	return c
}
func parseChild(parent *Node, seq *[]string) int {
	childNum, _ := strconv.Atoi((*seq)[0])
	metasNum, _ := strconv.Atoi((*seq)[1])
	node := NewNode()
	node.nodeNum = childNum
	node.metaNum = metasNum
	parent.childs = append(parent.childs, node)
	*seq = (*seq)[2:]
	result := 0
	if childNum > 0 {
		for i := 0; i < childNum; i++ {
			result += parseChild(node, seq)
		}
	}
	if metasNum > 0 {
		for i := 0; i < metasNum; i++ {
			v, _ := strconv.Atoi((*seq)[i])
			result += v
			fmt.Println(v)
			node.metas = append(node.metas, v)
		}
		*seq = (*seq)[metasNum:]
	}
	if childNum == 0 {
		for i := range node.metas {
			node.value += node.metas[i]
		}
	} else {
		for i := range node.metas {
			index := node.metas[i] - 1
			if index < childNum {
				node.value += node.childs[index].value
			}
		}
	}
	return result
}
func parse(seq *[]string) int {
	childNum, _ := strconv.Atoi((*seq)[0])
	metasNum, _ := strconv.Atoi((*seq)[1])
	node := NewNode()
	node.nodeNum = childNum
	node.metaNum = metasNum
	*seq = (*seq)[2:]
	result := 0
	if childNum > 0 {
		for i := 0; i < childNum; i++ {
			result += parseChild(node, seq)
		}
	}
	if metasNum > 0 {
		for i := 0; i < metasNum; i++ {
			v, _ := strconv.Atoi((*seq)[i])
			result += v
			fmt.Println(v)
			node.metas = append(node.metas, v)
		}
		*seq = (*seq)[metasNum:]
	}
	if childNum == 0 {
		for i := range node.metas {
			node.value += node.metas[i]
		}
	} else {
		for i := range node.metas {
			index := node.metas[i] - 1
			if index < childNum {
				node.value += node.childs[index].value
			}
		}
	}
	fmt.Println("root value ", node.value)
	return result
}

func main() {
	t1 := time.Now()
	seq := GetContent("/Volumes/MyDisk/studio/aoc/go/day8/input.txt")
	fmt.Println(seq)
	result := parse(&seq)

	fmt.Println("result :", result)

	t2 := time.Since(t1)

	fmt.Println("use time:", t2)

}
