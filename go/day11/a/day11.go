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
func getFuelCell(x, y, sn int) int {
	value := ((x+10)*y + sn) * (x + 10)
	value = int((value % 1000) / 100)
	return value - 5
}

func main() {
	t1 := time.Now()

	sn := 9221
	g := map[int]map[int]int{}
	maxValue := 0
	pos := []int{}
	MAX := 300
	for y := 1; y <= MAX; y++ {
		for x := 1; x <= MAX; x++ {
			mm, ok := g[y]
			if !ok {
				mm = make(map[int]int)
				g[y] = mm
			}
			mm[x] = getFuelCell(x, y, sn)

		}
	}
	for y := 1; y <= MAX-2; y++ {
		for x := 1; x <= MAX-2; x++ {
			v := 0
			for y1 := y; y1 <= y+2; y1++ {
				for x1 := x; x1 <= x+2; x1++ {
					v += g[y1][x1]

				}
			}
			if maxValue < v {
				maxValue = v
				pos = []int{x, y}
			}
		}
	}

	t2 := time.Since(t1)
	fmt.Println("Max=", maxValue)
	fmt.Println(pos)
	fmt.Println("use time:", t2)

}
