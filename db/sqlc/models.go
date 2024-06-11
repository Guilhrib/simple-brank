// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt pgtype.Timestamptz
}

type Entry struct {
	ID        int64
	AccountID pgtype.Int8
	// It can be negative or positive
	Amount    int64
	CreatedAt pgtype.Timestamptz
}

type Transfer struct {
	ID            int64
	FromAccountID pgtype.Int8
	ToAccountID   pgtype.Int8
	// It must be positive
	Amount    int64
	CreatedAt pgtype.Timestamptz
}
