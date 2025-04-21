package structure

import (
	"golang.org/x/exp/constraints"
	"slices"
	"sort"
	"wh.com/Go/constraints"
	//"wh.com/Go/constraints"
	//"wh.com/Go/constraints"
	//"wh.com/Go/constraints"
)

func RemoveDuplicates[T constraints.Ordered](items []T) []T {
	// 如果有0或1个元素，则返回切片本身。
	if len(items) < 2 {
		return items
	}

	// 对切片进行升序排序
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})

	// 使用唯一指针来跟踪不重复的元素
	uniqPointer := 0
	for i := 1; i < len(items); i++ {
		if items[uniqPointer] != items[i] { // 比较当前元素和唯一指针指向的元素
			uniqPointer++
			items[uniqPointer] = items[i] // 如果它们不相同，则将项写入唯一指针的右侧。
		}
	}

	// 返回去重后的切片
	return items[:uniqPointer+1]
}

const DefaultVal = iota

func NewPriority[T constraints.Ordered](values []T, priority map[T]int) *Priority[T] {
	return &Priority[T]{
		values:   RemoveDuplicates(values),
		priority: priority,
	}
}

// Priority 定义一个包含需要排序的字符串的类型
type Priority[T constraints.Ordered] struct {
	values   []T
	priority map[T]int
}

// Len 实现 sort.Interface 接口的方法
func (s *Priority[T]) Len() int {
	return len(s.values)
}

func (s *Priority[T]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

// Less 按照自定义优先级排序
func (s *Priority[T]) Less(i, j int) bool {
	return s.priority[s.values[i]] > s.priority[s.values[j]]
}

// CustomSort 自定义排序函数，接受字符串切片和优先级映射作为参数
func (s *Priority[T]) CustomSort() {
	s.values = slices.DeleteFunc(s.values, func(t T) bool {
		_, exit := s.priority[t]
		return !exit
	})
	if len(s.values) < 2 {
		return
	}
	sort.Sort(s)
}

// Max 自定义排序函数，接受字符串切片和优先级映射作为参数
func (s *Priority[T]) Max() *T {
	if len(s.values) < 1 {
		return nil
	}
	s.CustomSort()
	if len(s.values) < 1 {
		return nil
	}

	return &s.values[0]
}

// Min 自定义排序函数，接受字符串切片和优先级映射作为参数
func (s *Priority[T]) Min() *T {
	if len(s.values) < 1 {
		return nil
	}
	s.CustomSort()
	if len(s.values) < 1 {
		return nil
	}
	return &s.values[len(s.values)-1]
}
