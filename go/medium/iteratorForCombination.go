package main

type CombinationIterator struct {
	str  string
	cl   int
	now  int
	l    int
	next string
}

func Constructor(characters string, combinationLength int) CombinationIterator {

	obj := CombinationIterator{
		str: characters,
		cl:  combinationLength,
		now: 0,
		l:   len(characters),
	}

	return obj
}

func (this *CombinationIterator) Next() string {
	if this.HasNext() {
		indexes := getIndexes(this.l, this.cl, this.now)

		str := ""
		for _, v := range indexes {
			str += string(this.str[v])
		}
		this.now += 1
		return str
	} else {
		return ""
	}
}

func (this *CombinationIterator) HasNext() bool {

	c := combination(this.l, this.cl)

	if this.now <= c-1 {
		return true
	}
	return false
}

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func combination(n int, k int) int {
	if n < k {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func getIndexes(n int, k int, m int) []int {
	res := []int{}
	a := n
	b := k
	x := (combination(n, k) - 1) - m
	for i := 0; i < k; i++ {
		a = a - 1
		for combination(a, b) > x {
			a = a - 1
		}
		res = append(res, n-1-a)
		x = x - combination(a, b)
		b = b - 1
	}

	return res
}

func main() {

	obj := Constructor("abc", 2)
	println(obj.Next())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())

}
