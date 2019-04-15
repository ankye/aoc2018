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

type Node struct {
	name string
	pid  int
	step int
	x    int
	y    int
}
type Point struct {
	x int
	y int
}
type Opt struct {
	step int
	x    int
	y    int
	pid  int
}

const MAX = 400

func GetName(num int, upper bool) string {
	num = num - 1
	l := "abcdefghijklmnopqrstuvwxyz"
	one := int(num / len(l))
	two := num % len(l)

	s := fmt.Sprintf("%s%s", string(l[one]), string(l[two]))
	if one == 0 {
		s = string(l[two])
	}
	if upper {
		s = strings.ToUpper(s)
	}
	return s
}

func PK(g *[][]Node, x, y, pid, step int) bool {
	node := (*g)[x][y]

	if node.step == 0 {
		return false
	}

	if node.pid != pid {
		if node.step > step+1 {
			node.name = GetName(pid, false)
			node.pid = pid
			node.step = step + 1
			(*g)[x][y] = node
			return true
		} else if node.step == step+1 {
			node.name = "-"
			node.pid = 0
			(*g)[x][y] = node
			return true
		} else {
			return false
		}
	} else {
		if node.step > step+1 {
			node.name = GetName(pid, false)
			node.pid = pid
			node.step = step + 1
			(*g)[x][y] = node
			return true
		}
		return false
	}
}

func WalkN(g *[][]Node, step, x, y, pid int) bool {

	if x < 0 || x >= len(*g) {
		return false
	}
	if y < 0 || y >= len(*g) {
		return false
	}
	return PK(g, x, y, pid, step)
}

func WalkNode(g *[][]Node, searchQ *[]Opt, step, x, y, pid int) int {
	if pid == 0 {
		fmt.Println("=========")
	}
	if step > MAX {
		return 0
	}
	count := 0
	if WalkN(g, step, x, y-1, pid) {
		*searchQ = append(*searchQ, Opt{step + 1, x, y - 1, pid})
		count++
	}
	if WalkN(g, step, x, y+1, pid) {
		*searchQ = append(*searchQ, Opt{step + 1, x, y + 1, pid})
		count++
	}
	if WalkN(g, step, x-1, y, pid) {
		*searchQ = append(*searchQ, Opt{step + 1, x - 1, y, pid})
		count++
	}
	if WalkN(g, step, x+1, y, pid) {
		*searchQ = append(*searchQ, Opt{step + 1, x + 1, y, pid})
		count++
	}
	return count
}

func GetContent(filename string) []Point {
	fileIn, fileInErr := os.Open(filename)
	if fileInErr != nil {
		fmt.Println("error!")
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	var fileList []Point
	for {
		str, readerError := finReader.ReadString('\n')

		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		arr := strings.Split(str, ",")
		if len(arr) == 2 {
			//fmt.Println(arr)
			x, _ := strconv.Atoi(arr[1])
			y, _ := strconv.Atoi(arr[0])
			fileList = append(fileList, Point{x, y})
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
	lines := GetContent("/Volumes/MyDisk/studio/aoc/day6/input.txt")
	fmt.Println(lines)

	g := make([][]Node, MAX, MAX) //建议第一维
	for i := 0; i < MAX; i++ {
		m2 := make([]Node, MAX, MAX) //可用循环对m2赋值，默认建立初值为0
		g[i] = m2                    //建立第二维
		for j := 0; j < MAX; j++ {
			m2[j] = Node{"", 0, 999999, 0, 0}
		}
	}

	num := 1
	nodes := []Node{}

	for _, p := range lines {
		name := GetName(num, true)
		node := g[p.x][p.y]
		node.name = name
		node.pid = num
		node.step = 0
		node.x = p.x
		node.y = p.y
		g[p.x][p.y] = node
		num++
		//fmt.Println(p)
		nodes = append(nodes, node)
	}

	closeMap := make(map[int]Node)
	counterMap := make(map[int]int)

	searchQ := []Opt{}

	for _, node := range nodes {
		step := 0
		closeMap[node.pid] = node
		counterMap[node.pid] = 0

		count := WalkNode(&g, &searchQ, step, node.x, node.y, node.pid)
		if count > 0 {
			//do something

		}
	}

	for {
		if len(searchQ) == 0 {
			break
		}
		q := searchQ[0]
		WalkNode(&g, &searchQ, q.step, q.x, q.y, q.pid)
		searchQ = searchQ[1:]
	}
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			node := g[i][j]
			fmt.Print(node.name + " ")
		}
		fmt.Print("\n")
	}

	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			node := g[i][j]

			if node.pid > 0 {
				counterMap[node.pid]++
			}
		}
	}

	for i := 0; i < MAX; i++ {
		node := g[0][i]
		if _, ok := closeMap[node.pid]; ok {
			delete(closeMap, node.pid)
		}
		node = g[MAX-1][i]
		if _, ok := closeMap[node.pid]; ok {
			delete(closeMap, node.pid)
		}
		node = g[i][0]
		if _, ok := closeMap[node.pid]; ok {
			delete(closeMap, node.pid)
		}
		node = g[i][MAX-1]
		if _, ok := closeMap[node.pid]; ok {
			delete(closeMap, node.pid)
		}
	}

	fmt.Println(closeMap)

	maxNum := 0
	maxK := ""
	key := 0
	for k, _ := range closeMap {
		if maxNum < counterMap[k] {
			maxNum = counterMap[k]
			maxK = GetName(k, true)
			key = k
		}
	}

	fmt.Println(counterMap)

	fmt.Println(maxNum, maxK, key)

	t2 := time.Since(t1)
	fmt.Println("sum=", 0)
	fmt.Println("use time:", t2)

}
