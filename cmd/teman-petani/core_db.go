package main

import (
	"strconv"
	"time"

	"github.com/kurnhyalcantara/teman-petani/libs/database"
	_ "github.com/lib/pq"
)

func startDBConnections() {
	logger.Info().Msg("Starting DB Connections")

	openConnectionPostgresql()
}

func closeDBConnections() {
	logger.Info().Msg("Closing DB Connections")

	closeConnectionPostgresql()
}

func openConnectionPostgresql() {
	logger.Info().Msg("PostgreSQL - Connecting")

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
		logger.Fatal().Msgf("Failed to connecting db: %v", err)
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

	logger.Info().Msg("PostgreSQL - Connected")
}

func closeConnectionPostgresql() {
	logger.Info().Msg("PostgreSQL - Closing")

	if err := dbSql.CloseConnection(); err != nil {
		logger.Fatal().Msgf("Failed to closing connection: %v", err)
	}

	logger.Info().Msg("PostgreSQL - Closed")
}
