package simple

type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{reqs: []int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.reqs = append(rc.reqs, t)

	for rc.reqs[0] < t-3000 {
		rc.reqs = rc.reqs[1:]
	}

	return len(rc.reqs)
}
