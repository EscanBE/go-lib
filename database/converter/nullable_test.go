package converter

import (
	"database/sql"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

type TestJsonStruct struct {
	F1 int    `json:"f1"`
	F2 int    `json:"f2"`
	F3 string `json:"f3,omitempty"`
}

func TestJsonAsSqlNullableString(t *testing.T) {
	type args struct {
		v         any
		sureEmpty bool
	}
	tests := []struct {
		name    string
		args    args
		want    sql.NullString
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				v: TestJsonStruct{
					F1: 1,
					F2: 2,
				},
				sureEmpty: false,
			},
			want: sql.NullString{
				String: "{\"f1\":1,\"f2\":2}",
				Valid:  true,
			},
			wantErr: false,
		},
		{
			name: "sure empty will take priority even tho input is not empty",
			args: args{
				v: TestJsonStruct{
					F1: 1,
				},
				sureEmpty: true,
			},
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
			wantErr: false,
		},
		{
			name: "failed to marshal",
			args: args{
				v: func() {},
			},
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonAsSqlNullableString(tt.args.v, tt.args.sureEmpty)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonAsSqlNullableString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonAsSqlNullableString() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("keep invalid byte sequence", func(t *testing.T) {
		got, err := JsonAsSqlNullableString(TestJsonStruct{
			F1: 1,
			F2: 2,
			F3: "a\x80\xFF",
		}, false)
		if err != nil {
			t.Errorf("JsonAsSqlNullableString() error = %v, want no err", err)
			return
		}
		if !got.Valid {
			t.Errorf("JsonAsSqlNullableString() result need to be valid")
			return
		}
		if strings.Contains(got.String, "\"a\"") {
			t.Errorf("JsonAsSqlNullableString() result string should not remove invalid UTF8")
			return
		}
	})
}

func TestUtf8StringAsSqlNullableString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  sql.NullString
	}{
		{
			name:  "success",
			input: "123",
			want: sql.NullString{
				String: "123",
				Valid:  true,
			},
		},
		{
			name:  "success empty",
			input: "",
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			name:  "Should remove invalid UTF-8",
			input: "a\x80\xFF",
			want: sql.NullString{
				String: "a",
				Valid:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8StringAsSqlNullableString(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Utf8StringAsSqlNullableString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringAsSqlNullableString(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		removeInvalidUtf8 bool
		want              sql.NullString
	}{
		{
			name:              "do not remove invalid UTF-8",
			input:             "a\x80\xFF",
			removeInvalidUtf8: false,
			want: sql.NullString{
				String: "a\x80\xFF",
				Valid:  true,
			},
		},
		{
			name:              "remove invalid UTF-8",
			input:             "a\x80\xFF",
			removeInvalidUtf8: true,
			want: sql.NullString{
				String: "a",
				Valid:  true,
			},
		},
		{
			name:              "remove invalid UTF-8 and become null",
			input:             "\x80\xFF",
			removeInvalidUtf8: true,
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			name:              "empty is null",
			input:             "",
			removeInvalidUtf8: false,
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			name:              "blank is not null",
			input:             " \t\r\n",
			removeInvalidUtf8: false,
			want: sql.NullString{
				String: " \t\r\n",
				Valid:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringAsSqlNullableString(tt.input, tt.removeInvalidUtf8); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringAsSqlNullableString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooleanAsSqlNullableBoolean(t *testing.T) {
	tests := []struct {
		name  string
		input bool
		want  sql.NullBool
	}{
		{
			name:  "true is true not null",
			input: true,
			want: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		},
		{
			name:  "false is null",
			input: false,
			want: sql.NullBool{
				Bool:  false,
				Valid: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BooleanAsSqlNullableBoolean(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanAsSqlNullableBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigIntAsNullableSqlInt64(t *testing.T) {
	tests := []struct {
		name  string
		input *big.Int
		want  sql.NullInt64
	}{
		{
			name:  "pointer is nil should be NULL",
			input: nil,
			want: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		},
		{
			name:  "pointer is not null should not be NULL",
			input: &big.Int{},
			want: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
		},
		{
			name:  "pointer is not null with value should not be NULL",
			input: new(big.Int).SetInt64(999),
			want: sql.NullInt64{
				Int64: 999,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigIntAsNullableSqlInt64(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigIntAsNullableSqlInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
