// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 45: // ['-','-']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 99: // ['a','c']
			return 3
		case r == 100: // ['d','d']
			return 4
		case 101 <= r && r <= 122: // ['e','z']
			return 3
		case r == 123: // ['{','{']
			return 5
		case r == 125: // ['}','}']
			return 6
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 62: // ['>','>']
			return 7
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 122: // ['a','z']
			return 3
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 104: // ['a','h']
			return 3
		case r == 105: // ['i','i']
			return 8
		case 106 <= r && r <= 122: // ['j','z']
			return 3
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 102: // ['a','f']
			return 3
		case r == 103: // ['g','g']
			return 9
		case 104 <= r && r <= 122: // ['h','z']
			return 3
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 113: // ['a','q']
			return 3
		case r == 114: // ['r','r']
			return 10
		case 115 <= r && r <= 122: // ['s','z']
			return 3
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case r == 97: // ['a','a']
			return 11
		case 98 <= r && r <= 122: // ['b','z']
			return 3
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 111: // ['a','o']
			return 3
		case r == 112: // ['p','p']
			return 12
		case 113 <= r && r <= 122: // ['q','z']
			return 3
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 103: // ['a','g']
			return 3
		case r == 104: // ['h','h']
			return 13
		case 105 <= r && r <= 122: // ['i','z']
			return 3
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case 97 <= r && r <= 122: // ['a','z']
			return 3
		}
		return NoState
	},
}
