/*	Shunting Arm Algorithm
Author: Adrian McNulty
Based on instructional video "Shunting yard algorithm in Go"
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e?channelId=f9970e30-b336-4145-8af3-a2bbe2938f5e&channelName=Graph%20Theory
*/

package main

import (
	"fmt"
)

/*	Create a function to convert expressions from infix notation to postfix notation. */
func intopost(infix string) string {
	//	Create a map of the special characters and assign values to keep track of the precedence of the characters(i.e. * has a higher precedence than . which has precedence over |).
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//	Initialise an empty array of runes
	pofix := []rune{}
	//	Create a stack to temporarily hold operators as they are read from the string
	s := []rune{}
	//	Loop over the infix string and convert to postfix
	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		//	If character is not special or a bracket, append it to the end of the output
		default:
			pofix = append(pofix, r)
		}
	}
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}
	//	Cast pofix to a string
	return string(pofix)
}

func main() {
	//	Test program is working
	// Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))
	// Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))
	// Answer: abd|.c*.
	fmt.Println("Infix:   ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))
	// Answer: abb.+.c.
	fmt.Println("Infix:   ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
}
