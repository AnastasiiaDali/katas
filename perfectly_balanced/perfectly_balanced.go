package perfectly_balanced

import (
	"strings"

	"merge_lists/stack_data_structure"
)

func PerfectlyBalanced(expression string) bool {
	expressions := strings.Split(expression, "")

	stack := stack_data_structure.NewStack(len(expression)/2 + 1)

	for _, expr := range expressions {
		value, oppositeBrace := isOpeningBrace(expr)
		if value {
			stack.Push(expr)
		} else {
			poppedItem, ableToPop := stack.Pop()
			if poppedItem != oppositeBrace || !ableToPop {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func isOpeningBrace(brace string) (bool, string) {
	switch brace {
	case "{":
		return true, "}"
	case "(":
		return true, ")"
	case "[":
		return true, "]"
	case "}":
		return false, "{"
	case ")":
		return false, "("
	case "]":
		return false, "["
	default:
		return false, ""
	}
}

//func isClosingBrace(brace string) (bool, string) {
//	switch brace {
//	case "}":
//		return true, "{"
//	case ")":
//		return true, "("
//	case "]":
//		return true, "["
//	default:
//		return false, ""
//	}
//}

//{()}[)
//{}[)
//[)

//{[{}{}]}[()]
//{[{}]}[()]
//{[]}[()]
//{}[()]
//[()]
//[]
//
