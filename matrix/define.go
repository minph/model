package matrix

// Elem 矩阵结构体
type Elem struct {
	width  int       // 矩阵列数
	height int       // 矩阵行数
	value  []*Vector // 矩阵元素值
}

// Vector 向量构体
type Vector struct {
	dim   int       // 向量维数
	value []float64 // 向量元素值
}

// CheckVectors 检查多个向量是否可转为矩阵
func CheckVectors(value ...*Vector) bool {
	dim := value[0].dim

	for _, item := range value {
		if dim != item.dim {
			return false
		}
	}
	return true
}

// New 创建矩阵
func New(value ...*Vector) *Elem {

	if !CheckVectors(value...) {
		return nil
	}

	return &Elem{
		width:  value[0].dim,
		height: len(value),
		value:  value,
	}
}

// V 创建向量
func V(value ...float64) *Vector {
	return &Vector{
		dim:   len(value),
		value: value,
	}
}
