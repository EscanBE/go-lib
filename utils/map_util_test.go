package utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetKeys(t *testing.T) {
	type args struct {
		myMap map[int]bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				myMap: map[int]bool{},
			},
			want: make([]int, 0),
		},
		{
			name: "one",
			args: args{
				myMap: map[int]bool{1: false},
			},
			want: []int{1},
		},
		{
			name: "two",
			args: args{
				myMap: map[int]bool{1: false, 2: true},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKeys(tt.args.myMap)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKeysOf(t *testing.T) {
	type args struct {
		myMap         map[int]bool
		expectedValue bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				myMap: map[int]bool{
					1: true,
					2: false,
					3: true,
					4: false,
					5: false,
				},
				expectedValue: true,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				myMap: map[int]bool{
					1: true,
					2: false,
					3: true,
					4: false,
					5: false,
				},
				expectedValue: false,
			},
			want: []int{2, 4, 5},
		},
		{
			args: args{
				myMap: map[int]bool{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				},
				expectedValue: false,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKeysOf(tt.args.myMap, tt.args.expectedValue)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeysOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKeysOfTrue(t *testing.T) {
	type args struct {
		myMap map[int]bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				myMap: map[int]bool{
					1: true,
					2: false,
					3: true,
					4: false,
				},
			},
			want: []int{1, 3},
		},
		{
			args: args{
				myMap: map[int]bool{
					1: false,
					2: false,
					3: false,
					4: false,
				},
			},
			want: make([]int, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKeysOfTrue(tt.args.myMap)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeysOfTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoftCloneMap(t *testing.T) {
	var arr []byte
	v := &arr
	type args struct {
		myMap map[int]*[]byte
	}
	tests := []struct {
		name string
		args args
		want map[int]*[]byte
	}{
		{
			args: args{
				myMap: map[int]*[]byte{
					1: v,
					2: v,
				},
			},
			want: map[int]*[]byte{
				1: v,
				2: v,
			},
		},
		{
			args: args{
				myMap: map[int]*[]byte{},
			},
			want: map[int]*[]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SoftCloneMap(tt.args.myMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SoftCloneMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlideToTracker(t *testing.T) {
	tests := []struct {
		slice []int
		want  map[int]bool
	}{
		{
			slice: nil,
			want:  map[int]bool{},
		},
		{
			slice: []int{},
			want:  map[int]bool{},
		},
		{
			slice: []int{1},
			want:  map[int]bool{1: true},
		},
		{
			slice: []int{1, 2, 3},
			want:  map[int]bool{1: true, 2: true, 3: true},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SlideToTracker(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlideToTracker() = %v, want %v", got, tt.want)
			}
		})
	}
}