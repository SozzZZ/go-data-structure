package stack

type StackBase[T any] interface {
	//栈顶增加一个元素
	Push(e T) bool

	//出栈
	Pop(e T) bool

	//清空栈
	Clear(e T)

	//返回栈顶元素
	Peek(e T) (T, bool)

	//返回栈所包含元素的数量
	Size(e T) int

	//判断栈是否为空
	IsEmpty() bool
}

type Stack[T any] struct {
	StackBase[T]

	data []T
}

// 创建新的stack
func NewStack[T any]() *Stack[T] {
	stack := new(Stack[T])
	stack.data = make([]T, 0)

	return stack
}

// 栈顶增加一个元素
func (stack *Stack[T]) Push(e T) bool {
	stack.data = append(stack.data, e)

	return true
}

// 出栈
func (stack *Stack[T]) Pop(e T) bool {
	if len(stack.data) >= 1 {
		stack.data = stack.data[:len(stack.data)-1]
		return true
	}

	return false
}

// 清空栈
func (stack *Stack[T]) Clear() {
	stack.data = make([]T, 0)
}

// 返回栈顶元素，已经是否有值
func (stack *Stack[T]) Peek() (T, bool) {
	if len(stack.data) >= 1 {
		return stack.data[len(stack.data)-1], true
	}

	var zero T
	return zero, false
}

// 返回栈所包含元素的数量
func (stack *Stack[T]) Size() int {

	return len(stack.data)
}

// 判断栈是否为空
func (stack *Stack[T]) IsEmpty() bool {

	return len(stack.data) == 0
}
