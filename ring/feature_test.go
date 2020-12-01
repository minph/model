package ring

import (
	"reflect"
	"testing"
)

func TestElem_Remove(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Elem
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{
				value: []interface{}{0, 1, 2, 3, 4},
				index: 3,
			},
			want: &Elem{
				value: []interface{}{0, 1, 2, 4},
				index: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Remove(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
