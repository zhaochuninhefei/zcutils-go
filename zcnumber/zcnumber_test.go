package zcnumber

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestJSONInt64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		j       JSONInt64
		want    []byte
		wantErr bool
	}{
		{
			name: "test01",
			j:    JSONInt64(123),
			want: []byte(`123`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONInt64_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		j       JSONInt64
		args    args
		wantErr bool
	}{
		{
			name: "test",
			j:    JSONInt64(123),
			args: args{
				data: []byte("123"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type WhatEver struct {
	Order    JSONInt64 `json:"order"`
	OrderStr JSONInt64 `json:"orderStr"`
	// 假设这里还有其他字段
	Name string `json:"name"`
}

func TestUnmarshalJSONStrAndInt(t *testing.T) {
	jsonData := `{"order": 12345, "orderStr": "54321", "name": "test"}`
	var what WhatEver

	err := json.Unmarshal([]byte(jsonData), &what)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	want := WhatEver{
		Order:    12345,
		OrderStr: 54321,
		Name:     "test",
	}

	fmt.Printf("%+v\n", what) // 正确解析为int64类型，同时不影响其他字段
	// 比较 want 和 what
	if reflect.DeepEqual(what, want) {
		fmt.Println("PASS")
	} else {
		t.Fatal("FAIL")
	}
}
