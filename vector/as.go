package vector

// AsInt 退化为 []int
func AsInt(e Feature) []int {

	var r []int

	e.Foreach(func(index int, value interface{}) {
		if v, ok := value.(int); ok {
			r = append(r, v)
		}
	})

	return r
}

// AsInt 退化为 []float64
func AsFloat(e Feature) []float64 {

	var r []float64

	e.Foreach(func(index int, value interface{}) {
		if v, ok := value.(float64); ok {
			r = append(r, v)
		}
	})

	return r
}

// AsNum 退化为 []float64，且 int 类型也会包含
func AsNum(e Feature) []float64 {

	var r []float64

	e.Foreach(func(index int, value interface{}) {
		if v, ok := value.(float64); ok {
			r = append(r, v)
		}
		// 此时int 类型也会包含
		if v, ok := value.(int); ok {
			r = append(r, float64(v))
		}
	})

	return r
}

// AsString 退化为 []string
func AsString(e Feature) []string {

	var r []string

	e.Foreach(func(index int, value interface{}) {
		if v, ok := value.(string); ok {
			r = append(r, v)
		}
	})

	return r
}

// 生成相应接口方法

func (e *Elem) AsInt() []int {
	return AsInt(e)
}

func (e *Elem) AsFloat() []float64 {
	return AsFloat(e)
}

func (e *Elem) AsNum() []float64 {
	return AsNum(e)
}

func (e *Elem) AsString() []string {
	return AsString(e)
}
