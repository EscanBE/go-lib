package utils

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestPaging(t *testing.T) {
	type args struct {
		slice    []int
		pageSize int
	}

	generateArgs := func(pageSize int) args {
		return args{
			slice:    []int{1, 2, 3, 4, 5},
			pageSize: pageSize,
		}
	}

	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name:    "err when paging negative",
			args:    generateArgs(-1 * (rand.Int() + 1)),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "err when paging 0",
			args:    generateArgs(0),
			want:    nil,
			wantErr: true,
		},
		{
			name: "input empty slide should return empty result",
			args: args{
				slice:    make([]int, 0),
				pageSize: rand.Int() + 1, /*always > 0*/
			},
			want:    [][]int{{}},
			wantErr: false,
		},
		{
			name:    "1",
			args:    generateArgs(1),
			want:    [][]int{{1}, {2}, {3}, {4}, {5}},
			wantErr: false,
		},
		{
			name:    "2",
			args:    generateArgs(2),
			want:    [][]int{{1, 2}, {3, 4}, {5}},
			wantErr: false,
		},
		{
			name:    "3",
			args:    generateArgs(3),
			want:    [][]int{{1, 2, 3}, {4, 5}},
			wantErr: false,
		},
		{
			name:    "4",
			args:    generateArgs(4),
			want:    [][]int{{1, 2, 3, 4}, {5}},
			wantErr: false,
		},
		{
			name:    "5",
			args:    generateArgs(5),
			want:    [][]int{{1, 2, 3, 4, 5}},
			wantErr: false,
		},
		{
			name:    "6",
			args:    generateArgs(6),
			want:    [][]int{{1, 2, 3, 4, 5}},
			wantErr: false,
		},
		{
			name:    "7",
			args:    generateArgs(7),
			want:    [][]int{{1, 2, 3, 4, 5}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Paging(tt.args.slice, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Paging() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Paging() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUniqueElements(t *testing.T) {
	tests := []struct {
		slice []int
		want  []int
	}{
		{
			slice: []int{},
			want:  []int{},
		},
		{
			slice: []int{1},
			want:  []int{1},
		},
		{
			slice: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			slice: []int{1, 2, 3, 2, 3, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := GetUniqueElements(tt.slice...)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUniqueElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
