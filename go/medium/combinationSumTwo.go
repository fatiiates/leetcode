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

func combinationSum2(candidates []int, target int) [][]int {
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
				index:     i + 1,
				remainder: target - v,
			})
		}
	}

	for i := 0; i < len(bt); i++ {
		node := bt[i]
		values := make(map[int]bool)
		for j := node.index; j < len(candidates); j++ {
			v := candidates[j]
			if _, ok := values[v]; ok {
				continue
			}
			values[v] = true
			if node.remainder == v {
				res = append(res, append(append([]int{}, node.arr...), v))
				break
			} else if v < node.remainder && node.remainder-v >= v {
				new_node := Node{
					arr:       append(append([]int{}, node.arr...), v),
					index:     j + 1,
					remainder: node.remainder - v,
				}
				bt = append(bt, new_node)
			}
		}
	}

	return res
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
