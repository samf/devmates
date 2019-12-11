package main

import (
	"fmt"
	"strings"
)

/*
 * Problem: 🧟 Zombies in Matrix
 * Given a 2D grid, each cell is either a zombie 1 or a human 0. Zombies can
 * turn adjacent (up / down / left / right) human beings into zombies every
 * hour. Find out how many hours does it take to infect all humans?
 *
 *  Input:
 *  [[0, 1, 1, 0, 1],
 *   [0, 1, 0, 1, 0],
 *   [0, 0, 0, 0, 1],
 *   [0, 1, 0, 0, 0]]
 *
 *  Output: 2
 *
 *  Why?
 *  At the end of the 1st hour, the status of the grid:
 *  [[1, 1, 1, 1, 1],
 *   [1, 1, 1, 1, 1],
 *   [0, 1, 0, 1, 1],
 *   [1, 1, 1, 0, 1]]
 *
 *  At the end of the 2nd hour, the status of the grid:
 *  [[1, 1, 1, 1, 1],
 *   [1, 1, 1, 1, 1],
 *   [1, 1, 1, 1, 1],
 *   [1, 1, 1, 1, 1]]
 */

type (
	zombieLand [][]int
	coord      [2]int
)

func main() {
	var cases []zombieLand = []zombieLand{
		zombieLand{
			[]int{0, 1, 1, 0, 1},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 0, 0, 0, 1},
			[]int{0, 1, 0, 0, 0},
		},
		zombieLand{},
		zombieLand{
			[]int{0, 0, 0, 0, 0, 1},
		},
		zombieLand{
			[]int{0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
		},
		zombieLand{
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, c := range cases {
		fmt.Printf("%v\n", c)
		result, err := solve(c)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Answer: %v\n", result)
	}
}

func solve(z zombieLand) (int, error) {
	var rc int

	for {
		mutations, err := z.mutation()
		if err != nil {
			return 0, err
		}
		if len(mutations) == 0 {
			break
		}
		rc++
		z.mutate(mutations)
	}

	return rc, nil
}

func (z zombieLand) mutation() ([]coord, error) {
	var (
		coords []coord
		seen   bool
	)

	for row, rowStuff := range z {
		for col, cell := range rowStuff {
			if cell == 1 {
				seen = true
				if victim := z.victim(row-1, col); victim != nil {
					coords = append(coords, *victim)
				}
				if victim := z.victim(row+1, col); victim != nil {
					coords = append(coords, *victim)
				}
				if victim := z.victim(row, col-1); victim != nil {
					coords = append(coords, *victim)
				}
				if victim := z.victim(row, col+1); victim != nil {
					coords = append(coords, *victim)
				}
			}
		}
	}

	if !seen {
		return coords, fmt.Errorf("no zombies anywhere")
	}

	return coords, nil
}

func (z zombieLand) mutate(c []coord) {
	for _, victim := range c {
		z[victim[0]][victim[1]] = 1
	}
}

func (z zombieLand) victim(row, col int) *coord {
	rows := len(z)
	if rows == 0 {
		return nil
	}
	cols := len(z[0])
	if cols == 0 {
		return nil
	}

	// if we're out of bounds return nothing
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return nil
	}

	// if our victim isn't a human, return nothing
	if z[row][col] != 0 {
		return nil
	}

	// victim!
	found := coord{row, col}
	return &found
}

func (z zombieLand) String() string {
	var (
		rows []string
	)

	for _, row := range z {
		rString := fmt.Sprintf("%v", row)
		rows = append(rows, rString)
	}

	return "[" + strings.Join(rows, ",\n ") + "]"
}
