package matrix_test

import (
	"fmt"
	"github.com/minph/model/matrix"
)

func ExampleNew() {

	// 创建 2*2 矩阵
	A := matrix.New(
		matrix.V(1, 2),
		matrix.V(3, 4),
	)

	// B 是 A 的转置
	B := matrix.New(
		matrix.V(1, 3),
		matrix.V(2, 4),
	)

	fmt.Printf("%#v\n", A)
	fmt.Printf("%#v\n", B)

}
