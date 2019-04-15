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
	name      string
	char      int
	preNodes  []string
	nextNodes []string
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
		str = strings.Replace(str, "Step ", "", -1)
		str = strings.Replace(str, "must be finished before step ", "", -1)
		str = strings.Replace(str, "can begin.", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		arr := strings.Split(str, " ")
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

func main() {
	t1 := time.Now()
	lines := GetContent("/Volumes/MyDisk/studio/aoc/go/day7/input.txt")
	fmt.Println(lines)
	tree := make(map[string]Node)
	checkBooks := make([]string, 30)
	books := make([]string, 30)

	for i := range lines {
		line := lines[i]
		if _, ok := tree[line[0]]; !ok {
			v := int(line[0][0]) - 65
			tree[line[0]] = Node{line[0], v, []string{}, []string{}}
		}
		if _, ok := tree[line[1]]; !ok {
			v := int(line[1][0]) - 65
			tree[line[1]] = Node{line[1], v, []string{}, []string{}}
		}
		node, _ := tree[line[0]]
		nextNode, _ := tree[line[1]]

		nextNode.preNodes = append(nextNode.preNodes, node.name)
		node.nextNodes = append(node.nextNodes, nextNode.name)

		tree[line[1]] = nextNode
		tree[line[0]] = node

	}
	fmt.Println(tree)
	//	root := ""
	for k, v := range tree {
		if len(v.preNodes) == 0 {
			checkBooks[v.char] = k
		}
	}

	//fmt.Printf("Root : %s\n", root)
	fmt.Println("start Check == ")
	var rootNode *Node

	for {
		rootNode = nil
		for i := range checkBooks {
			if len(checkBooks[i]) > 0 {
				node, _ := tree[checkBooks[i]]
				check := true
				for j := range node.preNodes {
					name := node.preNodes[j]
					if _, ok := tree[name]; ok {
						check = false
						break
					}
				}
				if check {
					books[i] = checkBooks[i]
					checkBooks[i] = ""
				}
			}
		}
		for i := range books {
			if len(books[i]) > 0 {
				node, _ := tree[books[i]]
				rootNode = &node
				fmt.Print(rootNode.name)
				books[i] = ""
				delete(tree, node.name)
				break
			}
		}
		if rootNode == nil {
			break
		}
		if len(rootNode.nextNodes) > 0 {
			for i := range rootNode.nextNodes {
				name := rootNode.nextNodes[i]
				node, _ := tree[name]
				checkBooks[node.char] = node.name
				//fmt.Println("append name " + node.name)
			}
		}

	}
	fmt.Println("\nend Check == ")

	t2 := time.Since(t1)

	fmt.Println("use time:", t2)

}
