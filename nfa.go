/*	Thompson's Construction
Author: Adrian McNulty
Based on instructional video "go-thompson-final"
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
	Regex match function
Based on instructional video "Regex match function"
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d?channelId=f9970e30-b336-4145-8af3-a2bbe2938f5e&channelName=Graph%20Theory
*/

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
		case '+':
			// pop the top item off the stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// create a new accept state variable
			accept := state{}
			// make edge1 initial state of fragment and edge2 point at the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial
			// push the new fragment to the stack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

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

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

//	Create a  function that takes a regex in postfix notation and any string and return true or false
func pomatch(po string, s string) bool {
	// create a default value set to false
	ismatch := false
	// Create a nfa from the regular expression
	ponfa := poregtonfa(po)
	// Create an array to keep track of the current states we are in in the nfa
	current := []*state{}
	// Create an array to keep track of the states we can move to
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Loop through string s
	for _, r := range s {
		// Loop through the current array
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// Replace the old states with new ones(i.e value of next becomes value of current and then next becomes empty array )
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}
	return ismatch
}

func main() {
	fmt.Println(pomatch("ab.c*|", "cccc"))
}
