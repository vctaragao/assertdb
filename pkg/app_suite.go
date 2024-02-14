package pkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
)

type DBConnection struct {
	dbConn *bun.DB
	dbTx   *bun.Tx
}

func (db *DBConnection) Rollback() {
	if err := db.dbTx.Rollback(); err != nil {
		panic(err)
	}
}

func (db *DBConnection) BeginTx(ctx context.Context) {
	tx, err := db.dbConn.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	db.dbTx = &tx
}

func (db *DBConnection) Close() {
	db.Rollback()
	db.dbConn.Close()
}

type IntegrationSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *IntegrationSuite) SetupSuite() {
	s.initApp()
}

func (s *IntegrationSuite) TearDownSuite() {
	s.CloseApp()
}

func (s *IntegrationSuite) TearDownTest() {
	// s.softonDB.Rollback()
	// s.softonDB.BeginTx(s.ctx)
}

func (s *IntegrationSuite) initApp() {
	// go routerInstance.Start(":9200")
}

func (s *IntegrationSuite) CloseApp() {
}

func (s *IntegrationSuite) NewTestSuiteDbConn(dbConn *bun.DB) DBConnection {
	conn := DBConnection{dbConn: dbConn}
	conn.BeginTx(s.ctx)

	return conn
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}

	suite.Run(t, new(IntegrationSuite))
}
