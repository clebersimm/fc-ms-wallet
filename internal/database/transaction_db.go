package database

import (
	"database/sql"

	"github.com/clebersimm/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (t *TransactionDB) Create(transction *entity.Transaction) error {
	stmt, err := t.DB.Prepare("INSERT INTO transactions(id, account_id_from, account_id_to, amount, created_at) VALUEs (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transction.ID, transction.AccountFrom.ID, transction.AcctountTo.ID, transction.Amount, transction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
