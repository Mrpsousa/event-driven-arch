package datebase

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrpsousa/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(225), name varchar(225), email varchar(225), created_at date)")
	db.Exec("Create table accounts (id varchar(225), client_id varchar(225), balance int, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("roger", "eu@email.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindById() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountResult, err := s.accountDB.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountResult.ID)
	s.Equal(account.Client.ID, accountResult.Client.ID)
}
