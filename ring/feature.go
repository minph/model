package ring

import "github.com/minph/model/vector"

type Feature interface {
	Count() int
	Get() interface{}
	Insert(value interface{}) *Elem
	Next() interface{}
	Prev() interface{}
	Remove() *Elem
	Turn(index int) *Elem
	Vector() *vector.Elem
}

// Get 获取当前元素
func (e *Elem) Get() interface{} {
	return e.value[e.index]
}

// Count 获取所有元素个数
func (e *Elem) Count() int {
	return len(e.value)
}

// Prev 获取上一个元素
func (e *Elem) Prev() interface{} {
	e.index = replace(e.index-1, e.Count())
	return e.Get()
}

// Next 获取下一个元素
func (e *Elem) Next() interface{} {
	e.index = replace(e.index+1, e.Count())
	return e.Get()
}

// Remove 删去当前元素，然后指针指向上一位
func (e *Elem) Remove() *Elem {
	r := vector.New(e.value...)
	r.Turn(e.index)
	r.Remove()
	return FromVector(r)
}

// Turn 指针转到指定位置
func (e *Elem) Turn(index int) *Elem {
	e.index = replace(index, e.Count())
	return e
}

// Vector 转为向量集合
func (e *Elem) Vector() *vector.Elem {
	r := vector.New(e.value)
	r.Turn(e.index)
	return r
}

// Insert 在当前位置增加元素并指向此元素
func (e *Elem) Insert(value interface{}) *Elem {
	r := e.Vector()
	r.Insert(value)
	return FromVector(r)
}
