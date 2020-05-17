package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"calculator/queues"
	"calculator/slicestacks"
)

var (
	regexUpperCase = regexp.MustCompile("[A-Z]+")
)

const (
	openParanthesis  = '('
	closeParanthesis = ')'
)

type Calculator struct {
	mapVariables map[string]string
}

func (c *Calculator) Evaluate(expr string) (string, error) {
	if c.mapVariables == nil {
		c.mapVariables = make(map[string]string)
	}

	if ok := strings.Contains(expr, "="); ok {
		dataArr := strings.Split(expr, "=")

		key := strings.TrimSpace(dataArr[0])
		value := strings.TrimSpace(dataArr[1])

		if data, ok := c.mapVariables[key]; ok {
			newExpr := strings.Replace(expr, key, data, 1)
			return c.evaluate(newExpr)
		}

		// find key contains in string
		listCharacters := regexUpperCase.FindStringSubmatch(value)

		// find key and replace value of key
		if len(listCharacters) > 0 {
			for _, key := range listCharacters {
				// check into map by key
				if data, ok := c.mapVariables[key]; ok {
					newExpr := strings.Replace(value, key, data, -1)
					c.mapVariables[key] = newExpr

					return c.evaluate(newExpr)
				}
			}
		}

		c.mapVariables[key] = fmt.Sprintf("(%s)", value)
		return c.evaluate(fmt.Sprintf("(%s)", value))
	} else {
		listCharacters := regexUpperCase.FindStringSubmatch(expr)
		// find key and replace value of key
		if len(listCharacters) > 0 {
			for _, value := range listCharacters {
				if data, ok := c.mapVariables[value]; ok {
					newExpr := strings.Replace(expr, value, data, -1)

					return c.evaluate(newExpr)
				}
			}
		}
	}

	return c.evaluate(expr)
}

func (c *Calculator) evaluate(expr string) (string, error) {
	stack := slicestacks.New()
	queue := queues.New()

	for _, char := range []rune(expr) {
		// ignore space characters
		if char == 32 {
			continue
		}
		// only number to push into queue
		if 48 <= char && char <= 57 {
			queue.Enqueue(char)
		} else {
			// check if close characters
			if c.isCloseParantheses(char) {
				for {
					if stack.IsEmpty() {
						break
					}
					value, err := stack.Pop()
					if err != nil {
						return "", err
					}
					// pop stack until open characters
					if c.isOpenParantheses(value.(rune)) {
						break
					}
					queue.Enqueue(value.(rune))
				}
			} else {
				// set priovity
				if stack.Size() == 1 && !c.isOpenParantheses(char) {
					value, err := stack.Pop()
					if err != nil {
						return "", err
					}
					runeString := value.(rune)

					if string(runeString) == "*" || string(runeString) == "/" && string(char) == "+" || string(char) == "-" {
						queue.Enqueue(runeString)
						stack.Push(char)
					} else {
						stack.Push(value, char)
					}
				} else {
					// push into stack
					stack.Push(char)
				}
			}
		}
	}

	// pop all value from stack to queue
	if stack.Size() > 0 {
		for {
			if stack.IsEmpty() {
				break
			}
			value, err := stack.Pop()
			if err != nil {
				return "", err
			}
			queue.Enqueue(value.(rune))
		}
	}
	// clear stack
	stack.Clear()

	for {
		if queue.IsEmpty() {
			break
		}
		value, err := queue.Dequeue()
		if err != nil {
			continue
		}

		char := value.(rune)

		if 48 <= char && char <= 57 {
			stack.Push(string(char))
		} else {
			for {
				if stack.IsEmpty() {
					break
				}
				// if contant 1 value
				if stack.Size() == 1 {
					valueStackOne, _ := stack.Pop()
					valueOne, _ := strconv.Atoi(valueStackOne.(string))
					switch string(char) {
					case "+":
						return fmt.Sprintf("%d", valueOne), nil
					case "-":
						return fmt.Sprintf("-%d", valueOne), nil
					}
				}
				valueStackOne, err := stack.Pop()
				if err != nil {
					return "", err
				}
				valueStackTwo, err := stack.Pop()
				if err != nil {
					return "", err
				}

				valueOne, err := strconv.Atoi(valueStackOne.(string))
				if err != nil {
					return "", err
				}
				valueTwo, err := strconv.Atoi(valueStackTwo.(string))
				if err != nil {
					return "", err
				}

				switch string(char) {
				case "*":
					stack.Push(fmt.Sprintf("%d", valueOne*valueTwo))
				case "-":
					stack.Push(fmt.Sprintf("%d", valueOne-valueTwo))
				case "+":
					stack.Push(fmt.Sprintf("%d", valueOne+valueTwo))
				case "/":
					stack.Push(fmt.Sprintf("%0f", float64(valueOne)/float64(valueTwo)))
				}
				break
			}
		}
	}
	result, err := stack.Pop()
	if err != nil {
		return "", err
	}

	return result.(string), nil
}

func (c *Calculator) isOpenParantheses(character rune) bool {
	return character == openParanthesis
}

func (c *Calculator) isCloseParantheses(character rune) bool {
	return character == closeParanthesis
}
