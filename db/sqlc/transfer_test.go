package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"studies.com/util"
)

func createNewTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: pgtype.Int8{
			Int64: account1.ID,
			Valid: true,
		},
		ToAccountID: pgtype.Int8{
			Int64: account2.ID,
			Valid: true,
		},
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, transfer.ID)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createNewAccount(t)
	account2 := createNewAccount(t)
	createNewTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createNewAccount(t)
	account2 := createNewAccount(t)
	transfer1 := createNewTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestListTransfers(t *testing.T) {
	account1 := createNewAccount(t)
	account2 := createNewAccount(t)

	for i := 0; i < 10; i++ {
		createNewTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: pgtype.Int8{
			Int64: account1.ID,
			Valid: true,
		},
		ToAccountID: pgtype.Int8{
			Int64: account2.ID,
			Valid: true,
		},
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
