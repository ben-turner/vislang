package tokenizer

import (
	"bufio"
)

type simpleDefinition struct {
	Type    TokenType
	SubType TokenSubType
	Match   []rune
}

type functionMatch struct {
	Type    TokenType
	SubType TokenSubType
	Match   func(rune, *bufio.Reader) (bool, error)
}

var functionedDefinitions = []functionMatch{
	{
		Type:    ObjectReference,
		SubType: ObjectReferenceSimple,
		Match: func(rune, *bufio.Reader) (bool, error) {

		},
	},
	{
		Type:    ObjectReference,
		SubType: ObjectReferenceCreate,
		Match:   []rune{'^'},
	},
	{
		Type:    ObjectReference,
		SubType: ObjectReferenceDelete,
		Match:   []rune{'~'},
	},
	{
		Type:    ValueLiteral,
		SubType: ValueLiteralString,
		Match:   []rune{'"'},
	},
	{
		Type:    ValueLiteral,
		SubType: ValueLiteralInteger,
		Match:   []rune{'*'},
	},
}

var simpleDefinitions = []simpleDefinition{
	// Assignment Operators
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorAssign,
		Match:   []rune{'='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorAdd,
		Match:   []rune{'+', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorSubtract,
		Match:   []rune{'-', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorMultiply,
		Match:   []rune{'*', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorDivide,
		Match:   []rune{'/', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorPow,
		Match:   []rune{'*', '*', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorBitwiseAnd,
		Match:   []rune{'&', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorBitwiseOr,
		Match:   []rune{'|', '='},
	},
	{
		Type:    AssignmentOperator,
		SubType: AssignmentOperatorBitwiseXor,
		Match:   []rune{'^', '='},
	},
	// Combining Operators
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorEqual,
		Match:   []rune{'=', '='},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorNotequal,
		Match:   []rune{'!', '='},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorLessthan,
		Match:   []rune{'<'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorLessequal,
		Match:   []rune{'<', '='},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorGreaterthan,
		Match:   []rune{'>'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorGreaterequal,
		Match:   []rune{'>', '='},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorAnd,
		Match:   []rune{'&', '&'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorOr,
		Match:   []rune{'|', '|'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorXor,
		Match:   []rune{'^', '^'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorAdd,
		Match:   []rune{'+'},
	},
	{
		Type:    CombiningOperator,
		SubType: ControlOperatorSubtract,
		Match:   []rune{'-'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorMultiply,
		Match:   []rune{'*'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorDivide,
		Match:   []rune{'/'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorPow,
		Match:   []rune{'*', '*'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorBitwiseAnd,
		Match:   []rune{'&'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorBitwiseOr,
		Match:   []rune{'|'},
	},
	{
		Type:    CombiningOperator,
		SubType: CombiningOperatorBitwiseXor,
		Match:   []rune{'^'},
	},
	// Control Operators
	{
		Type:    ControlOperator,
		SubType: ControlOperatorIf,
		Match:   []rune{'i', 'f'},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorElse,
		Match:   []rune{'e', 'l', 's', 'e'},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorWhile,
		Match:   []rune{'w', 'h', 'i', 'l', 'e'},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorGroupstart,
		Match:   []rune{'('},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorGroupend,
		Match:   []rune{')'},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorBlockstart,
		Match:   []rune{'{'},
	},
	{
		Type:    ControlOperator,
		SubType: ControlOperatorBlockend,
		Match:   []rune{'}'},
	},
}
