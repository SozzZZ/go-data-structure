package queue

type QueueBase[T any] interface {
	//队列增加一个元素
	Enqueue(e T) bool

	//出队列
	Dequeue(e T) (T, bool)

	//清空队列
	Clear(e T)

	//返回队列首位元素
	Peek(e T) (T, bool)

	//返回包含元素的数量
	Size(e T) int

	//判断是否为空
	IsEmpty() bool
}

type Queue[T any] struct {
	QueueBase[T]

	data []T
}

// 创建新的队列
func NewQueue[T any]() *Queue[T] {
	queue := new(Queue[T])
	queue.data = make([]T, 0)

	return queue
}
