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

type die struct {
	faces string
	which int
}

func main() {
	values := []struct {
		word  string
		faces []string
	}{
		{
			word: "hello",
			faces: []string{
				"alcdef",
				"abcdef",
				"abchef",
				"abcdof",
				"abclef",
			},
		},
		{
			word: "hello",
			faces: []string{
				"abcdef",
				"abcdef",
				"abcdef",
				"abcdef",
				"abcdef",
			},
		},
		{
			word: "aaaa",
			faces: []string{
				"aaaaaa",
				"bbbbbb",
				"abcdef",
				"abcdef",
			},
		},
	}

	for _, value := range values {
		fmt.Printf("word: %q\ndice: %+v\n", value.word, value.faces)
		fmt.Println(pretty(attempt(value.word, strings2dice(value.faces))))
		fmt.Println()
	}
}

func attempt(word string, dice []die) *[][2]int {
	if len(word) == 0 {
		// special case: success
		return &[][2]int{}
	}

	r := []rune(word)[0]
	for i, d := range dice {
		where := strings.IndexRune(d.faces, r)
		if where == -1 {
			continue
		}

		resp := attempt(word[1:], remove(i, dice))
		if resp == nil {
			continue
		}

		me := [][2]int{
			[2]int{d.which, where},
		}

		all := append(me, *resp...)
		return &all
	}

	return nil
}

func remove(i int, dice []die) []die {
	// copy the input lest we corrupt it
	res := make([]die, len(dice))
	copy(res, dice)

	res[i] = res[len(dice)-1]
	return res[:len(dice)-1]
}

func pretty(res *[][2]int) string {
	if res == nil {
		return "Output: false"
	}

	var pretties []string
	for _, die := range *res {
		pretties = append(pretties, fmt.Sprintf("dice[%v][%v]", die[0], die[1]))
	}

	return fmt.Sprintf("Output: true\nWhy? %v", strings.Join(pretties, " + "))
}

func strings2dice(faces []string) []die {
	var dice []die

	for i, s := range faces {
		d := die{
			which: i,
			faces: s,
		}

		dice = append(dice, d)
	}

	return dice
}
