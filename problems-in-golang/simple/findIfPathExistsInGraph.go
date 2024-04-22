package simple

import (
	"fmt"
)

// func someFunc(edges [][]int, curr, dest, n int) bool {
// 	if n < 0 {
// 		return false
// 	}
// 	if curr == dest {
// 		return true
// 	}
// 	for i := 0; i < len(edges); i++ {
// 		if edges[i][0] == curr {
// 			if someFunc(edges, edges[i][1], dest, n-1) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func validPath(n int, edges [][]int, source int, destination int) (ans bool) {
// 	if len(edges) < 1 {
// 		return true
// 	}
// 	for i := 0; i < len(edges); i++ {
// 		if edges[i][0] == source {
// 			// does edges[i][1] connect to destination?
// 			if someFunc(edges, edges[i][1], destination, n) || someFunc1(edges, edges[i][0], destination, n) {
// 				ans = true
// 			}
// 		}
// 	}
// 	return ans
// }

// func check(edges [][]int, vis *map[int]bool, found *bool, curr, target, n int) {
// 	if curr == target {
// 		*found = true
// 		return
// 	}
// 	if n < 0 {
// 		return
// 	}
// 	for i := range edges {
// 		if (*vis)[(edges[i][0]+1)*(edges[i][1]+1)] {
// 			return
// 		}
// 		if edges[i][0] == curr {
// 			(*vis)[(edges[i][0]+1)*(edges[i][1]+1)] = true
// 			check(edges, vis, found, edges[i][1], target, n-1)
// 		} else if edges[i][1] == curr {
// 			(*vis)[(edges[i][0]+1)*(edges[i][1]+1)] = true
// 			check(edges, vis, found, edges[i][0], target, n-1)
// 		}
// 		(*vis)[(edges[i][0]+1)*(edges[i][1]+1)] = false

// 	}
// }

//	func validPath(n int, edges [][]int, source int, destination int) (ans bool) {
//		if n <= 1 {
//			return true
//		}
//		vis := make(map[int]bool)
//		check(edges, &vis, &ans, source, destination, n)
//		return ans
//	}

type (
	Graph map[int]Conn
	Conn  map[int]bool
)

func createGraph(edges [][]int) (res Graph) {
	res = Graph{}
	for _, edge := range edges {
		src, dst := edge[0], edge[1]

		if conn, ok := res[src]; ok {
			conn[dst] = true
		} else {
			res[src] = Conn{dst: true}
		}

		if conn, ok := res[dst]; ok {
			conn[src] = true
		} else {
			res[dst] = Conn{src: true}
		}
	}
	return res
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if n <= 1 {
		return true
	}
	g := createGraph(edges)
	vis := map[int]bool{}
	q := []int{source}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]
			if c == destination {
				return true
			}
			if !vis[c] {
				for n := range g[c] {
					q = append(q, n)
				}
			}
			vis[c] = true
		}
	}
	return false
}

func TestValidPath() {
	s5 := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	fmt.Println(validPath(10, s5, 5, 9)) // True

	s4 := [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}
	fmt.Println(validPath(10, s4, 7, 5)) // True

	s3 := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	fmt.Println(validPath(10, s3, 5, 9)) // True

	s := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println(validPath(3, s, 0, 2)) // True

	s1 := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	fmt.Println(validPath(6, s1, 0, 5)) // False

	s2 := [][]int{}
	fmt.Println(validPath(1, s2, 0, 0)) // True
}
