package main

import "fmt"

func main() {
	num := 14
	result := numberOfSteps(num)
	fmt.Println(result)

}
func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {

			num = num / 2

		} else {

			num = num - 1

		}
		count++
	}
	//fmt.Println(num)
	return count
}
