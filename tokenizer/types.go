package tokenizer

import (
	"fmt"
)

type TokenType uint8
type TokenSubType uint8

type Token struct {
	Type    TokenType
	SubType TokenSubType
	Value   []rune
}

type Definition struct {
	Type    TokenType
	SubType TokenSubType
	Start   string
	End     string
}

const (
	ObjectReference TokenType = iota
	ValueLiteral
	AssignmentOperator
	LogicalOperator
	CombiningOperator
	ControlOperator

	ObjectReferenceSimple TokenSubType = iota
	ObjectReferenceCreate
	ObjectReferenceDelete

	ValueLiteralString TokenSubType = iota
	ValueLiteralInteger
	ValueLiteralFloat
	ValueLiteralBool
	ValueLiteralNil

	AssignmentOperatorAssign TokenSubType = iota
	AssignmentOperatorAdd
	AssignmentOperatorSubtract
	AssignmentOperatorMultiply
	AssignmentOperatorDivide
	AssignmentOperatorPow
	AssignmentOperatorBitwiseAnd
	AssignmentOperatorBitwiseOr
	AssignmentOperatorBitwiseXor

	CombiningOperatorEqual TokenSubType = iota
	CombiningOperatorNotequal
	CombiningOperatorLessthan
	CombiningOperatorLessequal
	CombiningOperatorGreaterthan
	CombiningOperatorGreaterequal
	CombiningOperatorAnd
	CombiningOperatorOr
	CombiningOperatorXor
	CombiningOperatorAdd
	CombiningOperatorSubtract
	CombiningOperatorMultiply
	CombiningOperatorDivide
	CombiningOperatorPow
	CombiningOperatorBitwiseAnd
	CombiningOperatorBitwiseOr
	CombiningOperatorBitwiseXor

	ControlOperatorIf TokenSubType = iota
	ControlOperatorElse
	ControlOperatorWhile
	ControlOperatorGroupstart
	ControlOperatorGroupend
	ControlOperatorBlockstart
	ControlOperatorBlockend
)

func (t *Token) String() string {
	return fmt.Sprint("Token{%v, %v}", TokenSubType(t.Type, t.SubType), string(t.Value))
}

func TokenTypeString(t TokenType) string {
	switch t {
	case ObjectReference:
		return "ObjectReference"
	case ValueLiteral:
		return "ValueLiteral"
	case AssignmentOperator:
		return "AssignmentOperator"
	case CombiningOperator:
		return "CombiningOperator"
	case ControlOperator:
		return "ControlOperator"
	default:
		return "Invalid"
	}
}

func TokenSubTypeString(t TokenType, s TokenSubType) string {
	switch t {
	case ObjectReference:
		switch s {
		case ObjectReferenceSimple:
			return "ObjectReferenceSimple"
		case ObjectReferenceCreate:
			return "ObjectReferenceCreate"
		case ObjectReferenceDelete:
			return "ObjectReferenceDelete"
		}
	case ValueLiteral:
		switch s {
		case ValueLiteralString:
			return "ValueLiteralString"
		case ValueLiteralInteger:
			return "ValueLiteralInteger"
		case ValueLiteralFloat:
			return "ValueLiteralFloat"
		case ValueLiteralBool:
			return "ValueLiteralBool"
		case ValueLiteralNil:
			return "ValueLiteralNil"
		default:
			return "Invalid"
		}
	case AssignmentOperator:
		switch s {
		case AssignmentOperatorAssign:
			return "AssignmentOperatorAssign"
		case AssignmentOperatorAdd:
			return "AssignmentOperatorAdd"
		case AssignmentOperatorSubtract:
			return "AssignmentOperatorSubtract"
		case AssignmentOperatorMultiply:
			return "AssignmentOperatorMultiply"
		case AssignmentOperatorDivide:
			return "AssignmentOperatorDivide"
		case AssignmentOperatorPow:
			return "AssignmentOperatorPow"
		case AssignmentOperatorBitwiseXor:
			return "AssignmentOperatorBitwiseXor"
		case AssignmentOperatorBitwiseAnd:
			return "AssignmentOperatorBitwiseAnd"
		case AssignmentOperatorBitwiseOr:
			return "AssignmentOperatorBitwiseOr"
		default:
			return "Invalid"
		}
	case CombiningOperator:
		switch s {
		case CombiningOperatorEqual:
			return "CombiningOperatorEqual"
		case CombiningOperatorNotequal:
			return "CombiningOperatorNotequal"
		case CombiningOperatorLessthan:
			return "CombiningOperatorLessthan"
		case CombiningOperatorLessequal:
			return "CombiningOperatorLessequal"
		case CombiningOperatorGreaterthan:
			return "CombiningOperatorGreaterthan"
		case CombiningOperatorGreaterequal:
			return "CombiningOperatorGreaterequal"
		case CombiningOperatorAnd:
			return "CombiningOperatorAnd"
		case CombiningOperatorOr:
			return "CombiningOperatorOr"
		case CombiningOperatorAdd:
			return "CombiningOperatorAdd"
		case CombiningOperatorSubtract:
			return "CombiningOperatorSubtract"
		case CombiningOperatorMultiply:
			return "CombiningOperatorMultiply"
		case CombiningOperatorDivide:
			return "CombiningOperatorDivide"
		case CombiningOperatorPow:
			return "CombiningOperatorPow"
		case CombiningOperatorBitwiseAnd:
			return "CombiningOperatorBitwiseAnd"
		case CombiningOperatorBitwiseOr:
			return "CombiningOperatorBitwiseOr"
		case CombiningOperatorBitwiseXor:
			return "CombiningOperatorBitwiseXor"
		default:
			return "Invalid"
		}
	case ControlOperator:
		switch s {
		case ControlOperatorIf:
			return "ControlOperatorIf"
		case ControlOperatorElse:
			return "ControlOperatorElse"
		case ControlOperatorWhile:
			return "ControlOperatorWhile"
		case ControlOperatorGroupstart:
			return "ControlOperatorGroupstart"
		case ControlOperatorGroupend:
			return "ControlOperatorGroupend"
		case ControlOperatorBlockstart:
			return "ControlOperatorBlockstart"
		case ControlOperatorBlockend:
			return "ControlOperatorBlockend"
		default:
			return "Invalid"
		}
	default:
		return "Invalid"
	}
}
