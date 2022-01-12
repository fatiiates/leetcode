package main

func getMaxLen(nums []int) int {
	subs := [][]int{}
	subArr := []int{}
	for _, v := range nums {
		if v != 0 {
			subArr = append(subArr, v)
		} else if len(subArr) > 0 {
			subs = append(subs, subArr)
			subArr = []int{}
		}
	}
	subs = append(subs, subArr)

	// case_1 if there is a sub array with include negatives even
	case_1 := 0
	for _, v := range subs {
		negativeNumbers := 0
		for _, x := range v {
			if x < 0 {
				negativeNumbers++
			}
		}
		if negativeNumbers%2 == 0 {
			case_1 = max(case_1, len(v))
		}
	}
	if case_1 > 0 {
		return case_1
	}

	// case_2 remove prefix to first negative
	case_2 := 0
	for _, v := range subs {
		numbers := 0
		for _, x := range v {
			numbers++
			if x < 0 {
				break
			}
		}
		case_2 = max(case_2, len(v)-numbers)
	}

	// case_3 remove suffix to last negative
	case_3 := 0
	for _, v := range subs {
		numbers := 0
		for i := len(v) - 1; i >= 0; i-- {
			numbers++
			if v[i] < 0 {
				break
			}
		}
		case_3 = max(case_3, len(v)-numbers)
	}

	return max(case_1, max(case_2, case_3))
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
