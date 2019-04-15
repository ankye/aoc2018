package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Node struct {
	value    string
	nextNode *Node
	preNode  *Node
}

func NewNode(value string) *Node {
	node := new(Node)
	node.value = value
	node.nextNode = nil
	node.preNode = nil
	return node
}

func GetContent(filename string) [][]string {
	fileIn, fileInErr := os.Open(filename)
	if fileInErr != nil {
		fmt.Println("error!")
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	var fileList [][]string
	for {
		str, readerError := finReader.ReadString('\n')
		str = strings.TrimSpace(str)
		str = strings.Replace(str, "initial state: ", "data => ...", -1)
		arr := strings.Split(str, "=>")

		//fmt.Println(str)
		if len(arr) >= 2 {

			fileList = append(fileList, []string{strings.TrimSpace(arr[0]), strings.TrimSpace(arr[1])})
		}
		if readerError == io.EOF {
			break
		}

	}
	//fmt.Println("fileList", fileList)
	return fileList
}

func logArr(node *Node) {
	curNode := node
	for {
		print(curNode.value)
		if curNode.nextNode == nil {
			break
		}
		curNode = curNode.nextNode
	}
	print("\n")
}
func main() {
	t1 := time.Now()

	lines := GetContent("/Volumes/MyDisk/studio/aoc/go/day12/test.txt")
	fmt.Println(lines)
	g := map[string]string{}

	content := lines[0][1]
	root := NewNode(string(content[0]))
	curNode := root
	for i := 1; i < len(content); i++ {
		node := NewNode(string(content[i]))
		curNode.nextNode = node
		node.preNode = curNode
		curNode = node
	}
	for i := 0; i < 20; i++ {
		node := NewNode(".")
		curNode.nextNode = node
		node.preNode = curNode
		curNode = node
	}
	for i := 1; i < len(lines); i++ {
		g[lines[i][0]] = lines[i][1]
	}
	logArr(root)
	curNode = root

	for {
		if curNode.preNode == nil || curNode.preNode.preNode == nil {
			curNode = curNode.nextNode
			continue
		}
		key := curNode.preNode.preNode.value
		key += curNode.preNode.value
		key += curNode.value
		key += curNode.nextNode.value
		key += curNode.nextNode.nextNode.value
		if v, ok := g[key]; ok {
			curNode.preNode.preNode.value = "."
			curNode.preNode.value = "."
			curNode.nextNode.value = "."
			curNode.value = v
			curNode.nextNode.nextNode.value = "."
		}
		curNode = curNode.nextNode

		if curNode.nextNode != nil && curNode.nextNode.nextNode == nil {
			break
		}
	}
	logArr(root)

	t2 := time.Since(t1)

	fmt.Println("use time:", t2)

}
