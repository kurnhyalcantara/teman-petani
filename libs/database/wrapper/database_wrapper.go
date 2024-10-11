package wrapper

import "database/sql"

type DatabaseConnectionWrapper struct {
	DbSql *sql.DB
}

// Begin implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) Begin() (*sql.Tx, error) {
	return d.DbSql.Begin()
}

// Close implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) Close() error {
	return d.DbSql.Close()
}

// Open implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) Open(driverName string, source string) (*sql.DB, error) {
	db, err := sql.Open(driverName, source)
	if err != nil {
		return nil, err
	}

	d.DbSql = db
	return db, nil
}

// Ping implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) Ping() error {
	return d.DbSql.Close()
}

// SetMaxIdleConnections implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) SetMaxIdleConnections(n int) {
	d.DbSql.SetMaxIdleConns(n)
}

// SetMaxOpenConnections implements DatabaseConnectionInterface.
func (d *DatabaseConnectionWrapper) SetMaxOpenConnections(n int) {
	d.DbSql.SetMaxOpenConns(n)
}
