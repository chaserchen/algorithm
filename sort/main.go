package main

import (
	"fmt"
)

func main() {
	x := []int{4, 3, 2, 1, 5, -1}
	selection(x)
	fmt.Println(x)
}

func bubble(x []int) {  // 冒泡排序
	l := len(x)
	if l <= 1 {
		return
	}

	for i := 0; i < l - 1; i++ {  // 最多需要冒泡 len - 1 次
		hasChange := false
		for j := 0; j < l - 1 - i; j++ {
			if x[j] < x[j + 1] {
				value := x[j]
				x[j] = x[j + 1]
				x[j + 1] = value
				hasChange = true
			}
		}
		if !hasChange {  // 如果有一次冒泡过程中, 没有发生数据交换, 则说明已完成排序, 提前退出
			break
		}
	}
}

func insertion(x []int)  {  // 插入排序
	l := len(x)
	if l <= 1 {
		return
	}

	for i := 1; i < l; i++ {  // 初始已排序区域长度为1(index为0的区域), 遍历未排序区域
		value := x[i]
		for j := i - 1; j >= 0; j-- {  // 倒序遍历已排序区域
			if value < x[j] {  // 若发现更小值, 则做交换
				x[j + 1] = x[j]
				x[j] = value
			} else {
				break  // 若发现该值小于已排序区域中的任一值, 则不再需要交换, 继续处理下一个未排序区域数据
			}
		}
	}
}

func selection(x []int)  {  // 选择排序
	l := len(x)
	if l <= 1 {
		return
	}

	for i:= 0; i < l; i++ {  // 初始已排序区域长度为0, 控制添加已排序区域的末尾元素, 末尾元素为未排序区域的最小元素 
		minIndex := i
		for j := i + 1; j < l; j++ {  // 找到未排序区域的最小元素
			if x[j] < x[minIndex] {
				minIndex = j
			}
		}
		if minIndex > i {  // 若最小值不是在未排序区域的第一位, 则交换 
			minValue := x[minIndex]
			x[minIndex] = x[i]
			x[i] = minValue
		}
	}
}

