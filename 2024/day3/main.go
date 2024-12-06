package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	mulOp struct {
		X int
		Y int
	}
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalln("a path must be provided")
	}
	path := os.Args[1]
	fmt.Printf("reading path: %s\n", path)
	ops, err := parseMultiplicationOperations(path)
	if err != nil {
		log.Fatal(err)
	}
	actualSum := getSumFromMulOps(ops)
	fmt.Printf("Sum from ops: %d\n", actualSum)
}

func parseMultiplicationOperations(filePath string) ([]mulOp, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()
	sb := &strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if _, err := sb.WriteString(line); err != nil {
			return nil, fmt.Errorf("(*strings.Builder).WriteString: %w", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("(*bufio.Scanner).Err: %w", err)
	}
	return parseMemoryForMulOps(sb.String()), nil
}

type State string

const (
	StateUnknown      State = "unknown"
	StateLetterM      State = "M"
	StateLetterU      State = "U"
	StateLetterL      State = "L"
	StateOpenBracket  State = "("
	StateXDigit       State = "XDigit"
	StateYDigit       State = "YDigit"
	StateComma        State = "COMMA"
	StateCloseBracket State = ")"
)

func parseMemoryForMulOps(source string) []mulOp {
	prevState := StateUnknown
	ops := make([]mulOp, 0)
	var xVal, yVal int
	numDigitsSeen := 0
	for _, r := range source {
		switch prevState {
		case StateLetterM:
			if r != 'u' {
				prevState = StateUnknown
				break
			}
			prevState = StateLetterU
		case StateLetterU:
			if r != 'l' {
				prevState = StateUnknown
				break
			}
			prevState = StateLetterL
		case StateLetterL:
			if r != '(' {
				prevState = StateUnknown
				break
			}
			prevState = StateOpenBracket
		case StateOpenBracket:
			if r < '0' || r > '9' {
				prevState = StateUnknown
				break
			}
			prevState = StateXDigit
			numDigitsSeen = 1
			// Reset xVal and yVal
			xVal, yVal = 0, 0
			xVal = calculateValueBasedOnDigitNumber(xVal, r, numDigitsSeen)
		case StateXDigit:
			// If we previously saw a digit, we can only see another digit, or a comma.
			if r != ',' && (r < '0' || r > '9') {
				prevState = StateUnknown
				break
			}
			if r == ',' {
				prevState = StateComma
				break
			}
			// Another digit.
			numDigitsSeen++
			if numDigitsSeen > 3 {
				// Reset the value for X.
				xVal = 0
				prevState = StateUnknown
				break
			}
			xVal = calculateValueBasedOnDigitNumber(xVal, r, numDigitsSeen)
		case StateYDigit:
			// If we previously saw a digit, we can only see another digit, or a closing bracket.
			if r != ')' && (r < '0' || r > '9') {
				prevState = StateUnknown
				break
			}
			if r == ')' {
				// If we see a closing bracket here, we succesfully found a mul(X,Y).
				ops = append(ops, mulOp{X: xVal, Y: yVal})
				prevState = StateCloseBracket
				numDigitsSeen = 0
				xVal = 0
				yVal = 0
				break
			}
			// Another digit.
			numDigitsSeen++
			if numDigitsSeen > 3 {
				// Reset the value for Y.
				yVal = 0
				prevState = StateUnknown
				break
			}
			yVal = calculateValueBasedOnDigitNumber(yVal, r, numDigitsSeen)
		case StateComma:
			// Next value must be a digit.
			if r < '0' || r > '9' {
				prevState = StateUnknown
				break
			}
			prevState = StateYDigit
			numDigitsSeen = 1
			yVal = calculateValueBasedOnDigitNumber(yVal, r, numDigitsSeen)
		case StateCloseBracket:
			// We just reset to the unknown state unless the rune is an 'm'.
			if r == 'm' {
				prevState = StateLetterM
				break
			}
			prevState = StateUnknown
		case StateUnknown:
			if r == 'm' {
				prevState = StateLetterM
			}
		default:
			panic("unexpected state")
		}
	}
	return ops
}

func calculateValueBasedOnDigitNumber(currValue int, r rune, digitNum int) int {
	val := int(r - '0')
	if digitNum < 1 || digitNum > 3 {

	}
	if digitNum == 1 && currValue != 0 {
		panic("first digit should have currValue = 0")
	}
	valToReturn := val
	switch digitNum {
	case 1:
		break
	case 2:
		fallthrough
	case 3:
		return (currValue * 10) + val
	default:
		panic("digitNum was not between 0 and 3")
	}
	return valToReturn
}

func getSumFromMulOps(input []mulOp) int {
	sum := 0
	for _, op := range input {
		sum += op.X * op.Y
	}
	return sum
}
