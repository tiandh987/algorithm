package main

import "strconv"

func main() {

}

// leetcode 412
// Fizz Buzz
// https://leetcode-cn.com/problems/fizz-buzz/

/**
 * 方法一：模拟法
 * 思路：每遇到一个数字就判断，能被3整除？能被5整除？都能被整除？
 */
func fizzBuzz(n int) []string {
	var res []string

	for i := 1; i <= n; i++ {
		divisibleBy3 := (i%3 == 0)
		divisibleBy5 := (i%5 == 0)

		if divisibleBy3 && divisibleBy5 {
			res = append(res, "FizzBuzz")
		} else if divisibleBy3 {
			res = append(res, "Fizz")
		} else if divisibleBy5 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

/**
 * 方法二：字符串拼接
 * 思路：对方法一的改进
 *   能被3整除        -> fizz
 *   能被5整除        -> Buzz
 *   能被3 和 5 整除  -> fuzzBuzz
 *   不能被整除       -> 返回数字
 *
 *   在方法一中只有两个映射关系，如果更多，将会有更多的if分支判断条件
 */
func fizzBuzz(n int) []string {
	var res []string

	for i := 1; i <= n; i++ {
		divisibleBy3 := (i%3 == 0)
		divisibleBy5 := (i%5 == 0)

		numAnsStr := ""
		if divisibleBy3 {
			numAnsStr += "Fizz"
		}

		if divisibleBy5 {
			numAnsStr += "Buzz"
		}

		if numAnsStr == "" {
			numAnsStr += strconv.Itoa(i)
		}

		res = append(res, numAnsStr)
	}

	return res
}

/**
 * 方法三：hash表
 * 思路：前面的两种方法中，映射关系都是固定的。
 *   如果要替换掉映射关系，就需要一个一个的修改判断条件，非常不利于维护代码。
 *   可以将这种映射关系放在hash表中。
 */
func fizzBuzz(n int) []string {
	var res []string

	dict := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	key := []int{3, 5}

	for i := 1; i <= n; i++ {
		numAnsStr := ""
		//for key, value := range dict {
		//	divisible := (i % key == 0)
		//	if divisible {
		//		numAnsStr += value
		//	}
		//}

		// 直接遍历map会出现“BuzzFizz”,因为map遍历是无序的
		for _, key := range keys {
			divisible := (i%key == 0)
			if divisible {
				numAnsStr += dict[key]
			}
		}
		if numAnsStr == "" {
			numAnsStr = strconv.Itoa(i)
		}
		res = append(res, numAnsStr)
	}
	return res
}
