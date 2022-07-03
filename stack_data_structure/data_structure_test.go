package stack_data_structure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/stack_data_structure"
)

func TestStackData(t *testing.T) {
	t.Run("NewStack should return empty StackOfItems", func(t *testing.T) {
		//arrange
		capacity := 10

		//act
		stack := stack_data_structure.NewStack(capacity)

		//assert
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 10, cap(stack.StackOfItems))
	})
	t.Run("testing Push", func(t *testing.T) {
		t.Run("Push should return a StackOfItems with an item", func(t *testing.T) {
			//arrange
			item := "1"
			capacity := 5

			stack := stack_data_structure.NewStack(capacity)

			//act
			stack.Push(item)

			//assert
			assert.False(t, stack.IsEmpty())
		})
		t.Run("Push should return false if StackOfItems is already full", func(t *testing.T) {
			//arrange
			item1, item2, item3 := "1", "2", "3"
			capacity := 2

			stack := stack_data_structure.NewStack(capacity)

			stack.Push(item1)
			stack.Push(item2)

			//act
			boolValue := stack.Push(item3)

			//assert
			assert.False(t, boolValue)
		})
	})

	t.Run("testing Peek", func(t *testing.T) {
		t.Run("Peek should return the last item from the stuck but not remove it", func(t *testing.T) {
			//arrange
			item1, item2 := "1", "2"
			capacity := 5

			stack := stack_data_structure.NewStack(capacity)

			stack.Push(item1)
			stack.Push(item2)

			//act
			peekedItem, _ := stack.Peek()

			//assert
			assert.Equal(t, "2", peekedItem)
			assert.False(t, stack.IsEmpty())
		})
		t.Run("Peek should return false if StackOfItems is empty", func(t *testing.T) {
			//arrange
			capacity := 5
			stack := stack_data_structure.NewStack(capacity)

			//act
			_, boolValue := stack.Peek()

			//assert
			assert.False(t, false, boolValue)
		})
	})
	t.Run("testing Pop", func(t *testing.T) {
		//arrange
		item1, item2 := "1", "2"
		capacity := 2

		stack := stack_data_structure.NewStack(capacity)

		stack.Push(item1)
		stack.Push(item2)

		//act
		poppedItem, _ := stack.Pop()
		peekedItem, _ := stack.Peek()

		//assert
		assert.Equal(t, "2", poppedItem)
		assert.Equal(t, "1", peekedItem)
	})

	t.Run("testing IsEmpty", func(t *testing.T) {
		t.Run("IsEmpty should return true when StackOfItems is empty", func(t *testing.T) {
			//arrange
			capacity := 5
			stack := stack_data_structure.NewStack(capacity)

			//act
			boolValue := stack.IsEmpty()

			//assert
			assert.True(t, boolValue)
		})
		t.Run("IsEmpty should return false when StackOfItems is NOT empty", func(t *testing.T) {
			//arrange
			item1 := "1"
			capacity := 3

			stack := stack_data_structure.NewStack(capacity)
			stack.Push(item1)

			//act
			boolValue := stack.IsEmpty()

			//assert
			assert.False(t, boolValue)
		})
	})
	t.Run("testing IsFull", func(t *testing.T) {
		t.Run("IsFull should return true when StackOfItems is full", func(t *testing.T) {
			//arrange
			item1, item2 := "1", "2"
			capacity := 2
			stack := stack_data_structure.NewStack(capacity)

			stack.Push(item1)
			stack.Push(item2)

			//act
			boolValue := stack.IsFull()

			//assert
			assert.True(t, boolValue)
		})
		t.Run("IsFull should return false when StackOfItems is NOT full", func(t *testing.T) {
			//arrange
			item1 := "1"
			capacity := 3

			stack := stack_data_structure.NewStack(capacity)
			stack.Push(item1)

			//act
			boolValue := stack.IsFull()

			//assert
			assert.False(t, boolValue)
		})
	})

}
