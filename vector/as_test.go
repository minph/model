package vector

import (
	"reflect"
	"testing"
)

func TestAsInt(t *testing.T) {
	type args struct {
		e Feature
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "测试1",
			want: []int{1, 3},
			args: struct{ e Feature }{e: New(1, '2', 3, '4', 5.5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsInt(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsFloat(t *testing.T) {
	type args struct {
		e Feature
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "测试1",
			want: []float64{0.5, 5.5},
			args: struct{ e Feature }{e: New(0.5, 1, '2', 3, '4', 5.5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsFloat(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsNum(t *testing.T) {
	type args struct {
		e Feature
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "测试1",
			want: []float64{0.5, 1, 3, 5.5},
			args: struct{ e Feature }{e: New(0.5, 1, '2', 3, '4', 5.5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsNum(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsString(t *testing.T) {
	type args struct {
		e Feature
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "测试1",
			want: []string{"0.5,1", "5.5"},
			args: struct{ e Feature }{e: New("0.5,1", '2', 3, '4', "5.5")},
		},
		{
			name: "测试2",
			want: []string{"你好啊", "你好"},
			args: struct{ e Feature }{e: New("你好啊", '你', 3, '好', "你好")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsString(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
