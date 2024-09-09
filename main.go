package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type DisjointSet struct {
	parent []int
	sum    []int
}

func NewDisjointSet(n int, population []int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		sum:    make([]int, n),
	}
	for i := 0; i < n; i++ {
		ds.parent[i] = i
		ds.sum[i] = population[i]
	}
	return ds
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSet) Union(x, y int) {
	px, py := ds.Find(x), ds.Find(y)
	if px == py {
		return
	}
	ds.parent[py] = px
	ds.sum[px] += ds.sum[py]
}

func (ds *DisjointSet) GetSum(x int) int {
	return ds.sum[ds.Find(x)]
}

type Event struct {
	time   int
	typ    int // 0 for query, 1 for modification
	cities []int
	index  int // For queries, to maintain original order
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		var x int
		fmt.Sscanf(scanner.Text(), "%d", &x)
		return x
	}

	n, m, q := readInt(), readInt(), readInt()
	population := make([]int, n)
	for i := 0; i < n; i++ {
		population[i] = readInt()
	}

	events := make([]Event, 0, m+q)
	for i := 0; i < m; i++ {
		t, c := readInt(), readInt()
		cities := make([]int, c)
		for j := 0; j < c; j++ {
			cities[j] = readInt() - 1 // 转换为0-索引
		}
		events = append(events, Event{time: t, typ: 1, cities: cities})
	}

	for i := 0; i < q; i++ {
		t, s := readInt(), readInt()
		events = append(events, Event{time: t, typ: 0, cities: []int{s - 1}, index: i}) // 转换为0-索引
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].typ > events[j].typ
		}
		return events[i].time < events[j].time
	})

	ds := NewDisjointSet(n, population)
	results := make([]int, q)

	for _, event := range events {
		if event.typ == 1 {
			for i := 1; i < len(event.cities); i++ {
				ds.Union(event.cities[0], event.cities[i])
			}
		} else {
			results[event.index] = ds.GetSum(event.cities[0])
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
