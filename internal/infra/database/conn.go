package database

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/vctaragao/assertdb/internal/entity"
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

func Migrate(db *bun.DB) error {
	_, err := db.
		NewCreateTable().
		IfNotExists().
		Model((*entity.Banana)(nil)).
		Exec(context.Background())
	return err
}
