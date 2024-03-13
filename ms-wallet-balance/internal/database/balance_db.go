package database

import (
	"database/sql"

	"github.com/matheuslch/fc-ms-wallet-balance/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{DB: db}
}

func (b *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (id, account_id, balance, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.BalanceAmount, balance.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceDB) Update(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET balance = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.BalanceAmount, balance.UpdatedAt, balance.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceDB) FindByAccountID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := b.DB.Prepare("SELECT id, account_id, balance, updated_at FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountID)

	err = row.Scan(
		&balance.ID,
		&balance.AccountID,
		&balance.BalanceAmount,
		&balance.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return &entity.Balance{}, nil
	case err != nil:
		return nil, err
	default:
		return &balance, nil
	}

}
