package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kurnhyalcantara/teman-petani/libs/database/wrapper"
)

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	SslMode      string
	TimeZone     string
	MaxRetry     int
	Timeout      time.Duration
}

type DB struct {
	DriverName    string
	Config        *Config
	SqlDb         *sql.DB
	SqlTx         *sql.Tx
	Counter       int
	DBConnWrapper wrapper.DatabaseConnectionWrapper
}

func InitConnection(driverName string, config *Config) *DB {
	return &DB{
		DriverName:    driverName,
		Config:        config,
		DBConnWrapper: wrapper.DatabaseConnectionWrapper{},
	}
}

func (db *DB) Connect() error {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		db.Config.Host,
		db.Config.Port,
		db.Config.User,
		db.Config.Password,
		db.Config.DatabaseName,
		db.Config.SslMode,
		db.Config.TimeZone)

	var err error
	db.SqlDb, err = db.DBConnWrapper.Open(db.DriverName, connString)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) TryConnect() error {
	for {
		db.AddCounter()
		log.Printf("trying to connect %d times...", db.Counter)

		err := db.Connect()
		if err == nil {
			db.Counter = 0
			return nil
		}

		if db.Counter >= db.Config.MaxRetry {
			log.Println("stop reconnecting max retries exceeding")
			return err
		}
	}
}

func (db *DB) CheckConnection() error {
	if db.Counter > 0 {
		log.Println("server still trying connect to db")
	}
	if err := db.DBConnWrapper.Ping(); err != nil {
		db.DBConnWrapper.Close()
		return db.TryConnect()
	}
	return nil
}

func (db *DB) StartTransaction() error {
	var err error
	db.SqlTx, err = db.DBConnWrapper.Begin()
	return err
}

func (db *DB) AddCounter() {
	db.Counter++
}

func (db *DB) CloseConnection() error {
	if err := db.DBConnWrapper.Close(); err != nil {
		log.Fatal("failed close connection: ", err)
	}
	return nil
}

func (db *DB) GetTimeout() time.Duration {
	return db.Config.Timeout
}

func (db *DB) SetMaxIdleConnections(n int) {
	db.DBConnWrapper.SetMaxIdleConnections(n)
}

func (db *DB) SetMaxOpenConnections(n int) {
	db.DBConnWrapper.SetMaxOpenConnections(n)
}
