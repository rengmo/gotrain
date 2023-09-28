package week1

import (
	"fmt"
)

func init() {
	ms := NewMySlice[int]()
	ms.Add(1, 2, 3, 4, 5, 6, 7, 8)
	ms.Delete(0)
	ms.Add(9)
	ms.Delete(1)
	ms.Delete(2)
	ms.Delete(3)
	ms.Delete(4)
	ms.Delete(5)
	ms.Add(-100)
	fmt.Println(ms.Find(5))
	ms.Print()
}

type MySlice[T comparable] struct {
	// 虽然单独维护len和cap增加了复杂度，但是可以不用在每次删除的时候都挪动元素，提高了删除的性能
	len   int
	cap   int
	elems []*T
}

func NewMySlice[T comparable]() MySlice[T] {
	ms := MySlice[T]{
		len:   0,
		cap:   0,
		elems: make([]*T, 0),
	}
	return ms
}

func (ms *MySlice[T]) Delete(idx int) error {
	// 范围限制
	if idx < 0 || idx >= ms.cap {
		return fmt.Errorf("index out of range")
	}
	var empty *T
	// 该位置已经没元素，直接返回
	if ms.elems[idx] == empty {
		return fmt.Errorf("elemet has been deleted")
	}

	// 删除该位置的元素
	ms.elems[idx] = empty
	ms.len -= 1
	ms.reduceCap()
	return nil
}

func (ms *MySlice[T]) Add(elems ...T) {
	for i := 0; i < len(elems); i++ {
		ms.elems = append(ms.elems, &elems[i])
	}
	ms.len += len(elems)
	ms.cap += len(elems)
}

func (ms *MySlice[T]) Find(elem T) (int, error) {
	for index, val := range ms.elems {
		if val == nil {
			continue
		}
		if *val == elem {
			return index, nil
		}
	}
	return -1, fmt.Errorf("element not exist")
}

func (ms *MySlice[T]) reduceCap() {
	if float64(ms.len)/float64(ms.cap) < 0.6 {
		newElems := make([]*T, ms.len)
		var i int
		for _, val := range ms.elems {
			if val != nil {
				newElems[i] = val
				i++
			}
		}
		ms.cap = ms.len
		ms.elems = newElems
	}
}

func (ms *MySlice[T]) Print() {
	var empty T
	elems := make([]T, len(ms.elems))
	for index, val := range ms.elems {
		if val == nil {
			elems[index] = empty
		} else {
			elems[index] = *val
		}
	}
	fmt.Printf("%v", elems)
}
