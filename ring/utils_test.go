package ring

import "testing"

func Test_replace(t *testing.T) {
	type args struct {
		index int
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: struct {
				index int
				count int
			}{
				index: 1,
				count: 6,
			},
			want: 1,
		},

		{
			name: "测试2",
			args: struct {
				index int
				count int
			}{
				index: -7,
				count: 6,
			},
			want: 5,
		},
		{
			name: "测试3",
			args: struct {
				index int
				count int
			}{
				index: 16,
				count: 6,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.index, tt.args.count); got != tt.want {
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
