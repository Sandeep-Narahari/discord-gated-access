package database

import (
	"fmt"

	db "github.com/forbole/juno/v3/database"
	"github.com/forbole/juno/v3/database/postgresql"
	juno "github.com/forbole/juno/v3/types"
	"github.com/jmoiron/sqlx"
)

var _ db.Database = &Db{}

type Db struct {
	*postgresql.Database
	Sqlx *sqlx.DB
}

func Cast(database db.Database) *Db {
	ajunoDb, ok := (database).(*Db)
	if !ok {
		panic(fmt.Errorf("database is not a Ajuno instance"))
	}

	return ajunoDb
}

func Builder(ctx *db.Context) (db.Database, error) {
	database, err := postgresql.Builder(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("DATABASE", database)

	psqlDb, ok := (database).(*postgresql.Database)
	fmt.Println("PSQLDB", psqlDb)
	if !ok {
		return nil, fmt.Errorf("invalid database type")
	}

	return &Db{
		Database: psqlDb,
		Sqlx:     sqlx.NewDb(psqlDb.Sql, "postgresql"),
	}, nil
}

func (db *Db) SaveTx(_ *juno.Tx) error {
	return nil
}

// HasValidator overrides postgresql.Database to perform a no-op
func (db *Db) HasValidator(_ string) (bool, error) {
	return true, nil
}

// SaveValidators overrides postgresql.Database to perform a no-op
func (db *Db) SaveValidators(_ []*juno.Validator) error {
	return nil
}

// SaveCommitSignatures overrides postgresql.Database to perform a no-op
func (db *Db) SaveCommitSignatures(_ []*juno.CommitSig) error {
	return nil
}

// SaveMessage overrides postgresql.Database to perform a no-op
func (db *Db) SaveMessage(_ *juno.Message) error {
	return nil
}
