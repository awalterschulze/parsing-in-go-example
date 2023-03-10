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
			shift(3), // id
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
			nil,       // INVALID
			reduce(1), // $, reduce: Grammar
			nil,       // id
			nil,       // ->
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			nil,      // id
			shift(4), // ->
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(5), // id
			nil,      // ->
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // $, reduce: EdgeStmt
			nil,       // id
			nil,       // ->
		},
	},
}
