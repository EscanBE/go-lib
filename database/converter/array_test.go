package converter

import (
	"fmt"
	"github.com/lib/pq"
	"testing"
)

func TestToPostgresArray(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  interface{}
	}{
		{
			name:  "nil is GenericArray",
			input: nil,
			want:  pq.GenericArray{},
		},
		{
			name:  "bool is BoolArray",
			input: []bool{true},
			want:  pq.BoolArray{},
		},
		{
			name:  "float64 is Float64Array",
			input: []float64{},
			want:  pq.Float64Array{},
		},
		{
			name:  "float32 is Float32Array",
			input: []float32{},
			want:  pq.Float32Array{},
		},
		{
			name:  "int64 is Int64Array",
			input: []int64{},
			want:  pq.Int64Array{},
		},
		{
			name:  "int32 is Int32Array",
			input: []int32{},
			want:  pq.Int32Array{},
		},
		{
			name:  "int is GenericArray",
			input: []int{},
			want:  pq.GenericArray{},
		},
		{
			name:  "int16 is GenericArray",
			input: []int16{},
			want:  pq.GenericArray{},
		},
		{
			name:  "int8 is GenericArray",
			input: []int8{},
			want:  pq.GenericArray{},
		},
		{
			name:  "string is StringArray",
			input: []string{},
			want:  pq.StringArray{},
		},
		{
			name:  "bytea is ByteaArray",
			input: [][]byte{},
			want:  pq.ByteaArray{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToPostgresArray(tt.input)
			if !sameType(got, tt.want) {
				t.Errorf("ToPostgresArray() returns un-expected type! Got %T, want %T or *%T", got, tt.want, tt.want)
			}
		})
	}
}

func sameType(got, want interface{}) bool {
	return fmt.Sprintf("%T", got) == fmt.Sprintf("%T", want) ||
		fmt.Sprintf("%T", got) == fmt.Sprintf("*%T", want)
}
