package zcslice

import (
	"fmt"
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

func TestSubtract(t *testing.T) {
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
				slice2: []string{"a", "甲a", "c"},
			},
			want: []string{"b", "乙b", "丙c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
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

func TestReverseBytes(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				input: []byte{1, 2, 3, 4},
			},
			want: []byte{4, 3, 2, 1},
		},
		{
			name: "test2",
			args: args{
				input: []byte{13, 21, 3, 14, 53, 69, 71, 8},
			},
			want: []byte{8, 71, 69, 53, 14, 3, 21, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseBytes(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBytes() = %v, want %v", got, tt.want)
			}
			fmt.Printf("ReverseBytes() = %v, want %v\n", got, tt.want)
		})
	}
}
