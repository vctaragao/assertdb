package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func Connect() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	sqldb.SetMaxIdleConns(1000)
	sqldb.SetConnMaxLifetime(0)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	return db
}
