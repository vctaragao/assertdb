package pkg

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

type DBConnection struct {
	DbConn *bun.DB
	DbTx   *bun.Tx
}

func (db *DBConnection) Rollback() {
	if err := db.DbTx.Rollback(); err != nil {
		log.Fatal(err)
	}
}

func (db *DBConnection) BeginTx(ctx context.Context) {
	tx, err := db.DbConn.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	db.DbTx = &tx
}

func (db *DBConnection) Close() error {
	return db.DbConn.Close()
}
