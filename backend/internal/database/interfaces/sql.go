package interfaces

import (
	"context"
	"database/sql"
)

type SQLInterface interface {
	Conn(ctx context.Context) (SQLConnInterface, error)
	PrepareContext(ctx context.Context, query string) (StmtInterface, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (RowsInterface, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (SqlResultInterface, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) RowInterface
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TransactionInterface, error)
	GetDB() *sql.DB
	Close() error
	Exec(query string, args ...interface{}) (SqlResultInterface, error)
}

type SQLConnInterface interface {
	PrepareContext(ctx context.Context, query string) (StmtInterface, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (RowsInterface, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (SqlResultInterface, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) RowInterface
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TransactionInterface, error)
	Close() error
	UseDatabase(ctx context.Context, dbName string) error
	TableExists(ctx context.Context, tableName string) (bool, error)
}

type TransactionInterface interface {
	Commit() error
	Exec(query string, args ...any) (SqlResultInterface, error)
	ExecContext(ctx context.Context, query string, args ...any) (SqlResultInterface, error)
	Prepare(query string) (StmtInterface, error)
	PrepareContext(ctx context.Context, query string) (StmtInterface, error)
	Query(query string, args ...any) (RowsInterface, error)
	QueryContext(ctx context.Context, query string, args ...any) (RowsInterface, error)
	QueryRow(query string, args ...any) RowInterface
	QueryRowContext(ctx context.Context, query string, args ...any) RowInterface
	Rollback() error
}
type StmtInterface interface {
	Close() error
	ExecContext(ctx context.Context, args ...interface{}) (SqlResultInterface, error)
	QueryContext(ctx context.Context, args ...interface{}) (RowsInterface, error)
}

type RowsInterface interface {
	Close() error
	Err() error
	Next() bool
	Scan(dest ...interface{}) error
}

type RowInterface interface {
	Scan(dest ...interface{}) error
}

type SQLQuerier interface {
	PrepareContext(ctx context.Context, query string) (StmtInterface, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (RowsInterface, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (SqlResultInterface, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) RowInterface
	// 	GetConn() (*sql.Conn, error)
	// 	GetDB() (*sql.DB, error)
	// 	GetTransaction() *sql.Tx
}

type SqlResultInterface interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
