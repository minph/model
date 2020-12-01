package vector_test

import (
	"fmt"
	"github.com/minph/model/vector"
)

func ExampleNew() {
	v := vector.New(1, 2, 3).Push(4)
	fmt.Printf("%#v", v)

	// Output
	// &vector.Elem{value:[]interface {}{1, 2, 3, 4}, index:0}

}
