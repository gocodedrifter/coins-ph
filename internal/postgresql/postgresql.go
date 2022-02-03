package postgresql

import (
	"database/sql"
)

func newNullString(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid: len(str)>0,
	}
}

func newNullInt64(amount int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: amount,
		Valid: amount>=0,
	}
}
