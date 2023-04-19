package zcslice

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		slice []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				slice: []string{"a", "b", "c", "甲a", "乙b", "丙c"},
				str:   "乙b",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				slice: []string{"a", "b", "c"},
				str:   "乙b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.str); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
