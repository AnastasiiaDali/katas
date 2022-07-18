package postfix_evaluation

import (
	"fmt"
	"strconv"

	"merge_lists/stack_data_structure"
)

func EvaluatePostfix(expression string) int {
	stack := stack_data_structure.NewStack(len(expression))
	var result int
	for _, i := range expression {
		k := string(i)
		isOperator := isOperation(k)
		if !isOperator {
			stack.Push(k)
		} else {
			value1, _ := stack.Pop()
			value2, _ := stack.Pop()
			integer1, _ := strconv.Atoi(value1)
			integer2, _ := strconv.Atoi(value2)
			result = calculator(k, integer1, integer2)
			stack.Push(fmt.Sprint(result))
		}
	}
	return result
}

func isOperation(value string) bool {
	if value != "+" && value != "*" && value != "/" {
		return false
	}
	return true
}

func calculator(operator string, value1 int, value2 int) int {
	result := 0

	switch operator {
	case "+":
		result = value2 + value1
	case "*":
		result = value2 * value1
	case "/":
		result = value2 / value1
	}

	return result

}
