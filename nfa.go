/*	Thompson's Construction
Author: Adrian McNulty
Based on instructional video "go-thompson-final"
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
*/

// Use Thompson's Construction to build an nfa from a regular expression written in postfix notation
// Using 3 special characters
package main

import (
	"fmt"
)

// Store states and arrows.
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

// keep track of the initial and accept states of your nfa
type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(pofix string) *nfa {

	// Create an array of pointers to nfas
	nfastack := []*nfa{}

	// Loop through the pofix expression a rune at a time
	for _, r := range pofix {
		// Do something to the stack depending on the character
		switch r {
		case '.':
			// pop the top 2 items off the stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// the first edge of the accept state of frag1 should point to the initial state of frag2
			frag1.accept.edge1 = frag2.initial
			// append a new pointer to an nfa struct that represents the new nfa (frag1 + frag2)
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// pop the top 2 items off the stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// create 2 new states and join them to the two fragments
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			// append a new pointer to an nfa struct that represents the new nfa (frag1 + frag2)
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':
			// pop the top item off the stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			// append a new pointer to an nfa struct that represents the new nfa (frag1 + frag2)
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	// Return the entire nfa which is the last item on the stack
	// Need to add error checking here
	return nfastack[0]
}

//	Create a  function that takes a regex in postfix notation and any string and return true or false
func pomatch(po string, s string) bool {
	// create a default value set to false
	ismatch := false

	return ismatch
}

func main() {
	fmt.Println(pomatch("ab.c*", "cccc"))
}
