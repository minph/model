package vector

// get 不越界获取元素，私有方法，如果越界返回 nil
func (e *Elem) get(index int) interface{} {
	if 0 <= index && index < e.Count() {
		return e.value[index]
	} else if index < 0 {
		// 指针指向零
		e.index = 0
	} else if index >= e.Count() {
		// 指针指向末尾
		index = e.Count() - 1
	}

	return nil

}

// Add 在向量集合头部增加元素，当前指针位置元素不变
func (e *Elem) Add(value interface{}) *Elem {
	var r []interface{}
	r = append(r, value)
	r = append(r, e.value...)
	e.value = r
	e.Next()
	return e
}

// Count 获取所有元素个数
func (e *Elem) Count() int {
	return len(e.value)
}

// Filter 从原向量集合过滤产生新向量集合
func (e *Elem) Filter(filterFunc FilterFunc) *Elem {
	newList := New()
	e.Foreach(func(index int, value interface{}) {
		if filterFunc(index, value) {
			newList.Push(value)
		}
	})
	return newList
}

// Get 获取当前元素
func (e *Elem) Get() interface{} {
	return e.get(e.index)
}

// Index 获取元素位置
func (e *Elem) Index() int {
	return e.index
}

// First 获取第一个元素
func (e *Elem) First() interface{} {
	e.index = 0
	return e.get(e.index)
}

// Foreach 遍历元素
func (e *Elem) Foreach(rangeFunc MapFunc) {

	for index := 0; index < e.Count(); index++ {
		// 执行函数
		rangeFunc(index, e.get(index))
	}
}

// Last 获取最后一个元素
func (e *Elem) Last() interface{} {
	e.index = e.Count() - 1
	return e.get(e.index)
}

// Map 从原向量集合派生新向量集合
func (e *Elem) Map(deriveFunc DeriveFunc) *Elem {
	newList := New()
	e.Foreach(func(index int, value interface{}) {
		newList.Push(deriveFunc(index, value))
	})
	return newList
}

// Next 获取下一个元素
func (e *Elem) Next() interface{} {
	e.index++ // 指针加1
	return e.get(e.index)
}

// Prev 获取上一个元素
func (e *Elem) Prev() interface{} {
	e.index-- // 指针减1
	return e.get(e.index)
}

// Reverse 生成反向向量集合
func (e *Elem) Reverse() *Elem {
	var r = &Elem{}
	for index := e.Count() - 1; index >= 0; index-- {
		r.value = append(r.value, e.get(index))
	}
	return r
}

// ReverseForeach 反向遍历元素
func (e *Elem) ReverseForeach(rangeFunc MapFunc) {

	for index := e.Count() - 1; index >= 0; index-- {
		// 执行函数
		rangeFunc(index, e.get(index))
	}
}

// Push 在向量集合尾部增加元素，当前指针位置元素不变
func (e *Elem) Push(value interface{}) *Elem {
	e.value = append(e.value, value)
	return e
}

// Remove 移除当前指针位置元素，然后指针指向上一位
func (e *Elem) Remove() *Elem {
	var r []interface{}
	r = append(r, e.value[0:e.index]...)
	r = append(r, e.value[e.index+1:]...)
	e.value = r
	e.Prev()
	return e
}

// Shift 删除第一个元素，并返回第一个元素的值，然后指针指向上一位
func (e *Elem) Shift() interface{} {
	value := e.get(0)
	e.value = e.value[1:]
	e.Prev()
	return value
}

// Turn 将指针转到指定位置，越界则指向第一个元素
func (e *Elem) Turn(index int) *Elem {
	if 0 <= index && index < e.Count() {
		e.index = index
	} else {
		// 指针归零
		e.index = 0
	}
	return e
}

// Value 获取所有元素
func (e *Elem) Value() []interface{} {
	return e.value
}

// Insert 在指针位置加入元素并指向此元素
func (e *Elem) Insert(value interface{}) *Elem {
	var r []interface{}
	r = append(r, e.value[0:e.index]...)
	r = append(r, value)
	r = append(r, e.value[e.index:]...)
	e.value = r
	return e
}
