package converter

import (
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

// JsonAsSqlNullableString do transform an object into a string that nullable in the underlying database
func JsonAsSqlNullableString(v interface{}, len int) (sql.NullString, error) {
	if v == nil || len < 1 {
		return sql.NullString{}, nil
	}
	bz, err := json.Marshal(v)
	if err != nil {
		return sql.NullString{}, errors.Wrap(err, "problem while trying to marshal a nullable json")
	}
	return StringAsSqlNullableString(string(bz), false), nil
}

// Utf8StringAsSqlNullableString receives a string, then remove invalid UTF characters and then transform the string into a string object that nullable in the underlying database, if the string is empty, it will be NULL
func Utf8StringAsSqlNullableString(s string) sql.NullString {
	return StringAsSqlNullableString(s, true)
}

// StringAsSqlNullableString do transform a string into a string object that nullable in the underlying database, if the string is empty, it will be NULL
func StringAsSqlNullableString(s string, removeInvalidUtf8 bool) sql.NullString {
	if len(s) < 1 {
		return sql.NullString{}
	}
	if removeInvalidUtf8 {
		s = strings.ToValidUTF8(s, "")
		if len(s) < 1 {
			return sql.NullString{}
		}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// BooleanAsSqlNullableBoolean do transform a boolean into a boolean object that nullable in the underlying database, if the boolean value is FALSE, it will be NULL
func BooleanAsSqlNullableBoolean(b bool) sql.NullBool {
	if !b {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  true,
		Valid: true,
	}
}

// BigIntAsNullableSqlInt64 do transform a big.Int into a int64 object that nullable in the underlying database, if the pointer is NIL, it will be NULL
func BigIntAsNullableSqlInt64(s *big.Int) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: s.Int64(),
		Valid: true,
	}
}
