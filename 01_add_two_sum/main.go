package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}

func twoSum(nums []int, target int) []int {

	remainSet := make(map[int]int, len(nums))

	for index, value := range nums {
		if i, exsist := remainSet[value]; exsist {
			return []int{i, index}
		}
		remainSet[target-value] = index
	}
	return nil
}
