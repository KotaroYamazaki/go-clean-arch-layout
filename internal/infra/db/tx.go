package db

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ctxTxKey = struct{}{}

func GetContextExecutor(ctx context.Context) boil.ContextExecutor {
	if tx, ok := ctx.Value(&ctxTxKey).(*sql.Tx); ok {
		return tx
	}

	return boil.GetContextDB()
}

func RunTransaction(ctx context.Context, fn func(context.Context) error) error {
	db, ok := GetContextExecutor(ctx).(boil.ContextBeginner)
	if !ok {
		panic("The database in the context does not support boil.ContextBeginner")
	}

	return runTxWithDB(ctx, db, fn)
}

func runTxWithDB(ctx context.Context, db boil.ContextBeginner, fn func(context.Context) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
			panic(err)
		}
	}()

	ctxWithTx := context.WithValue(ctx, &ctxTxKey, tx)
	if err := fn(ctxWithTx); err != nil {
		_ = tx.Rollback()

		return err
	}
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()

		return err
	}

	return nil
}
