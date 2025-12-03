package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
*

	只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。 可以使用 for 循环遍历数组，
	结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素

*
*/

func oneTimeNumber(number [10]int) int {

	mapA := make(map[int]int)

	for _, numb := range number {
		value, exist := mapA[numb]

		if exist {
			mapA[numb] = value + 1
		} else {
			mapA[numb] = 1
		}
	}
	for key, value := range mapA {
		if value == 1 {
			return key
		}

	}

	return 0

}

/*
*

	判断一个整数是否是回文数 （正序（从左向右）和倒序（从右向左）读都是一样的整数）

*
*/

func isPlind(x int) bool {
	if x < 0 {
		return false
	}
	chr := strconv.Itoa(x)
	star, end := 0, len(chr)-1
	for star < end {
		if chr[star] != chr[end] {
			return false
		}
		star++
		end--
	}
	return true
}

/*
*

	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

*
*/

func isValid(str string) bool {
	stack := []rune{}
	mapB := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if mapB[char] != top {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

/*
*

	查找字符串数组中的最长公共前缀

*
*/
func longest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	pre := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(pre) && j < len(strs[i]) && pre[j] == strs[i][j] {
			j++

		}
		pre = pre[:j]
		if pre == "" {
			break
		}

	}
	return pre

}

/*
*

	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

*
*/

func addOne(inta []int) []int {
	for i := len(inta) - 1; i >= 0; i-- {
		if inta[i] != 9 {
			inta[i]++
			return inta

		}
		inta[i] = 0

	}
	inta2 := make([]int, len(inta)+1)
	inta2[0] = 1
	return inta2

}

/*
*

	删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录
	不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。

*
*/

func delNum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(numbers); j++ {
		if numbers[j] != numbers[i] {
			i++
			numbers[i] = numbers[j]
		}

	}
	return i + 1

}

/*
*

	合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的
	区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
	将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中

*
*/

func sumIterv(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]

	})
	newintervals := [][]int{intervals[0]}

	for i := 0; i < len(intervals); i++ {
		last := newintervals[len(newintervals)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			newintervals = append(newintervals, intervals[i])
		}
	}
	return newintervals

}

/**

	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

**/

func findNum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num

		if j, exists := numMap[complement]; exists {

			return []int{nums[j], nums[i]}
		}
		numMap[num] = i
	}
	return []int{}

}
func main() {

	oneTimeNumbernums1 := [10]int{1, 1, 1, 2, 2, 3, 4, 4, 5, 5}
	fmt.Println("不重复数字：", oneTimeNumber(oneTimeNumbernums1))

	isPlindnumb2 := 12321
	fmt.Println(isPlind(isPlindnumb2))

	isValidstr := "()[]{}"
	fmt.Println("括号：", isValid(isValidstr))

	longeststrs := []string{
		"abc", "abcsd", "ab",
	}
	fmt.Println("最长前缀：", longest(longeststrs))

	addOnenumsAdd := []int{1, 1, 9}
	fmt.Println("addOne:", addOne(addOnenumsAdd))

	delNumNumbers := []int{0, 0, 1, 1, 2, 3, 4, 4, 5}
	fmt.Println("delNum:", delNum(delNumNumbers))

	sumItervintervals := [][]int{{1, 4}, {2, 3}, {5, 8}, {15, 18}}
	fmt.Println("sumIterv:", sumIterv(sumItervintervals))

	findNumNumbs := []int{1, 3, 5, 4, 6}
	targetNum := 4
	fmt.Println(findNum(findNumNumbs, targetNum))

}
