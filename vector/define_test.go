package vector

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		values []interface{}
	}

	tests := []struct {
		name string
		args args
		want *Elem
	}{
		{
			name: "测试1",
			want: New(1, 2, 3, 4),
			args: struct{ values []interface{} }{values: []interface{}{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
