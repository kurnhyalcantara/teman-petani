package wrapper

import "database/sql"

type DatabaseConnectionInterface interface {
	Open(driverName string, source string) (*sql.DB, error)
	Ping() error
	Close() error
	SetMaxIdleConnections(n int)
	SetMaxOpenConnections(n int)
	Begin() (*sql.Tx, error)
}
