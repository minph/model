package ring

import "github.com/minph/model/vector"

// Elem 数据环结构体
type Elem struct {
	value []interface{} // 所有元素
	index int           // 当前指针位置
}

// New 创建数据环
func New(values ...interface{}) *Elem {
	return &Elem{
		value: values,
		index: 0,
	}
}

// FromVector 从向量集合创建数据环
func FromVector(vector *vector.Elem) *Elem {
	return &Elem{
		value: vector.Value(),
		index: replace(vector.Index(), vector.Count()),
	}
}
