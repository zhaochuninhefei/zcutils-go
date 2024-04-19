package protobuffer

import (
	"reflect"
	"testing"
)

func TestBuffer_Bytes(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if got := b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeFixed32(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeFixed32()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFixed32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeFixed32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeFixed64(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeFixed64()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFixed64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeFixed64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeRawBytes(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		alloc bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeRawBytes(tt.args.alloc)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeRawBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeRawBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeVarint(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				buf:           make([]byte, 1),
				idx:           0,
				deterministic: false,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeVarint()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeVarint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeVarint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeZigzag32(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeZigzag32()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeZigzag32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeZigzag32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DecodeZigzag64(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			got, err := b.DecodeZigzag64()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeZigzag64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeZigzag64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_EncodeFixed32(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeFixed32(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeFixed32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_EncodeFixed64(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeFixed64(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeFixed64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_EncodeRawBytes(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeRawBytes(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeRawBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_EncodeVarint(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeVarint(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeVarint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_EncodeZigzag32(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeZigzag32(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeZigzag32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_EncodeZigzag64(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if err := b.EncodeZigzag64(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeZigzag64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_Reset(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			b.Reset()
		})
	}
}

func TestBuffer_SetBuf(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			b.SetBuf(tt.args.buf)
		})
	}
}

func TestBuffer_SetDeterministic(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	type args struct {
		deterministic bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			b.SetDeterministic(tt.args.deterministic)
		})
	}
}

func TestBuffer_Unread(t *testing.T) {
	type fields struct {
		buf           []byte
		idx           int
		deterministic bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				buf:           tt.fields.buf,
				idx:           tt.fields.idx,
				deterministic: tt.fields.deterministic,
			}
			if got := b.Unread(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unread() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeVarint(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DecodeVarint(tt.args.b)
			if got != tt.want {
				t.Errorf("DecodeVarint() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DecodeVarint() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEncodeVarint(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeVarint(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeVarint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBuffer(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want *Buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuffer(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSizeVarint(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SizeVarint(tt.args.v); got != tt.want {
				t.Errorf("SizeVarint() = %v, want %v", got, tt.want)
			}
		})
	}
}
