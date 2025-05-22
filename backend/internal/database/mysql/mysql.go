package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type MySQL struct {
	db *sql.DB
}
type MySQLConn struct {
	conn *sql.Conn
}
type MySQLTransaction struct {
	tx *sql.Tx
}
type MySQLStmt struct {
	stmt *sql.Stmt
}

type MySQLRows struct {
	rows *sql.Rows
}

type MySQLRow struct {
	row *sql.Row
}

type MySQLDataConn struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

func NewMySQLConnection(md MySQLDataConn, maxConn *int) (interfaces.SQLInterface, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", md.User, md.Password, md.Host, md.Port, md.Database)
	fmt.Println("Connecting to MySQL database...")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if maxConn != nil {
		db.SetMaxOpenConns(*maxConn)
	}
	db.SetConnMaxLifetime(time.Second * 12)
	// db.SetConnMaxIdleTime(time.Minute * 5)

	return &MySQL{db}, nil
}

func (m *MySQL) Conn(ctx context.Context) (interfaces.SQLConnInterface, error) {
	stats := m.db.Stats()
	openConns := stats.OpenConnections
	inUse := stats.InUse
	maxOpenConns := stats.MaxOpenConnections
	log.Println("Open connections:", openConns)
	log.Println("InUse connections:", inUse)

	if maxOpenConns != 0 && maxOpenConns == inUse {
		return nil, fmt.Errorf("max open connections")
	}

	conn, err := m.db.Conn(ctx)
	if err != nil {
		log.Println("Error obteniendo conexión desde db.Conn:", err.Error())
		return nil, err
	}

	_, err = conn.ExecContext(ctx, "SET wait_timeout = 10")
	if err != nil {
		return nil, err
	}

	// Configura un temporizador para cerrar la conexión después de 30 segundos de inactividad.
	timer := time.NewTimer(300 * time.Second)

	go func() {
		select {
		case <-timer.C:
			err := conn.Close()
			if err != nil && !strings.Contains(err.Error(), "connection is already closed") {
				log.Println("Error cerrando la conexión:", err.Error())
			} else if err == nil {
				log.Println("Connection closed due to timeout")
			}
		case <-ctx.Done():
			// Cancelación del contexto, cerrar la conexión.
			err := conn.Close()
			if err != nil && !strings.Contains(err.Error(), "connection is already closed") {
				log.Println("Error cerrando la conexión:", err.Error())
			} else if err == nil {
				log.Println("Connection closed due to context cancellation")
			}
		}
	}()

	return &MySQLConn{conn}, nil
}

func (m *MySQL) Exec(query string, args ...interface{}) (interfaces.SqlResultInterface, error) {
	return m.db.Exec(query, args...)
}

func (m *MySQL) ExecContext(ctx context.Context, query string, args ...interface{}) (interfaces.SqlResultInterface, error) {
	return m.db.ExecContext(ctx, query, args...)
}

func (m *MySQL) QueryContext(ctx context.Context, query string, args ...interface{}) (interfaces.RowsInterface, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	return &MySQLRows{rows}, err
}

func (m *MySQL) QueryRowContext(ctx context.Context, query string, args ...interface{}) interfaces.RowInterface {
	return m.db.QueryRowContext(ctx, query, args...)
}

func (m *MySQL) PrepareContext(ctx context.Context, query string) (interfaces.StmtInterface, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	return &MySQLStmt{stmt}, err
}

func (m *MySQL) BeginTx(ctx context.Context, opts *sql.TxOptions) (interfaces.TransactionInterface, error) {
	tx, err := m.db.BeginTx(ctx, opts)
	return &MySQLTransaction{tx}, err
}

func (m *MySQL) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}
func (m *MySQL) GetDB() *sql.DB {
	return m.db
}

func (c *MySQLConn) TableExists(ctx context.Context, tableName string) (bool, error) {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := c.conn.QueryContext(ctx, query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (c *MySQLConn) UseDatabase(ctx context.Context, dbName string) error {
	_, err := c.conn.ExecContext(ctx, "USE "+dbName)
	if err != nil {
		return err
	}
	log.Println("Ejecutando cambio de DB a " + dbName + "")
	return nil
}

func (c *MySQLConn) ExecContext(ctx context.Context, query string, args ...interface{}) (interfaces.SqlResultInterface, error) {
	return c.conn.ExecContext(ctx, query, args...)
}

func (c *MySQLConn) QueryContext(ctx context.Context, query string, args ...interface{}) (interfaces.RowsInterface, error) {
	rows, err := c.conn.QueryContext(ctx, query, args...)
	return &MySQLRows{rows}, err
}

func (c *MySQLConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) interfaces.RowInterface {
	return c.conn.QueryRowContext(ctx, query, args...)
}

func (c *MySQLConn) PrepareContext(ctx context.Context, query string) (interfaces.StmtInterface, error) {
	stmt, err := c.conn.PrepareContext(ctx, query)
	return &MySQLStmt{stmt}, err
}

func (c *MySQLConn) BeginTx(ctx context.Context, opts *sql.TxOptions) (interfaces.TransactionInterface, error) {
	tx, err := c.conn.BeginTx(ctx, opts)
	return &MySQLTransaction{tx}, err
}

func (c *MySQLConn) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (s *MySQLStmt) Close() error {
	return s.stmt.Close()
}

func (s *MySQLStmt) ExecContext(ctx context.Context, args ...interface{}) (interfaces.SqlResultInterface, error) {
	return s.stmt.ExecContext(ctx, args...)
}

func (s *MySQLStmt) QueryContext(ctx context.Context, args ...interface{}) (interfaces.RowsInterface, error) {
	rows, err := s.stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return &MySQLRows{rows}, nil
}
func (r *MySQLRows) Close() error {
	return r.rows.Close()
}

func (r *MySQLRows) Err() error {
	return r.rows.Err()
}

func (r *MySQLRows) Next() bool {
	return r.rows.Next()
}

func (r *MySQLRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *MySQLRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

func (t *MySQLTransaction) Commit() error {
	return t.tx.Commit()
}

func (t *MySQLTransaction) Exec(query string, args ...any) (interfaces.SqlResultInterface, error) {
	return t.tx.Exec(query, args...)
}
func (t *MySQLTransaction) ExecContext(ctx context.Context, query string, args ...any) (interfaces.SqlResultInterface, error) {
	return t.tx.ExecContext(ctx, query, args...)
}
func (t *MySQLTransaction) Prepare(query string) (interfaces.StmtInterface, error) {
	stmt, err := t.tx.Prepare(query)
	return &MySQLStmt{stmt}, err
}

func (t *MySQLTransaction) PrepareContext(ctx context.Context, query string) (interfaces.StmtInterface, error) {
	stmt, err := t.tx.PrepareContext(ctx, query)
	return &MySQLStmt{stmt}, err
}

func (t *MySQLTransaction) Query(query string, args ...any) (interfaces.RowsInterface, error) {
	rows, err := t.tx.Query(query, args...)
	return &MySQLRows{rows}, err
}

func (t *MySQLTransaction) QueryContext(ctx context.Context, query string, args ...any) (interfaces.RowsInterface, error) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &MySQLRows{rows}, nil
}

func (t *MySQLTransaction) QueryRow(query string, args ...any) interfaces.RowInterface {
	return t.tx.QueryRow(query, args...)
}

func (t *MySQLTransaction) QueryRowContext(ctx context.Context, query string, args ...any) interfaces.RowInterface {
	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *MySQLTransaction) Rollback() error {
	return t.tx.Rollback()
}
