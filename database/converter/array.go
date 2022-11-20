package converter

import (
	"database/sql"
	"database/sql/driver"
	"github.com/lib/pq"
)

// ToPostgresArray transforms the input array into array in underlying database
func ToPostgresArray(a interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	return pq.Array(a)
}
