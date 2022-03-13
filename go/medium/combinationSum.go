package main

import (
	"fmt"
	"sort"
)

type Node struct {
	arr       []int
	index     int
	remainder int
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	m := make(map[int]bool)
	bt := []Node{}
	sort.Ints(candidates)

	for i, v := range candidates {
		if v > target {
			continue
		}
		if _, ok := m[v]; !ok {
			m[v] = true
			if v == target {
				res = append(res, []int{v})
				continue
			}
			bt = append(bt, Node{
				arr:       []int{v},
				index:     i,
				remainder: target - v,
			})
		}
	}

	for i := 0; i < len(bt); i++ {
		node := bt[i]
		for j := node.index; j < len(candidates); j++ {
			fmt.Println(node)
			v := candidates[j]
			if node.remainder == v {
				res = append(res, append(append([]int{}, node.arr...), v))
				break
			} else if v < node.remainder && node.remainder-v >= v {
				new_node := Node{
					arr:       append(append([]int{}, node.arr...), v),
					index:     j,
					remainder: node.remainder - v,
				}
				bt = append(bt, new_node)
			}
		}

	}

	return res
}

func main() {
	// fmt.Println(combinationSum([]int{2, 3, 6, 7, 1}, 7))
	// fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}
