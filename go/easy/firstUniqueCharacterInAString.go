package main

import "fmt"

type Queue struct {
	c chan rune
}

func createQueue() Queue {
	return Queue{
		c: make(chan rune, 300),
	}
}

func (q *Queue) push(val rune) {
	q.c <- val
}

func (q *Queue) pop() rune {
	return <-q.c
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	q := createQueue()

	for i, v := range s {
		if _, ok := m[v]; ok {
			m[v] = -1
		} else {
			q.push(v)
			m[v] = i
		}
	}
	l := len(q.c)

	for i := 0; i < l; i++ {
		v := m[q.pop()]
		if v != -1 {
			return v
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
