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
	x  int
	y  int
	cx int
	cy int
	vx int
	vy int
	p  string
}

func NewNode(x, y, vx, vy int) *Node {
	node := new(Node)
	node.x = x
	node.y = y
	node.cx = x
	node.cy = y
	node.vx = vx
	node.vy = vy
	node.p = "#"
	return node
}

func GetContent(filename string) [][]int {
	fileIn, fileInErr := os.Open(filename)
	if fileInErr != nil {
		fmt.Println("error!")
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	var fileList [][]int
	for {
		str, readerError := finReader.ReadString('\n')
		str = strings.TrimSpace(str)
		str = strings.Replace(str, "position=<", "", -1)
		str = strings.Replace(str, "> velocity=<", ",", -1)
		str = strings.Replace(str, ">", "", -1)
		arr := strings.Split(str, ",")
		//fmt.Println(str)
		if len(arr) >= 2 {
			x, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(arr[1]))

			vx, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
			vy, _ := strconv.Atoi(strings.TrimSpace(arr[3]))

			fileList = append(fileList, []int{x, y, vx, vy})
		}
		if readerError == io.EOF {
			break
		}

	}
	//fmt.Println("fileList", fileList)
	return fileList
}
func isExist(g map[int]map[int]int, x, y int) bool {
	result := false
	if v, ok := g[y][x]; ok && v == 1 {
		result = true
	}
	return result
}
func main() {
	t1 := time.Now()
	lines := GetContent("/Volumes/MyDisk/studio/aoc/go/day10/input.txt")
	fmt.Println(lines)
	g := map[int]map[int]int{}

	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	nodes := []*Node{}
	for i := range lines {
		line := lines[i]
		x := line[0]
		y := line[1]
		if maxX < x {
			maxX = x
		}
		if minX > x {
			minX = x
		}
		if maxY < y {
			maxY = y
		}
		if minY > y {
			minY = y
		}

		mm, ok := g[y]
		if !ok {
			mm = make(map[int]int)
			g[y] = mm
		}
		mm[x] = 1
		nodes = append(nodes, NewNode(x, y, line[2], line[3]))

	}
	fmt.Println(minX, maxX, minY, maxY)
	step := 0
	maxStep := 500000
	for {
		if step > maxStep {
			break
		}
		r1 := isExist(g, maxX-1, maxY) && isExist(g, maxX-2, maxY)
		r2 := isExist(g, maxX, maxY-1) && isExist(g, maxX, maxY-2)
		if isExist(g, maxX, maxY) && (r1 || r2) {
			fmt.Println("==========================", step)

			for i := minY; i <= maxY; i++ {
				for j := minX; j <= maxX; j++ {
					if v, ok := g[i][j]; ok && v == 1 {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Print("\n")
			}
		}
		step++
		g = map[int]map[int]int{}
		maxX = 0
		minX = 9999
		maxY = 0
		minY = 9999
		//LCPGPXGL
		for i := range nodes {
			node := nodes[i]
			node.cx += node.vx
			node.cy += node.vy

			x := node.cx
			y := node.cy
			if maxX < x {
				maxX = x
			}
			if minX > x {
				minX = x
			}
			if maxY < y {
				maxY = y
			}
			if minY > y {
				minY = y
			}

			mm, ok := g[y]
			if !ok {
				mm = make(map[int]int)
				g[y] = mm
			}
			mm[x] = 1

		}

	}

	t2 := time.Since(t1)
	fmt.Println("sum=", 0)
	fmt.Println("use time:", t2)

}
