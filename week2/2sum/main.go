package main

func main() {

}

func twoSum(nums []int, target int) []int {
	preNums := map[int]int{}
	res := []int{}

	for key, value := range nums {
		preNum := target - value
		preNumIndex, ok := preNums[preNum]

		if ok {
			return append(res, preNumIndex, key)
		}

		preNums[value] = key
	}
	return res
}
