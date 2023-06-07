package set

type Set[T comparable] interface {
	//向集合添加一个元素
	Add(e T) bool

	//添加所有元素
	AddAll(c ...T) bool

	//删除集合里的某个元素
	Delete(e T) bool

	//移除集合中指定的所有元素
	DeleteAll(c ...T) bool

	//移除集合中除指定元素外的所有元素
	RetainAll(c ...T) bool

	//判断集合里是否有某个元素
	Has(e T) bool

	//判断集合中是否包含有全部指定元素
	HasAll(c ...T) bool

	//获取set中不存在于另外一个set的元素。
	Difference(o Set[T]) Set[T]

	// IsEmpty 判断集合是否为空
	IsEmpty() bool

	//清空集合
	Clear()

	//返回集合所包含元素的数量
	Size(e T) int

	//返回一个包含集合所有元素的数组
	Values(e T) []*T
}

type HashSet[T comparable] struct {
	Set[T]

	data map[T]struct{}
}

// 创建新的hashSet
func NewHashSet[T comparable]() *HashSet[T] {
	set := new(HashSet[T])
	set.data = make(map[T]struct{})

	return set
}

// 通过切片创建新的hashSet
func NewHashSetFromSlice[T comparable](c []T) *HashSet[T] {
	set := new(HashSet[T])
	set.data = make(map[T]struct{}, len(c))

	set.AddAll(c...)

	return set
}

// 添加一个元素到集合中，返回是否添加成功
func (set *HashSet[T]) Add(e T) bool {
	_, isExist := set.data[e]
	if isExist {
		return false
	}

	set.data[e] = struct{}{}
	return true
}

// 添加一个元素到集合中，返回是否添加成功
func (set *HashSet[T]) AddAll(c ...T) bool {
	isChanged := false
	for _, e := range c {
		_, isExist := set.data[e]
		if !isExist {
			set.data[e] = struct{}{}

			isChanged = true
		}

	}
	return isChanged
}

// 删除集合中的一个元素，返回是否删除成功
func (set *HashSet[T]) Delete(e T) bool {
	_, isExist := set.data[e]
	if !isExist {
		return false
	}

	delete(set.data, e)
	return true
}

// 移除集合中指定的所有元素
func (set *HashSet[T]) DeleteAll(c ...T) bool {
	isChanged := false
	for _, e := range c {
		_, isExist := set.data[e]
		if isExist {
			delete(set.data, e)

			isChanged = true
		}

	}
	return isChanged

}

// 移除集合中除指定元素外的所有元素
func (set *HashSet[T]) RetainAll(c ...T) bool {
	isChanged := false
	newSet := NewHashSetFromSlice(c)

	for e := range set.data {
		if !newSet.Has(e) {
			delete(set.data, e)
			isChanged = true
		}
	}

	return isChanged
}

// 判断集合中是否包含某个元素
func (set *HashSet[T]) Has(e T) bool {
	_, isExist := set.data[e]

	return isExist
}

// 判断集合中是否包含有全部指定元素
func (set *HashSet[T]) HasAll(c ...T) bool {
	for _, v := range c {
		if !set.Has(v) {
			return false
		}
	}

	return true
}

func (set *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	var diff []T
	for e := range other.data {
		if !set.Has(e) {
			diff = append(diff, e)
		}
	}

	return NewHashSetFromSlice(diff)
}

// 清空集合
func (set *HashSet[T]) Clear() {
	set.data = make(map[T]struct{})
}

// 返回集合的大小
func (set *HashSet[T]) Size(e T) int {
	return len(set.data)
}

// 返回包含集合所有元素指针的数组
func (set *HashSet[T]) Values(e T) []*T {
	var values []*T
	for k, _ := range set.data {
		values = append(values, &k)
	}
	return values
}
