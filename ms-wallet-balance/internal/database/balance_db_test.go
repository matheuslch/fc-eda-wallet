package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/matheuslch/fc-ms-wallet-balance/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE balances (id varchar(255), account_id varchar(255), balance float, updated_at datetime)")

	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance, _ := entity.NewBalance(
		"1",
		500,
		time.Now(),
	)
	err := s.balanceDB.Save(balance)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestUpdate() {
	balance, _ := entity.NewBalance(
		"1",
		500,
		time.Now(),
	)
	err := s.balanceDB.Save(balance)
	s.Nil(err)

	balance.BalanceAmount = 1000
	err = s.balanceDB.Update(balance)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestFindByAccountID() {
	balance, _ := entity.NewBalance(
		"1",
		500,
		time.Now(),
	)
	err := s.balanceDB.Save(balance)
	s.Nil(err)

	balanceFind, err := s.balanceDB.FindByAccountID(balance.AccountID)
	s.Nil(err)
	s.Equal(balanceFind.ID, balance.ID)
	s.Equal(balanceFind.AccountID, balance.AccountID)
	s.Equal(balanceFind.BalanceAmount, balance.BalanceAmount)
}
