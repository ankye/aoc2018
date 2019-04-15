package main

import (
	"fmt"
	"time"
)

type Node struct {
	num      int
	preNode  *Node
	nextNode *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.num = value
	node.preNode = nil
	node.nextNode = nil
	return node
}

func main() {
	t1 := time.Now()
	root := NewNode(0)
	root.nextNode = root
	root.preNode = root
	players := 466
	step := 71436 * 100
	count := 0
	curNode := root

	wins := make([]int, players)
	maxValue := 0
	for {
		if count > step {
			break
		}

		for i := 0; i < players; i++ {
			count++
			value := count
			if count > step {
				break
			}
			//println("cur", count)
			if value%23 == 0 {
				wins[i] += value

				for j := 0; j < 7; j++ {

					curNode = curNode.preNode

				}
				wins[i] += curNode.num
				node := curNode.preNode
				curNode = curNode.nextNode
				curNode.preNode = node
				node.nextNode = curNode
				if maxValue < wins[i] {
					maxValue = wins[i]
				}
				//println(i, wins[i])
			} else {
				node := NewNode(count)

				curNode = curNode.nextNode

				nextNode := curNode.nextNode
				curNode.nextNode = node
				node.preNode = curNode
				nextNode.preNode = node
				node.nextNode = nextNode
				curNode = node

			}

		}
	}

	// curNode = root
	// for {
	// 	if curNode.nextNode == root {
	// 		break
	// 	}
	// 	print(curNode.num, " ")
	// 	curNode = curNode.nextNode
	// }
	println("MaxValue:", maxValue)
	println(curNode.num)
	t2 := time.Since(t1)

	fmt.Println("use time:", t2)

}
