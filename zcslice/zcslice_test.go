package zcslice

import (
	"reflect"
	"testing"
)

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

func TestDiff(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				slice1: []string{"a", "b", "c", "甲a", "乙b", "丙c"},
				slice2: []string{"a", "b", "c"},
			},
			want: []string{"甲a", "乙b", "丙c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimAndRmSpace(t *testing.T) {
	type args struct {
		slice1 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				slice1: []string{"a ", " b", " c	", "	", "  乙  b ", "丙c"},
			},
			want: []string{"a", "b", "c", "乙  b", "丙c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAndRmSpace(tt.args.slice1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimAndRmSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
