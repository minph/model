package vector

import (
	"reflect"
	"testing"
)

func TestElem_Count(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 0},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElem_Filter(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	type args struct {
		filterFunc FilterFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Elem
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{-2, -1, 0, 1, 2, 3}, index: 0},
			want: New(-2, 0, 2),
			args: struct{ filterFunc FilterFunc }{filterFunc: func(index int, value interface{}) bool {
				if v, ok := value.(int); ok {
					return v%2 == 0
				}
				return false
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Filter(tt.args.filterFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElem_First(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElem_Index(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Index(); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElem_Next(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 2},
			want: nil,
		}, {
			name: "测试2",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElem_Prev(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 0},
			want: nil,
		}, {
			name: "测试2",
			fields: struct {
				value []interface{}
				index int
			}{value: []interface{}{1, 2, 3}, index: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				value: []interface{}{1, 2, 3, 4, 5},
				index: 2,
			},
			want: &Elem{
				value: []interface{}{1, 2, 4, 5},
				index: 1,
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

func TestElem_Insert(t *testing.T) {
	type fields struct {
		value []interface{}
		index int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Elem
	}{
		{
			name: "测试1",
			fields: struct {
				value []interface{}
				index int
			}{
				value: []interface{}{1, 2, 3, 4, 5},
				index: 2,
			},
			want: &Elem{
				value: []interface{}{1, 2, 9, 3, 4, 5},
				index: 2,
			},
			args: struct{ value interface{} }{value: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elem{
				value: tt.fields.value,
				index: tt.fields.index,
			}
			if got := e.Insert(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
