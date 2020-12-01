package matrix

import (
	"reflect"
	"testing"
)

func TestCheckVectors(t *testing.T) {
	type args struct {
		value []*Vector
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			want: true,
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
					V(2, 4, 5),
				},
			},
		},
		{
			name: "测试2",
			want: true,
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
				},
			},
		},
		{
			name: "测试3",
			want: false,
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
					V(2),
				},
			},
		},
		{
			name: "测试4",
			want: false,
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
					V(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckVectors(tt.args.value...); got != tt.want {
				t.Errorf("CheckVectors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		value []*Vector
	}
	tests := []struct {
		name string
		args args
		want *Elem
	}{
		{
			name: "测试1",
			want: &Elem{
				width:  3,
				height: 2,
				value: []*Vector{
					V(1, 3, 5),
					V(2, 4, 5),
				},
			},
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
					V(2, 4, 5),
				},
			},
		},

		{
			name: "测试2",
			want: &Elem{
				width:  3,
				height: 1,
				value: []*Vector{
					V(1, 3, 5),
				},
			},
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(1, 3, 5),
				},
			},
		},
		{
			name: "测试3",
			want: &Elem{
				width:  0,
				height: 1,
				value: []*Vector{
					V(),
				},
			},
			args: struct{ value []*Vector }{
				value: []*Vector{
					V(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV(t *testing.T) {
	type args struct {
		value []float64
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "测试1",
			args: struct{ value []float64 }{value: []float64{1, 2, 3, 4.5}},
			want: &Vector{
				dim:   4,
				value: []float64{1, 2, 3, 4.5},
			},
		},
		{
			name: "测试2",
			args: struct{ value []float64 }{value: []float64{}},
			want: &Vector{
				dim:   0,
				value: []float64{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := V(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("V() = %v, want %v", got, tt.want)
			}
		})
	}
}
