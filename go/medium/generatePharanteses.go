package main

import "fmt"

type Node struct {
	s string
	c int
}

func generateParenthesis(n int) []string {
	arr := []string{}

	bt := []Node{
		Node{
			s: "(",
			c: 1,
		},
	}
	for len(bt) > 0 {
		p := popFromQueue(&bt)
		if len(p.s)+p.c == n*2 {
			close := ""
			for i := 0; i < p.c; i++ {
				close += ")"
			}
			arr = append(arr, p.s+close)
		} else if p.c == 0 {
			new_node := Node{
				s: p.s + "(",
				c: 1,
			}
			bt = append(bt, new_node)
		} else {
			bt = append(bt, Node{
				s: p.s + "(",
				c: p.c + 1,
			})
			bt = append(bt, Node{
				s: p.s + ")",
				c: p.c - 1,
			})
		}
	}

	return arr
}

func popFromQueue(n *[]Node) Node {
	f := (*n)[0]
	(*n) = (*n)[1:]
	return f
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
