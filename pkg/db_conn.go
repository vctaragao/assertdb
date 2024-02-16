package pkg

import (
	"context"
	"database/sql"
	"log"
)

type DBConnection struct {
	DbConn *sql.DB
	DbTx   *sql.Tx
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

	db.DbTx = tx
}

func (db *DBConnection) Close() error {
	return db.DbConn.Close()
}
