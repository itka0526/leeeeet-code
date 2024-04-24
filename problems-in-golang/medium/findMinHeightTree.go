package medium

import (
	"fmt"
	"math"
)

type Tree map[int][]int

func buildTree(s [][]int) (Tree, map[int]bool) {
	t := Tree{}
	nodes := map[int]bool{}
	for i := range s {
		nodes[s[i][0]] = true
		nodes[s[i][1]] = true
	}
	for node := range nodes {
		tmp := []int{}
		for i := range s {
			if s[i][0] == node {
				tmp = append(tmp, s[i][1])
			}
			if s[i][1] == node {
				tmp = append(tmp, s[i][0])
			}
		}
		t[node] = tmp
	}
	return t, nodes
}

func dfsMHT(t Tree, node int, curr int, nodes int, lvl *int, vis *map[int]bool) {
	if (*vis)[node] {
		return
	}
	(*vis)[node] = true
	if curr > *lvl {
		*lvl = curr
	}
	for _, conn := range t[node] {
		dfsMHT(t, conn, curr+1, nodes, lvl, vis)
	}
}

func helper(n int, e [][]int) []int {
	t, nodes := buildTree(e)
	res := map[int][]int{}
	for node := range nodes {
		if len(t[node]) <= 1 && len(t) > 2 {
			continue
		}
		l := math.MinInt
		dfsMHT(t, node, 0, n, &l, &map[int]bool{})
		if l == math.MinInt {
			continue
		}
		if _, ok := res[l]; !ok {
			res[l] = []int{node}
		} else {
			res[l] = append(res[l], node)
		}
	}
	minLvl := math.MaxInt
	for k := range res {
		if minLvl > k {
			minLvl = k
		}
	}
	return res[minLvl]
}

func findMinHeightTreesDFS(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// dfs approach TLE
	return helper(n, edges)
}

type M struct {
	root  int
	value int
	level int
}

func findMinHeightTreesFailedBFS(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	t, nodes := buildTree(edges)
	res := []M{}
	for root := range nodes {
		fmt.Println("__________________")
		best := M{root: root, value: root, level: 0}
		q := []M{best}
		mem := make(map[int]bool)
		for len(q) > 0 {
			c := q[0]
			q = q[1:]
			fmt.Printf("%+v\n", c)
			if c.level > best.level {
				best = c
			}
			if mem[c.value] {
				continue
			}
			mem[c.value] = true
			for _, conn := range t[c.value] {
				q = append(q, M{root: root, value: conn, level: c.level + 1})
			}
		}
		if len(res) < 1 {
			res = append(res, best)
		} else if res[0].level > best.level {
			res = []M{best}
		} else if res[0].level == best.level {
			res = append(res, best)
		}
	}
	ans := make([]int, len(res))
	for i := range ans {
		ans[i] = res[i].root
	}
	return ans
}

type Edge int

type Graph map[Edge]*Info

type Info struct {
	connCount int
	conn      []Edge
}

func findMinHeightTreesWorks(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := Graph{}
	for _, pair := range edges {
		a := Edge(pair[0])
		b := Edge(pair[1])
		if _, ok := g[a]; ok {
			g[a].conn = append(g[a].conn, b)
		} else {
			g[a] = &Info{connCount: 0, conn: []Edge{b}}
		}
		if _, ok := g[b]; ok {
			g[b].conn = append(g[b].conn, a)
		} else {
			g[b] = &Info{connCount: 0, conn: []Edge{a}}
		}
		g[a].connCount += 1
		g[b].connCount += 1
	}
	for k, v := range g {
		fmt.Println(k, *v)
	}
	q := []Edge{}
	for k, v := range g {
		if v.connCount == 1 {
			q = append(q, k)
		}
	}
	for n > 2 {
		s := len(q)
		n -= s
		for s > 0 {
			curr := q[0]
			q = q[1:]
			for _, conn := range g[curr].conn {
				g[conn].connCount -= 1
				if g[conn].connCount == 1 {
					q = append(q, conn)
				}
			}
			s -= 1
		}
	}
	ans := make([]int, len(q))
	for i := range q {
		ans[i] = int(q[i])
	}
	return ans
}

type (
	GEdge int
	GConn map[GEdge]bool
)
type GGraph map[GEdge]GConn

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := GGraph{}
	for _, e := range edges {
		a, b := GEdge(e[0]), GEdge(e[1])
		if _, ok := g[a]; ok {
			g[a][b] = true
		} else {
			g[a] = GConn{b: true}
		}
		if _, ok := g[b]; ok {
			g[b][a] = true
		} else {
			g[b] = GConn{a: true}
		}
	}
	fmt.Println(g)
	q := []GEdge{}
	for e, conn := range g {
		if len(conn) == 1 {
			q = append(q, e)
		}
	}
	for n > 2 {
		s := len(q)
		n -= s
		for s > 0 {
			c := q[0]
			q = q[1:]
			for conn := range g[c] {
				delete(g[conn], c)
				if len(g[conn]) == 1 {
					q = append(q, conn)
				}
			}
			s -= 1
		}
	}
	ans := []int{}
	for _, k := range q {
		ans = append(ans, int(k))
	}
	return ans
}

func TestFindMinHeightTrees() {
	s := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	fmt.Println(findMinHeightTrees(6, s))
}
