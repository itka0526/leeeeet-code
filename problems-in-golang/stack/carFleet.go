package stack

import (
	"fmt"
	"sort"
)

// type car struct {
// 	reachedTarget bool
// 	speed         int
// 	pos           int
// }

// func (c *car) updatePos() {
// 	c.pos += c.speed
// }

// func carFleetSolution(target int, position []int, speed []int) int {
// 	cars := make([]car, len(position))

// 	for i := range position {
// 		cars[i] = car{speed: speed[i], pos: position[i], reachedTarget: false}
// 	}

// 	if len(cars) == 1 {
// 		return 1
// 	}

// 	allTheCarsReachedTarget := 0

// 	for allTheCarsReachedTarget < len(cars) {
// 		for i := range cars {
// 			if cars[i].reachedTarget {
// 				continue
// 			}

// 			cars[i].updatePos()

// 			if cars[i].pos >= target {
// 				cars[i].reachedTarget = true
// 				allTheCarsReachedTarget++
// 			}
// 		}

// 		// check if the cars have met each other
// 		for i := 0; i < len(cars); i++ {
// 			for j := 0; j < len(cars); j++ {
// 				if cars[i] == cars[j] {
// 					continue
// 				}

// 				if cars[i].pos == cars[j].pos && !cars[i].reachedTarget && !cars[j].reachedTarget {
// 					// now we must make them a fleet
// 					// if cars become a fleet we take the minimum speed of the fleet
// 					slowest := slowestCar(cars[i], cars[j])

// 					cars[i].speed = slowest
// 					cars[j].speed = slowest

// 					fmt.Println("BECAME A FLEET!")
// 				}
// 			}
// 		}
// 	}

// 	for _, c := range cars {
// 		fmt.Printf("Speed: %d, Position: %d, Reached: %t \n", c.speed, c.pos, c.reachedTarget)
// 	}
// 	panic("not done")
// }

// func slowestCar(a, b car) int {
// 	if a.speed < b.speed {
// 		return a.speed
// 	}
// 	return b.speed
// }

type car struct {
	p int
	s int
}

func carFleetSolution(target int, position []int, speed []int) int {
	fleets := []float32{}
	cars := make([]car, len(position))

	for i := range position {
		cars[i] = car{p: position[i], s: speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].p < cars[j].p
	})

	for i := len(cars) - 1; i >= 0; i-- {
		t := float32(target-cars[i].p) / float32(cars[i].s)
		fleets = append(fleets, t)

		if len(fleets) >= 2 && fleets[len(fleets)-1] <= fleets[len(fleets)-2] {
			fleets = fleets[:len(fleets)-1]
		}
	}

	return len(fleets)
}

func TestCarFleetSolution() {
	fmt.Println(carFleetSolution(12, []int{3, 5, 7}, []int{3, 2, 1}))
	// fmt.Println(carFleetSolution(12, []int{0, 2, 4}, []int{4, 2, 1}))
}
