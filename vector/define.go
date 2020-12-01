package vector

// Elem 向量集合结构体
type Elem struct {
	value []interface{}
	index int // 当前元素位置
}

// MapFunc 遍历函数
type MapFunc func(index int, value interface{})

// DeriveFunc 派生函数
type DeriveFunc func(index int, value interface{}) interface{}

// FilterFunc 过滤函数
type FilterFunc func(index int, value interface{}) bool

// New 创建向量集合
func New(values ...interface{}) *Elem {
	return &Elem{
		value: values,
		index: 0,
	}
}

// Feature 向量集合接口
type Feature interface {
	Add(value interface{}) *Elem
	Count() int
	Filter(filterFunc FilterFunc) *Elem
	Get() interface{}
	Index() int
	Insert(value interface{}) *Elem
	First() interface{}
	Foreach(rangeFunc MapFunc)
	Last() interface{}
	Map(deriveFunc DeriveFunc) *Elem
	Next() interface{}
	Prev() interface{}
	Reverse() *Elem
	ReverseForeach(rangeFunc MapFunc)
	Push(value interface{}) *Elem
	Remove() *Elem
	Shift() interface{}
	Turn(index int) *Elem
	Value() []interface{}
}

type AsType interface {
	AsInt() []int
	AsFloat() []float64
	AsNum() []float64
	AsString() []string
}
