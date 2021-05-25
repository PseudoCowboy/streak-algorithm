package main

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回表达式的值
 * @param s string字符串 待计算的表达式
 * @return int整型
 */
func solve(s string) int {
	// 把每次运算的结果放到该栈中
	stack := make([]int, 0)

	// 字符串转换为字符切片
	data := /*[]byte(s)*/ s

	// 表示字符串中每一个数字
	number := 0
	// 初识状态操作符
	operation := byte('+')

	// 处理字符串
	for i := 0; i < len(data); i++ {
		// 当前字符
		c := data[i]

		// 0~9 字符是数字
		if c >= '0' && c <= '9' {
			number = 10*number + int(c-'0')
			//continue 不能写continue，因为如果是最后一个字符的话，需要特殊处理
		}

		// '('
		// 遇到左括号特殊处理，还会进入到递归中
		if c == '(' {
			// 遇到的第一个左括号
			count := 1
			// 进入递归需要处理的开始于截止
			start, end := i+1, i+1

			// 每一次遇到括号，就是一层递归
			// 每一次递归都是去掉最外层的一对括号，直到没有括号
			// count == 0 就是最有一个右括号
			for count != 0 {
				// 遇到左括号 ++
				if data[end] == '(' {
					count++
				}
				//遇到右括号 --
				if data[end] == ')' {
					count--
				}

				end++
			}

			// 已经处理好的字符位置
			i = end - 1

			// 进入递归，并顺便传入需要处理的字符串
			// start, end-1 为需要处理的字符串的开始于截止位置
			number = solve(s[start:end])
		}

		// c < '0' || c > '9' 表示 '+' / '-' / '*'
		// i == len(data)-1 表示已经是最后一个字符了，此时需要进行最后的一次运算
		if c == '+' || c == '-' || c == '*' || i == len(data)-1 {
			switch operation {
			case '+':
				stack = append(stack, number)
			case '-':
				stack = append(stack, -number)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * number
			}

			// 重置数字
			number = 0

			// 当前的操作符号，前面已经对括号进行了处理，这里不会有括号
			//  i == len(data)-1 时，operation 无意义了
			operation = c
		}
	}

	// 累加栈中每一个数字
	result := 0
	for len(stack) != 0 {
		result += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return result
}
