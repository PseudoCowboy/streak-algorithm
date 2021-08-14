package main

func main() {
	unhappyFriends(4, [][]int{
		{1, 3, 2},
		{2, 3, 0},
		{1, 3, 0},
		{0, 2, 1},
	}, [][]int{
		{1, 3},
		{0, 2},
	})
}

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	m := make(map[int]int)
	count := 0
	for i := range pairs {
		m[pairs[i][0]] = pairs[i][1]
		m[pairs[i][1]] = pairs[i][0]
	}
	rank := make([][]int, n)
	for i := range rank {
		rank[i] = make([]int, n)
	}
	for i := range preferences {
		for priority, person := range preferences[i] {
			rank[i][person] = priority
		}
	}
	for i := range pairs {
		count += calc(pairs[i][0], pairs[i][1], rank, preferences, m)
		count += calc(pairs[i][1], pairs[i][0], rank, preferences, m)
	}
	return count
}

func calc(person int, pair int, rank, preferences [][]int, m map[int]int) int {
	priority := rank[person][pair]
	for i := 0; i < priority; i++ {
		prefer := preferences[person][i]
		preferspari := m[prefer]
		if rank[prefer][person] < rank[prefer][preferspari] {
			return 1
		}
	}
	return 0
}
