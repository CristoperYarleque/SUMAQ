package database

import (
	"fmt"

	"github.com/cyarleque/sumaq/config"
	"github.com/cyarleque/sumaq/internal/database/interfaces"
	"github.com/cyarleque/sumaq/internal/database/mysql"
)

type Connections struct {
	SQL interfaces.SQLInterface
}

type ConnectionFactoryInterface interface {
	CreateConnection(dbType string, config *config.DatabaseConfig, maxConn *int) (*Connections, error)
}
type ConnectionFactory struct {
}

func NewConnectionFactory() ConnectionFactoryInterface {
	return &ConnectionFactory{}
}

const (
	MySQLType   = "mysql"
	PotgresType = "potgres"
)

func (cf *ConnectionFactory) CreateConnection(dbType string, config *config.DatabaseConfig, maxConns *int) (*Connections, error) {

	switch dbType {
	case MySQLType, PotgresType:
		data := mysql.MySQLDataConn{Host: config.Host, User: config.Username, Password: config.Password, Database: config.Database, Port: config.Port}
		db, err := mysql.NewMySQLConnection(data, maxConns)
		return &Connections{SQL: db}, err
	default:
		return nil, fmt.Errorf("unsupported database type")
	}
}
