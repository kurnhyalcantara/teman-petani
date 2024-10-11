package main

import (
	"log"
	"strconv"
	"time"

	"github.com/kurnhyalcantara/teman-petani/libs/database"
	_ "github.com/lib/pq"
)

func startDBConnections() {
	log.Println("Starting DB Connections")

	openConnectionPostgresql()
}

func closeDBConnections() {
	log.Println("Closing DB Connections")

	closeConnectionPostgresql()
}

func openConnectionPostgresql() {
	log.Println("PostgreSQL - Connecting")

	maxRetry, err := strconv.Atoi(appConfig.DbRetry)
	if err != nil {
		maxRetry = 3
	}

	timeout, err := strconv.Atoi(appConfig.DbTimeout)
	if err != nil {
		timeout = 120
	}

	dbSql = database.InitConnection("postgres", &database.Config{
		Host:         appConfig.DbHost,
		Port:         appConfig.DbPort,
		User:         appConfig.DbUsername,
		Password:     appConfig.DbPassword,
		DatabaseName: appConfig.DbName,
		SslMode:      appConfig.DbSslMode,
		TimeZone:     appConfig.DbTimeZone,
		MaxRetry:     maxRetry,
		Timeout:      time.Duration(timeout),
	})

	if err := dbSql.Connect(); err != nil {
		log.Fatal("Failed to connecting db: ", err)
	}

	maxIdle, err := strconv.Atoi(appConfig.DbMaxIdle)
	if err != nil {
		maxIdle = 2
	}

	maxOpen, err := strconv.Atoi(appConfig.DbMaxOpen)
	if err != nil {
		maxOpen = 100
	}

	dbSql.SetMaxIdleConnections(maxIdle)
	dbSql.SetMaxOpenConnections(maxOpen)

	log.Println("PostgreSQL - Connected")
}

func closeConnectionPostgresql() {
	log.Println("PostgreSQL - Closing")

	if err := dbSql.CloseConnection(); err != nil {
		log.Fatal("Failed to closing connection: ", err)
	}

	log.Println("PostgreSQL - Closed")
}
