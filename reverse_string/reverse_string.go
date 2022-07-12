package reverse_string

import (
	"strings"

	"merge_lists/stack_data_structure"
)

func ReverseString(initialStr string) string {
	var reversedStr string
	var reversedArray []string

	//push them one by one and then pop
	//reverse me
	splitedStr := strings.Split(initialStr, "")

	stack := stack_data_structure.NewStack(len(splitedStr))
	for _, item := range splitedStr {
		stack.Push(item)
	}

	for i := 0; i < len(splitedStr); i++ {
		removedItem, _ := stack.Pop()
		reversedArray = append(reversedArray, removedItem)
	}

	reversedStr = strings.Join(reversedArray, "")

	return reversedStr
}
