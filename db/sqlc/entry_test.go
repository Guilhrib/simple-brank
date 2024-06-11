package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"studies.com/util"
)

func createNewEntry(t *testing.T, account1 Account) Entry {
	arg := CreateEntryParams{
		AccountID: pgtype.Int8{
			Int64: account1.ID,
			Valid: true,
		},
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, entry.ID)
	require.Equal(t, account1.ID, entry.AccountID.Int64)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := createNewAccount(t)
	createNewEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account1 := createNewAccount(t)
	entry := createNewEntry(t, account1)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotZero(t, entry2.ID)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
}

func TestListEntries(t *testing.T) {
	account1 := createNewAccount(t)

	for i := 0; i < 10; i++ {
		createNewEntry(t, account1)
	}

	arg := ListEntriesParams{
		AccountID: pgtype.Int8{
			Int64: account1.ID,
			Valid: true,
		},
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
