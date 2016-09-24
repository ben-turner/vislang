package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const (
	TypeStringLit int = iota
	TypeNumberLit
	TypeComparator
	TypeAssignmentOperator
	TypeLogicalOperator
	TypeBitwiseOperator
	TypeArithmaticOperator
	TypeObjectReference
	TypeInstructionEnd
	TypeBlockStart
	TypeBlockEnd
)

func Parse(r *bufio.Reader) (tokens []Token, linesRead int, err error) {
	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		var t Token

		tokens = append(tokens, t)
		fmt.Fprintf(os.Stderr, "%v: %v\n", t.TypeString(), strings.Replace(string(t.Value), "\n", "\\n", -1))
	}
	return
}

func (t Token) TypeString() string {
	switch t.Type {
	case TypeStringLit:
		return "StringLit"
	case TypeNumberLit:
		return "NumberLit"
	case TypeComparator:
		return "Comparator"
	case TypeAssignmentOperator:
		return "AssignmentOperator"
	case TypeLogicalOperator:
		return "LogicalOperator"
	case TypeBitwiseOperator:
		return "BitwiseOperator"
	case TypeArithmaticOperator:
		return "ArithmaticOperator"
	case TypeObjectReference:
		return "ObjectReference"
	case TypeInstructionEnd:
		return "InstructionEnd"
	case TypeBlockStart:
		return "BlockStart"
	case TypeBlockEnd:
		return "BlockEnd"
	}
	return "Invalid"
}
