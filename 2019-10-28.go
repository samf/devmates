package main

/*
 *Input:
 *  word = "hello"
 *  dice = [[a, l, c, d, e, f], [a, b, c, d, e, f],
 *          [a, b, c, h, e, f], [a, b, c, d, o, f],
 *          [a, b, c, l, e, f]]
 *  Output: true
 *  Why? dice[2][3] + dice[1][4] + dice[0][1] + dice[4][3] + dice[3][4]
 *
 *  Input:
 *  word = "hello"
 *  dice = [[a, b, c, d, e, f], [a, b, c, d, e, f],
 *          [a, b, c, d, e, f], [a, b, c, d, e, f],
 *          [a, b, c, d, e, f]]
 *  Output: false
 *
 *  Input:
 *  word = "aaaa"
 *  dice = [[a, a, a, a, a, a], [b, b, b, b, b, b],
 *          [a, b, c, d, e, f], [a, b, c, d, e, f]]
 *  Output: false
 */

import (
	"fmt"
	"strings"
)

func main() {
	values := []struct {
		word string
		dice []string
	}{
		{
			word: "hello",
			dice: []string{
				"alcdef",
				"abcdef",
				"abchef",
				"abcdof",
				"abclef",
			},
		},
		{
			word: "hello",
			dice: []string{
				"abcdef",
				"abcdef",
				"abcdef",
				"abcdef",
				"abcdef",
			},
		},
		{
			word: "aaaa",
			dice: []string{
				"aaaaaa",
				"bbbbbb",
				"abcdef",
				"abcdef",
			},
		},
	}

	for _, value := range values {
		fmt.Printf("word: %q\ndice: %+v\n", value.word, value.dice)
		fmt.Println(pretty(attempt(value.word, value.dice)))
	}
}

func attempt(word string, dice []string) *[][]int {
	if len(word) == 0 {
		// special case: success
		return &[][]int{}
	}

	r := []rune(word)[0]
	for i, die := range dice {
		where := strings.IndexRune(die, r)
		if where == -1 {
			continue
		}

		remaining := remove(i, dice)
		resp := attempt(word[1:], remaining)

		if resp == nil {
			continue
		}

		me := [][]int{
			[]int{i, where},
		}
		res := *resp

		me = append(me, res...)
		return &me
	}

	return nil
}

func pretty(res *[][]int) string {
	if res == nil {
		return "Output: false"
	}

	answer := *res

	var pretties []string
	for _, die := range answer {
		if len(die) != 2 {
			panic("wtf")
		}

		pretties = append(pretties, fmt.Sprintf("dice[%v][%v] ", die[0], die[1]))
	}

	return fmt.Sprintf("Output: true\nWhy: %v", strings.Join(pretties, ", "))
}

func remove(which int, dice []string) []string {
	switch which {
	case 0:
		return dice[1:]
	case len(dice) - 1:
		return dice[:len(dice)-1]
	}

	return append(dice[:which], dice[which+1:]...)
}
