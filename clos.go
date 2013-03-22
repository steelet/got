package main

import (
	"strconv"
)

func addition(a int, b int) (int, string) {
	return a + b, "(" + strconv.Itoa(a) + "+" + strconv.Itoa(b) + ")"
}

func multi(a int, b int) (int, string) {
	return a * b, "(" + strconv.Itoa(a) + "*" + strconv.Itoa(b) + ")"
}

func makeMathClosure(base int, fn func(int, int) (int, string)) func(int) int {
	last_answer := -1
  	return func(num int) int {
		print(last_answer)
		answer, txt := fn(base, num)
		print(txt)
		last_answer = answer
		return answer 
	}
}

func main() {
	timesTwo := makeMathClosure(2, multi)
	plusTwo := makeMathClosure(2, addition)
	println(timesTwo(4))
	println(plusTwo(timesTwo(4)))
	println(plusTwo(4))
}
