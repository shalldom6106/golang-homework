package main

import (
	"fmt"
	"sort"
)

func main() {
	//136. 只出现一次的数字
	arry := []int{1, 2, 3, 4, 1, 3}
	singleNum(arry)
	//回文数
	fmt.Println("判断是否为回文数：", isPalindrome(121))
	//有效的括号
	fmt.Println("括号是否有效：", isValid("([{]})"))
	//最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("最长公共前缀为：", longestCommonPrefix(strs))
	//删除排序数组中的重复项
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	a, b := removeDuplicates(nums)
	fmt.Println("去重后的数组长度和数组为：", a, b)
	//加一
	digits := []int{4, 3, 2, 1}
	fmt.Println("加一数组为：", plusOne(digits))
	//合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("合并区间后为：", merge(intervals))
	//两数之和
	andNums := []int{3, 2, 4}
	target := 6
	fmt.Println(target, "相加的数组为：", twoSum(andNums, target))

}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素
func singleNum(arr []int) {
	numMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]]++
	}
	for key, value := range numMap {
		if value == 1 {
			fmt.Println("数组唯一数：", key)
		}
	}
}

// 判断一个整数是否是回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := 0
	for x > num {
		num = num*10 + x%10
		x /= 10
	}
	// 当数字长度为奇数时，可以通过 num/10 去除处于中位的数字
	return x == num || x == num/10
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	slice := []rune{}
	runeMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		// 如果是左括号，则追加切片
		if c == '(' || c == '[' || c == '{' {
			slice = append(slice, c)
		} else {
			if len(slice) == 0 {
				return false
			}
			// 如果与最后一个元素不匹配，则返回 false
			if slice[len(slice)-1] != runeMap[c] {
				return false
			}
			// 出栈
			slice = slice[:len(slice)-1]
		}
	}
	return len(slice) == 0

}

// 编写一个函数来查找字符串数组中的最长公共前缀
// 如果不存在公共前缀，返回空字符串 ""
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sameString := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(strs) > 0 && len(strs[i]) < len(sameString) || strs[i][:len(sameString)] != sameString {
			sameString = sameString[:len(sameString)-1]
		}
	}
	return sameString
}

// 给定一个排序数组，你需要在原地删除重复出现的元素
func removeDuplicates(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, []int{}
	}
	ptr := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[ptr] {
			ptr++
			nums[ptr] = nums[i]
		}
	}
	return ptr + 1, nums[:ptr+1]
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	temp := append([]int{1}, digits...)
	return temp
}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	//按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//初始化合并后的区间
	mergeArry := [][]int{intervals[0]}
	for i := 0; i < len(intervals); i++ {
		//最后一个区间
		lastArry := mergeArry[len(mergeArry)-1]
		//当前区间
		currentArry := intervals[i]
		//区间重叠时合并区间
		if currentArry[0] <= lastArry[1] {
			lastArry[1] = max(lastArry[1], currentArry[1])
		} else { //否则添加区间
			mergeArry = append(mergeArry, currentArry)
		}
	}
	return mergeArry
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
	mm := make(map[int]int)
	for i, value := range nums {
		//计算剩余的数
		lastNum := target - value
		//查找剩余的数
		if j, otherValue := mm[lastNum]; otherValue {
			return []int{nums[j], nums[i]}
		}
		mm[value] = i
	}
	return nil //未找到整数返回空
}
