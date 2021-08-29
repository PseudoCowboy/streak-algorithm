package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSessions([]int{2, 3, 3, 4, 4, 4, 5, 6, 7, 10}, 12))
}

func minSessions(tasks []int, sessionTime int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	fmt.Println(tasks)
	used := make([]int, len(tasks))
	total := 0
	for i := 0; i < len(tasks); i++ {
		if used[i] == 0 {
			used[i]++
			subused := make([]int, len(tasks[i+1:]))
			patch(subused, used[i+1:])
			min := sessionTime
			ans := make([]int, 0)
			dfs(tasks[i+1:], sessionTime-tasks[i], &subused, &ans, []int{}, &min, 0)
			for j := range ans {
				used[i+1+ans[j]]++
			}
			total++
		}
	}
	return total
}

func patch(a, b []int) {
	for i := range a {
		a[i] = b[i]
	}
}

func dfs(task []int, target int, used *[]int, ans *[]int, current []int, min *int, start int) {
	if target < *min {
		*min = target
		temp := make([]int, len(current))
		copy(temp, current)
		*ans = temp
	}

	for i := start; i < len(task); i++ {
		if (*used)[i] == 0 && target-task[i] >= 0 {
			(*used)[i]++
			current = append(current, i)
			dfs(task, target-task[i], used, ans, current, min, i+1)
			current = current[:len(current)-1]
			(*used)[i]--
		}
	}
}
