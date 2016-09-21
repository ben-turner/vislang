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

type Token struct {
	Type  int
	Value []rune
}

func Parse(r *bufio.Reader) (tokens []Token, linesRead int, err error) {
	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		var t Token
		switch c {
		case '\'':
			fallthrough
		case '"':
			t = Token{
				TypeStringLit,
				nil,
			}
			for d, _, err := r.ReadRune(); d != c; d, _, err = r.ReadRune() {
				if err != nil {
					return nil, linesRead, errors.New("Error parsing string: " + err.Error())
				}
				t.Value = append(t.Value, d)
			}
		case '<':
			fallthrough
		case '>':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == '=' {
				t = Token{
					TypeComparator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeComparator,
					[]rune{c},
				}
			}
		case '=':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == '=' {
				t = Token{
					TypeComparator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeAssignmentOperator,
					[]rune{c},
				}
			}
		case '!':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == '=' {
				t = Token{
					TypeComparator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeLogicalOperator,
					[]rune{c},
				}
			}
		case '&':
			fallthrough
		case '|':
			fallthrough
		case '^':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == c {
				t = Token{
					TypeLogicalOperator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeBitwiseOperator,
					[]rune{c},
				}
			}
		case '+':
			fallthrough
		case '-':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == c || d == '=' {
				t = Token{
					TypeAssignmentOperator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeArithmaticOperator,
					[]rune{c},
				}
			}
		case '*':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			switch d {
			case '*':
				t = Token{
					TypeArithmaticOperator,
					[]rune{c, d},
				}
			case '=':
				t = Token{
					TypeAssignmentOperator,
					[]rune{c, d},
				}
			default:
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeArithmaticOperator,
					[]rune{c},
				}
			}
		case '%':
			fallthrough
		case '/':
			d, _, err := r.ReadRune()
			if err != nil {
				return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
			}
			if d == '=' {
				t = Token{
					TypeAssignmentOperator,
					[]rune{c, d},
				}
			} else {
				if err := r.UnreadRune(); err != nil {
					return nil, linesRead, errors.New("Error parsing operator: " + err.Error())
				}
				t = Token{
					TypeArithmaticOperator,
					[]rune{c},
				}
			}
		case '\n':
			linesRead++
			fallthrough
		case ';':
			t = Token{
				TypeInstructionEnd,
				[]rune{c},
			}
		case '(':
			t = Token{
				TypeBlockStart,
				[]rune{c},
			}
		case ')':
			t = Token{
				TypeBlockEnd,
				[]rune{c},
			}
		case '#':
			for {
				c, _, err := r.ReadRune()
				if err != nil {
					return nil, linesRead, err
				}
				if c == '\n' {
					break
				}
			}
		case ' ':
			continue
		default:
			if unicode.IsDigit(c) {
				t = Token{
					TypeNumberLit,
					[]rune{c},
				}
				for c, _, err := r.ReadRune(); unicode.IsDigit(c) && err != io.EOF; c, _, err = r.ReadRune() {
					if err != nil {
						return nil, linesRead, err
					}
					t.Value = append(t.Value, c)
				}
			} else if unicode.IsLetter(c) {
				t = Token{
					TypeObjectReference,
					[]rune{c},
				}
				for c, _, err := r.ReadRune(); unicode.IsLetter(c) && err != io.EOF; c, _, err = r.ReadRune() {
					if err != nil {
						return nil, linesRead, err
					}
					t.Value = append(t.Value, c)
				}
			} else {
				return nil, linesRead, errors.New("Invalid identifier: " + string(c))
			}
			if err := r.UnreadRune(); err != nil {
				return nil, linesRead, err
			}
		}
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
