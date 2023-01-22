// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(2), // id
			nil,      // ->
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // $
			nil,          // id
			nil,          // ->
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			nil,      // id
			shift(3), // ->
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(4), // id
			nil,      // ->
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // $, reduce: Grammar
			nil,       // id
			nil,       // ->
		},
	},
}
