// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package vatdb

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DataInstance struct {
	ID        pgtype.UUID
	Timestamp pgtype.Timestamp
	Value     []byte
}
