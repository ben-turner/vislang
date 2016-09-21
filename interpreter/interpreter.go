package interpreter

import (
	"encoding/binary"
	"fmt"
	. "github.com/ben-turner/vislang/tokenizer"
	"math"
	"os"
	"strconv"
)

type ValueType uint

const (
	Int ValueType = iota
	String
	Empty
)

func Run(tokens []Token) error {
	var err error
	for next := 0; next < len(tokens) && next >= 0; _, _, next, err = exec(tokens, next) {
		if err != nil {
			return err
		}
	}
	return nil
}

func exec(tokens []Token, i int) (ValueType, []byte, int, error) {
	if i > len(tokens) || i < 0 {
		return Empty, nil, -1, fmt.Errorf("Ran out of code!")
	}
	t := tokens[i]
	switch t.Type {
	case TypeObjectReference:
		if string(t.Value) == "log" {
			valType, out, next, err := exec(tokens, i+1)
			if err != nil {
				return Empty, nil, -1, err
			}
			switch valType {
			case String:
				fmt.Fprintf(os.Stderr, "%v\n", string(out))
			case Int:
				num, _ := binary.Varint(out)
				fmt.Fprintf(os.Stderr, "%v\n", num)
			}

			return Empty, nil, next, nil
		} else {
			return Empty, nil, -1, fmt.Errorf("Reference %v not recognized.", t.Value)
		}
	case TypeArithmaticOperator:
		valtype, a, next, err := exec(tokens, i+1)
		if err != nil {
			return Empty, nil, -1, err
		}
		if valtype != Int {
			return Empty, nil, -1, fmt.Errorf("Unsupported type")
		}
		aInt, _ := binary.Varint(a)
		valtype, b, next, err := exec(tokens, next)
		if err != nil {
			return Empty, nil, -1, err
		}
		if valtype != Int {
			return Empty, nil, -1, fmt.Errorf("Unsupported type")
		}
		bInt, _ := binary.Varint(b)
		var out int64
		switch string(t.Value) {
		case "+":
			out = aInt + bInt
		case "-":
			out = aInt - bInt
		case "*":
			out = aInt * bInt
		case "/":
			out = aInt / bInt
		}
		buf := make([]byte, 8)
		c := binary.PutVarint(buf, out)
		return Int, buf[:c], next, nil
	case TypeBlockStart:
		valtype, curBytes, next, err := exec(tokens, i+1)
		if err != nil {
			return Empty, nil, -1, err
		}
		if valtype != Int {
			return Empty, nil, -1, fmt.Errorf("Unsupported type")
		}
		currentVal, _ := binary.Varint(curBytes)
		for {
			if next >= len(tokens) {
				return Empty, nil, -1, fmt.Errorf("Ran out of code!")
			}
			nextToken := tokens[next]
			if nextToken.Type == TypeBlockEnd {
				return Int, curBytes, next + 1, nil
			}
			if nextToken.Type == TypeArithmaticOperator {
				var nextBytes []byte
				valtype, nextBytes, next, err = exec(tokens, next+1)
				if err != nil {
					return Empty, nil, -1, err
				}
				if valtype != Int {
					return Empty, nil, -1, fmt.Errorf("Unsupported type")
				}
				nextVal, _ := binary.Varint(nextBytes)
				switch string(nextToken.Value) {
				case "+":
					currentVal += nextVal
				case "-":
					currentVal -= nextVal
				case "/":
					currentVal /= nextVal
				case "*":
					currentVal *= nextVal
				case "**":
					currentVal = int64(math.Pow(float64(currentVal), float64(nextVal)))
				}
			}
		}
		buf := make([]byte, 8)
		c := binary.PutVarint(buf, currentVal)
		return Int, buf[:c], next, nil
	case TypeStringLit:
		return String, []byte(string(t.Value)), i + 1, nil
	case TypeNumberLit:
		num, _ := strconv.Atoi(string(t.Value))
		buf := make([]byte, 8)
		c := binary.PutVarint(buf, int64(num))
		return Int, buf[:c], i + 1, nil
	}
	return Empty, nil, i + 1, nil
}
