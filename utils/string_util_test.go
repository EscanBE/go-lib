package utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnyOf(t *testing.T) {
	tests := []struct {
		source string
		inAny  []string
		want   bool
	}{
		{
			source: "",
			inAny:  []string{"1", "2"},
			want:   false,
		},
		{
			source: "1",
			inAny:  []string{"1", "2"},
			want:   true,
		},
		{
			source: "3",
			inAny:  []string{"1", "2"},
			want:   false,
		},
		{
			source: "3",
			inAny:  []string{"333", "2", "33"},
			want:   false,
		},
		{
			source: "3",
			inAny:  []string{"333", "3", "33"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			if got := AnyOf(tt.source, tt.inAny...); got != tt.want {
				t.Errorf("AnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConditionalString(t *testing.T) {
	tests := []struct {
		expression bool
		whenTrue   string
		whenFalse  string
		want       string
	}{
		{
			expression: true,
			whenTrue:   "t",
			whenFalse:  "f",
			want:       "t",
		},
		{
			expression: false,
			whenTrue:   "t",
			whenFalse:  "f",
			want:       "f",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ConditionalString(tt.expression, tt.whenTrue, tt.whenFalse); got != tt.want {
				t.Errorf("ConditionalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonBlankString(t *testing.T) {
	tests := []struct {
		candidates []string
		want       string
		wantErr    bool
	}{
		{
			candidates: []string{},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"", ""},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
		{
			candidates: []string{"", "1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
		{
			candidates: []string{" ", ""},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"", " ", "1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
		{
			candidates: []string{"", "\t", "\n", "\r", ""},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"", "\t", "\n", "\r", "1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := FirstNonBlankString(tt.candidates...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNonBlankString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstNonBlankString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonEmptyString(t *testing.T) {
	tests := []struct {
		candidates []string
		want       string
		wantErr    bool
	}{
		{
			candidates: []string{},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"", ""},
			want:       "",
			wantErr:    true,
		},
		{
			candidates: []string{"1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
		{
			candidates: []string{"", "1", "", "2", ""},
			want:       "1",
			wantErr:    false,
		},
		{
			candidates: []string{"", " ", "1", "", "2", ""},
			want:       " ",
			wantErr:    false,
		},
		{
			candidates: []string{"", "\t", "1", "", "2", ""},
			want:       "\t",
			wantErr:    false,
		},
		{
			candidates: []string{"", "\n", "1", "", "2", ""},
			want:       "\n",
			wantErr:    false,
		},
		{
			candidates: []string{"", "\r", "1", "", "2", ""},
			want:       "\r",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := FirstNonEmptyString(tt.candidates...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNonEmptyString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstNonEmptyString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		text string
		want bool
	}{
		{
			text: "",
			want: true,
		},
		{
			text: " ",
			want: true,
		},
		{
			text: "\n",
			want: true,
		},
		{
			text: "\t",
			want: true,
		},
		{
			text: "\r",
			want: true,
		},
		{
			text: "    ",
			want: true,
		},
		{
			text: "\n\n\n",
			want: true,
		},
		{
			text: "\t\t\t",
			want: true,
		},
		{
			text: "\r\r\r",
			want: true,
		},
		{
			text: " \t\r\n",
			want: true,
		},
		{
			text: " \t\r\nx  \t\t\r\n\r\n",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if got := IsBlank(tt.text); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTrimmable(t *testing.T) {
	tests := []struct {
		text string
		want bool
	}{
		{
			text: "",
			want: false,
		},
		{
			text: " ",
			want: true,
		},
		{
			text: "\n",
			want: true,
		},
		{
			text: "\t",
			want: true,
		},
		{
			text: "\r",
			want: true,
		},
		{
			text: "    ",
			want: true,
		},
		{
			text: "\n\n\n",
			want: true,
		},
		{
			text: "\t\t\t",
			want: true,
		},
		{
			text: "\r\r\r",
			want: true,
		},
		{
			text: " \t\r\n",
			want: true,
		},
		{
			text: " \t\r\nx  \t\t\r\n\r\n",
			want: true,
		},
		{
			text: "x",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if got := IsTrimmable(tt.text); got != tt.want {
				t.Errorf("IsTrimmable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveEmptyString(t *testing.T) {
	tests := []struct {
		source []string
		want   []string
	}{
		{
			source: []string{"1", "2", "3"},
			want:   []string{"1", "2", "3"},
		},
		{
			source: []string{"1", "2", "3", "\r", "\n", "\t", " ", "\r\n", "\n\n", "\t\t", "  "},
			want:   []string{"1", "2", "3"},
		},
		{
			source: []string{},
			want:   []string{},
		},
		{
			source: nil,
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RemoveEmptyString(tt.source)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
