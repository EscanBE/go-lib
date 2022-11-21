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

func TestPutToMapAsKeys(t *testing.T) {
	tests := []struct {
		name         string
		tracker      map[int]bool
		slice        []int
		defaultValue bool
		behavior     PutToMapAsKeyBehavior
		want         map[int]bool
		wantErr      bool
	}{
		{
			name:    "empty slice",
			tracker: map[int]bool{1: false},
			slice:   []int{},
			want:    map[int]bool{1: false},
		},
		{
			name:         "RejectAllWhenAnyDuplicatedKey(1)",
			tracker:      map[int]bool{1: false},
			slice:        []int{2},
			defaultValue: true,
			behavior:     RejectAllWhenAnyDuplicatedKey,
			want:         map[int]bool{1: false, 2: true},
		},
		{
			name:         "RejectAllWhenAnyDuplicatedKey(2)",
			tracker:      map[int]bool{1: false, 2: false},
			slice:        []int{2, 3},
			defaultValue: true,
			behavior:     RejectAllWhenAnyDuplicatedKey,
			want:         map[int]bool{1: false, 2: false},
			wantErr:      true,
		},
		{
			name:         "SkipDuplicatedKeys(1)",
			tracker:      map[int]bool{1: false, 2: false},
			slice:        []int{2, 3},
			defaultValue: true,
			behavior:     SkipDuplicatedKeys,
			want:         map[int]bool{1: false, 2: false, 3: true},
		},
		{
			name:         "SkipDuplicatedKeys(2)",
			tracker:      map[int]bool{1: false, 2: false},
			slice:        []int{3, 4},
			defaultValue: true,
			behavior:     SkipDuplicatedKeys,
			want:         map[int]bool{1: false, 2: false, 3: true, 4: true},
		},
		{
			name:         "AcceptAllAndOverrideDuplicatedKeys(1)",
			tracker:      map[int]bool{1: false},
			slice:        []int{1, 2},
			defaultValue: true,
			behavior:     AcceptAllAndOverrideDuplicatedKeys,
			want:         map[int]bool{1: true, 2: true},
		},
		{
			name:         "AcceptAllAndOverrideDuplicatedKeys(2)",
			tracker:      map[int]bool{1: false},
			slice:        []int{2, 3},
			defaultValue: true,
			behavior:     AcceptAllAndOverrideDuplicatedKeys,
			want:         map[int]bool{1: false, 2: true, 3: true},
		},
		{
			name:         "AcceptOnlyDuplicatedKeysAndOverrideThem(1)",
			tracker:      map[int]bool{1: false, 2: true, 4: false},
			slice:        []int{1, 2, 3},
			defaultValue: true,
			behavior:     AcceptOnlyDuplicatedKeysAndOverrideThem,
			want:         map[int]bool{1: true, 2: true, 4: false},
		},
		{
			name:         "not supported behavior",
			tracker:      map[int]bool{1: false},
			slice:        []int{1, 2},
			defaultValue: true,
			behavior:     PutToMapAsKeyBehavior(99),
			want:         map[int]bool{1: false},
			wantErr:      true,
		},
		{
			name:    "tracker can not be nil",
			tracker: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PutToMapAsKeys(tt.tracker, tt.slice, tt.defaultValue, tt.behavior)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("SlideToTracker() = %t, want %t", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.tracker, tt.want) {
				t.Errorf("SlideToTracker() map %v, want %v", tt.tracker, tt.want)
			}
		})
	}
}
